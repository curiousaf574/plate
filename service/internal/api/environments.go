package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plate/service/internal/models"
)

func (s *Server) handleListEnvironments(c *gin.Context) {
	// Return developer-friendly fake data
	environments := []EnvironmentResponse{
		{
			ID:     1,
			Name:   "development",
			Type:   "development",
			Region: "us-west-2",
			Status: "active",
		},
		{
			ID:     2,
			Name:   "staging",
			Type:   "staging",
			Region: "us-west-2",
			Status: "active",
		},
		{
			ID:     3,
			Name:   "production",
			Type:   "production",
			Region: "us-west-2",
			Status: "active",
		},
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