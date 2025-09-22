package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plate/service/internal/models"
)

// Project handlers
func (s *Server) handleListProjects(c *gin.Context) {
	// Return developer-friendly fake data
	projects := []ProjectResponse{
		{
			ID:          1,
			Name:        "web-app",
			Description: "Frontend React application for the main website",
			Repository:  "https://github.com/yourorg/web-app.git",
			Runtime:     "nodejs",
			Status:      "active",
			LastDeploy:  "2 hours ago",
			Environments: []string{"development", "staging", "production"},
		},
		{
			ID:          2,
			Name:        "api-service",
			Description: "REST API backend service written in Go",
			Repository:  "https://github.com/yourorg/api-service.git",
			Runtime:     "go",
			Status:      "active",
			LastDeploy:  "1 day ago",
			Environments: []string{"development", "staging", "production"},
		},
		{
			ID:          3,
			Name:        "data-processor",
			Description: "Python service for data processing and analytics",
			Repository:  "https://github.com/yourorg/data-processor.git",
			Runtime:     "python",
			Status:      "building",
			LastDeploy:  "never",
			Environments: []string{"development"},
		},
		{
			ID:          4,
			Name:        "worker-service",
			Description: "Background job processing service",
			Repository:  "https://github.com/yourorg/worker-service.git",
			Runtime:     "go",
			Status:      "inactive",
			LastDeploy:  "1 week ago",
			Environments: []string{"production"},
		},
	}
	c.JSON(http.StatusOK, projects)
}

func (s *Server) handleCreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.services.Project.Create(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (s *Server) handleGetProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	project, err := s.services.Project.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (s *Server) handleUpdateProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project.ID = uint(id)
	if err := s.services.Project.Update(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (s *Server) handleDeleteProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	if err := s.services.Project.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}