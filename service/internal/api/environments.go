package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plate/service/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) handleListEnvironments(c *gin.Context) {
	// Get namespaces managed by Plate
	namespaces, err := s.services.Kubernetes.GetClientset().CoreV1().Namespaces().List(c, metav1.ListOptions{
		LabelSelector: "managed-by=plate",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get namespaces"})
		return
	}

	var environments []EnvironmentResponse
	for i, ns := range namespaces.Items {
		envType := ns.Labels["environment"]
		if envType == "" {
			envType = ns.Name
		}

		// Determine status based on namespace phase
		status := "active"
		if ns.Status.Phase != "Active" {
			status = "inactive"
		}

		environments = append(environments, EnvironmentResponse{
			ID:     uint(i + 1),
			Name:   ns.Name,
			Type:   envType,
			Region: "local",
			Status: status,
		})
	}

	c.JSON(http.StatusOK, environments)
}

func (s *Server) handleCreateEnvironment(c *gin.Context) {
	var environment models.Environment
	if err := c.ShouldBindJSON(&environment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.services.Environment.Create(&environment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, environment)
}

func (s *Server) handleGetEnvironment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid environment ID"})
		return
	}

	environment, err := s.services.Environment.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Environment not found"})
		return
	}

	c.JSON(http.StatusOK, environment)
}

func (s *Server) handleUpdateEnvironment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid environment ID"})
		return
	}

	var environment models.Environment
	if err := c.ShouldBindJSON(&environment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	environment.ID = uint(id)
	if err := s.services.Environment.Update(&environment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, environment)
}