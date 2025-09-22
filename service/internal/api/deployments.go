package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/plate/service/internal/models"
)

func (s *Server) handleListDeployments(c *gin.Context) {
	// Return developer-friendly fake data
	deployments := []DeploymentResponse{
		{
			ID:          1,
			Application: "web-app",
			Environment: "production",
			Version:     "v1.2.3",
			Status:      "live",
			URL:         "https://web-app.yourapp.com",
			DeployedAt:  "2 hours ago",
			Duration:    "45s",
		},
		{
			ID:          2,
			Application: "web-app",
			Environment: "staging",
			Version:     "v1.2.4-beta",
			Status:      "live",
			URL:         "https://web-app-staging.yourapp.com",
			DeployedAt:  "4 hours ago",
			Duration:    "52s",
		},
		{
			ID:          3,
			Application: "api-service",
			Environment: "production",
			Version:     "v2.1.0",
			Status:      "live",
			URL:         "https://api.yourapp.com",
			DeployedAt:  "1 day ago",
			Duration:    "1m 23s",
		},
		{
			ID:          4,
			Application: "api-service",
			Environment: "development",
			Version:     "v2.1.1-dev",
			Status:      "failed",
			URL:         "",
			DeployedAt:  "2 hours ago",
			Duration:    "failed at 30s",
		},
		{
			ID:          5,
			Application: "worker-service",
			Environment: "production",
			Version:     "v1.0.5",
			Status:      "live",
			URL:         "",
			DeployedAt:  "3 days ago",
			Duration:    "1m 10s",
		},
		{
			ID:          6,
			Application: "data-processor",
			Environment: "development",
			Version:     "v0.1.0-dev",
			Status:      "building",
			URL:         "",
			DeployedAt:  "building...",
			Duration:    "2m 15s",
		},
	}
	c.JSON(http.StatusOK, deployments)
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
		// Return detailed status for all or specific environment
		statuses := []DetailedStatusResponse{
			{
				Environment: "production",
				Region:      "us-west-2",
				Status:      "healthy",
				Applications: []ApplicationStatusResponse{
					{
						Application: "web-app",
						Environment: "production",
						Status:      "live",
						Health:      "healthy",
						URL:         "https://web-app.yourapp.com",
						Version:     "v1.2.3",
						Uptime:      "99.9%",
						LastUpdate:  "2 hours ago",
					},
					{
						Application: "api-service",
						Environment: "production",
						Status:      "live",
						Health:      "healthy",
						URL:         "https://api.yourapp.com",
						Version:     "v2.1.0",
						Uptime:      "99.8%",
						LastUpdate:  "1 day ago",
					},
				},
				LastUpdate: "2025-09-19T10:30:00Z",
			},
			{
				Environment: "staging",
				Region:      "us-west-2",
				Status:      "healthy",
				Applications: []ApplicationStatusResponse{
					{
						Application: "web-app",
						Environment: "staging",
						Status:      "live",
						Health:      "healthy",
						URL:         "https://web-app-staging.yourapp.com",
						Version:     "v1.2.4-beta",
						Uptime:      "99.5%",
						LastUpdate:  "4 hours ago",
					},
				},
				LastUpdate: "2025-09-19T09:15:00Z",
			},
			{
				Environment: "development",
				Region:      "us-west-2",
				Status:      "degraded",
				Applications: []ApplicationStatusResponse{
					{
						Application: "api-service",
						Environment: "development",
						Status:      "failed",
						Health:      "unhealthy",
						URL:         "",
						Version:     "v2.1.1-dev",
						Uptime:      "0%",
						LastUpdate:  "2 hours ago",
					},
					{
						Application: "data-processor",
						Environment: "development",
						Status:      "building",
						Health:      "unknown",
						URL:         "",
						Version:     "v0.1.0-dev",
						Uptime:      "N/A",
						LastUpdate:  "building...",
					},
				},
				LastUpdate: "2025-09-19T08:45:00Z",
			},
		}

		if env != "" {
			for _, status := range statuses {
				if status.Environment == env {
					c.JSON(http.StatusOK, []DetailedStatusResponse{status})
					return
				}
			}
			c.JSON(http.StatusOK, []DetailedStatusResponse{})
			return
		}

		c.JSON(http.StatusOK, statuses)
		return
	}

	// Return simple status
	status := gin.H{
		"web-app-production":         "live",
		"web-app-staging":           "live", 
		"api-service-production":    "live",
		"api-service-development":   "failed",
		"worker-service-production": "live",
		"data-processor-development": "building",
	}

	if env != "" {
		filtered := gin.H{}
		for key, value := range status {
			if env == "production" && strings.Contains(key, "production") {
				filtered[key] = value
			} else if env == "staging" && strings.Contains(key, "staging") {
				filtered[key] = value
			} else if env == "development" && strings.Contains(key, "development") {
				filtered[key] = value
			}
		}
		c.JSON(http.StatusOK, filtered)
		return
	}

	c.JSON(http.StatusOK, status)
}