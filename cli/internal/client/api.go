package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type APIClient struct {
	client  *resty.Client
	baseURL string
	token   string
}

type DeploymentStatus struct {
	Environment string    `json:"environment"`
	Status      string    `json:"status"`
	URL         string    `json:"url,omitempty"`
	Pods        []Pod     `json:"pods,omitempty"`
	LastUpdate  time.Time `json:"last_update"`
}

type Pod struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Ready  string `json:"ready"`
}

func NewAPIClient() *APIClient {
	baseURL := viper.GetString("api-url")
	token := viper.GetString("token")

	client := resty.New()
	client.SetTimeout(30 * time.Second)
	
	if token != "" {
		client.SetAuthToken(token)
	}

	return &APIClient{
		client:  client,
		baseURL: baseURL,
		token:   token,
	}
}

func (c *APIClient) Deploy(environment string) error {
	resp, err := c.client.R().
		SetBody(map[string]string{
			"environment": environment,
		}).
		Post(c.baseURL + "/api/v1/deploy")

	if err != nil {
		return fmt.Errorf("failed to make deploy request: %w", err)
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("deploy request failed with status %d: %s", resp.StatusCode(), resp.String())
	}

	return nil
}

func (c *APIClient) GetStatus(environment string, detailed bool) (string, error) {
	req := c.client.R()
	
	if environment != "" {
		req.SetQueryParam("env", environment)
	}
	
	if detailed {
		req.SetQueryParam("detailed", "true")
	}

	resp, err := req.Get(c.baseURL + "/api/v1/status")
	if err != nil {
		return "", fmt.Errorf("failed to get status: %w", err)
	}

	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("status request failed with status %d: %s", resp.StatusCode(), resp.String())
	}

	if detailed {
		return c.formatDetailedStatus(resp.Body())
	}

	return resp.String(), nil
}

func (c *APIClient) WatchDeployment(environment string) error {
	// Simple polling implementation
	for {
		status, err := c.GetStatus(environment, false)
		if err != nil {
			return err
		}

		fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), status)

		// Check if deployment is complete
		if c.isDeploymentComplete(status) {
			break
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}

func (c *APIClient) formatDetailedStatus(data []byte) (string, error) {
	var statuses []DeploymentStatus
	if err := json.Unmarshal(data, &statuses); err != nil {
		return string(data), nil // Fallback to raw response
	}

	result := ""
	for _, status := range statuses {
		result += fmt.Sprintf("Environment: %s\n", status.Environment)
		result += fmt.Sprintf("Status: %s\n", status.Status)
		if status.URL != "" {
			result += fmt.Sprintf("URL: %s\n", status.URL)
		}
		result += fmt.Sprintf("Last Update: %s\n", status.LastUpdate.Format("2006-01-02 15:04:05"))
		
		if len(status.Pods) > 0 {
			result += "Pods:\n"
			for _, pod := range status.Pods {
				result += fmt.Sprintf("  - %s: %s (%s)\n", pod.Name, pod.Status, pod.Ready)
			}
		}
		result += "\n"
	}

	return result, nil
}

func (c *APIClient) isDeploymentComplete(status string) bool {
	// Simple heuristic - check if status contains "running" or "failed"
	return fmt.Sprintf("%s", status) != "" // Placeholder logic
}