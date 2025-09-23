package api

import "time"

// Developer-friendly response structures that hide infrastructure complexity

type ProjectResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Repository  string `json:"repository"`
	Runtime     string `json:"runtime"`
	Status      string `json:"status"`
	LastDeploy  string `json:"last_deploy"`
	Environments []string `json:"environments"`
	CreatedAt   time.Time `json:"created_at"`
}

type DeploymentResponse struct {
	ID                uint   `json:"id"`
	Application       string `json:"application"`
	Environment       string `json:"environment"`
	Version           string `json:"version"`
	Status            string `json:"status"`
	URL               string `json:"url,omitempty"`
	DeployedAt        string `json:"deployed_at"`
	Duration          string `json:"duration,omitempty"`
	DesiredReplicas   int32  `json:"desired_replicas"`
	ReadyReplicas     int32  `json:"ready_replicas"`
	AvailableReplicas int32  `json:"available_replicas"`
}

type EnvironmentResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"` // development, staging, production
	Region string `json:"region"`
	Status string `json:"status"`
}

type ApplicationStatusResponse struct {
	Application string `json:"application"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
	Health      string `json:"health"`
	URL         string `json:"url,omitempty"`
	Version     string `json:"version"`
	Uptime      string `json:"uptime"`
	LastUpdate  string `json:"last_update"`
}

type DetailedStatusResponse struct {
	Environment  string                      `json:"environment"`
	Region       string                      `json:"region"`
	Status       string                      `json:"status"`
	Applications []ApplicationStatusResponse `json:"applications"`
	LastUpdate   string                      `json:"last_update"`
}

type BuildLogResponse struct {
	Timestamp string `json:"timestamp"`
	Stage     string `json:"stage"`
	Message   string `json:"message"`
	Level     string `json:"level"`
}

type DeploymentLogsResponse struct {
	Application string             `json:"application"`
	Environment string             `json:"environment"`
	Logs        []BuildLogResponse `json:"logs"`
}