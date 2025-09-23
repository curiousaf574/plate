package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/plate/service/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) handleListDeployments(c *gin.Context) {
	// Get deployments from all Plate-managed namespaces
	namespaces := []string{"dev", "staging", "production"}
	var allDeployments []DeploymentResponse

	for _, namespace := range namespaces {
		deployments, err := s.services.Kubernetes.GetClientset().AppsV1().Deployments(namespace).List(c, metav1.ListOptions{
			LabelSelector: "managed-by=plate",
		})
		if err != nil {
			continue
		}

		for i, deployment := range deployments.Items {
			appName := deployment.Labels["app"]
			if appName == "" {
				appName = deployment.Name
			}

			version := deployment.Labels["version"]
			if version == "" {
				version = "latest"
			}

			status := "live"
			if deployment.Status.ReadyReplicas == 0 {
				status = "failed"
			} else if deployment.Status.ReadyReplicas < *deployment.Spec.Replicas {
				status = "building"
			}

			deployedAt := "unknown"
			if deployment.CreationTimestamp.Time.IsZero() == false {
				deployedAt = deployment.CreationTimestamp.Format("2006-01-02 15:04")
			}

			// Generate URL for web services
			url := ""
			if strings.Contains(appName, "web") || strings.Contains(appName, "frontend") {
				url = fmt.Sprintf("https://%s.%s.yourapp.com", appName, namespace)
			} else if strings.Contains(appName, "api") {
				url = fmt.Sprintf("https://api.%s.yourapp.com", namespace)
			}

			allDeployments = append(allDeployments, DeploymentResponse{
				ID:                uint(len(allDeployments) + i + 1),
				Application:       appName,
				Environment:       namespace,
				Version:           version,
				Status:            status,
				URL:               url,
				DeployedAt:        deployedAt,
				Duration:          "N/A",
				DesiredReplicas:   *deployment.Spec.Replicas,
				ReadyReplicas:     deployment.Status.ReadyReplicas,
				AvailableReplicas: deployment.Status.AvailableReplicas,
			})
		}
	}

	c.JSON(http.StatusOK, allDeployments)
}

func (s *Server) handleCreateDeployment(c *gin.Context) {
	var deployment models.Deployment
	if err := c.ShouldBindJSON(&deployment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.services.Deployment.Create(&deployment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, deployment)
}

func (s *Server) handleGetDeployment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid deployment ID"})
		return
	}

	deployment, err := s.services.Deployment.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Deployment not found"})
		return
	}

	c.JSON(http.StatusOK, deployment)
}

func (s *Server) handleDeleteDeployment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid deployment ID"})
		return
	}

	if err := s.services.Deployment.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deployment deleted successfully"})
}

func (s *Server) handleGetDeploymentLogs(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid deployment ID"})
		return
	}

	logs, err := s.services.Deployment.GetLogs(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (s *Server) handleDeploy(c *gin.Context) {
	var req struct {
		ProjectID     uint   `json:"project_id" binding:"required"`
		EnvironmentID uint   `json:"environment_id" binding:"required"`
		Version       string `json:"version"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deployment, err := s.services.Deployment.Deploy(req.ProjectID, req.EnvironmentID, req.Version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, deployment)
}

func (s *Server) handleGetStatus(c *gin.Context) {
	env := c.Query("env")
	detailed := c.Query("detailed") == "true"

	if detailed {
		// Return detailed status from Kubernetes
		namespaces := []string{"dev", "staging", "production"}
		if env != "" {
			namespaces = []string{env}
		}

		var statuses []DetailedStatusResponse
		for _, namespace := range namespaces {
			deployments, err := s.services.Kubernetes.GetClientset().AppsV1().Deployments(namespace).List(c, metav1.ListOptions{
				LabelSelector: "managed-by=plate",
			})
			if err != nil {
				continue
			}

			var applications []ApplicationStatusResponse
			overallStatus := "healthy"

			for _, deployment := range deployments.Items {
				appName := deployment.Labels["app"]
				if appName == "" {
					appName = deployment.Name
				}
				version := deployment.Labels["version"]
				if version == "" {
					version = "latest"
				}

				status := "live"
				health := "healthy"
				uptime := "99.9%"
				
				if deployment.Status.ReadyReplicas == 0 {
					status = "failed"
					health = "unhealthy"
					uptime = "0%"
					overallStatus = "degraded"
				} else if deployment.Status.ReadyReplicas < *deployment.Spec.Replicas {
					status = "building"
					health = "unknown"
					uptime = "N/A"
					if overallStatus == "healthy" {
						overallStatus = "degraded"
					}
				}

				url := ""
				if strings.Contains(appName, "web") || strings.Contains(appName, "frontend") {
					url = fmt.Sprintf("https://%s.%s.yourapp.com", appName, namespace)
				} else if strings.Contains(appName, "api") {
					url = fmt.Sprintf("https://api.%s.yourapp.com", namespace)
				}

				lastUpdate := "unknown"
				if !deployment.CreationTimestamp.Time.IsZero() {
					lastUpdate = deployment.CreationTimestamp.Format("2006-01-02 15:04")
				}

				applications = append(applications, ApplicationStatusResponse{
					Application: appName,
					Environment: namespace,
					Status:      status,
					Health:      health,
					URL:         url,
					Version:     version,
					Uptime:      uptime,
					LastUpdate:  lastUpdate,
				})
			}

			if len(applications) > 0 {
				statuses = append(statuses, DetailedStatusResponse{
					Environment:  namespace,
					Region:       "local",
					Status:       overallStatus,
					Applications: applications,
					LastUpdate:   "2025-09-22T10:30:00Z",
				})
			}
		}

		c.JSON(http.StatusOK, statuses)
		return
	}

	// Return simple status from Kubernetes
	namespaces := []string{"dev", "staging", "production"}
	status := gin.H{}

	for _, namespace := range namespaces {
		deployments, err := s.services.Kubernetes.GetClientset().AppsV1().Deployments(namespace).List(c, metav1.ListOptions{
			LabelSelector: "managed-by=plate",
		})
		if err != nil {
			continue
		}

		for _, deployment := range deployments.Items {
			appName := deployment.Labels["app"]
			if appName == "" {
				appName = deployment.Name
			}

			key := fmt.Sprintf("%s-%s", appName, namespace)
			deploymentStatus := "live"
			
			if deployment.Status.ReadyReplicas == 0 {
				deploymentStatus = "failed"
			} else if deployment.Status.ReadyReplicas < *deployment.Spec.Replicas {
				deploymentStatus = "building"
			}

			status[key] = deploymentStatus
		}
	}

	// Filter by environment if specified
	if env != "" {
		filtered := gin.H{}
		for key, value := range status {
			if strings.Contains(key, "-"+env) {
				filtered[key] = value
			}
		}
		c.JSON(http.StatusOK, filtered)
		return
	}

	c.JSON(http.StatusOK, status)
}