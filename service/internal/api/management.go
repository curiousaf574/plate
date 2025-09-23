package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Deployment management handlers
func (s *Server) handleScaleDeployment(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	
	var req struct {
		Replicas int32 `json:"replicas" binding:"required,min=0"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.services.Kubernetes.ScaleDeployment(namespace, name, req.Replicas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deployment scaled successfully"})
}

func (s *Server) handleStopDeployment(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	if err := s.services.Kubernetes.StopDeployment(namespace, name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deployment stopped successfully"})
}

func (s *Server) handleStartDeployment(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	
	var req struct {
		Replicas int32 `json:"replicas"`
	}
	
	// Optional body for specifying replicas
	c.ShouldBindJSON(&req)
	if req.Replicas <= 0 {
		req.Replicas = 1
	}

	if err := s.services.Kubernetes.StartDeployment(namespace, name, req.Replicas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deployment started successfully"})
}

func (s *Server) handleRestartDeployment(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	if err := s.services.Kubernetes.RestartDeployment(namespace, name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deployment restarted successfully"})
}

func (s *Server) handleGetDeploymentStatus(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	status, err := s.services.Kubernetes.GetDeploymentStatus(namespace, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, status)
}

// App-specific handlers (works across environments)
func (s *Server) handleGetAppStatus(c *gin.Context) {
	appName := c.Param("name")
	
	// Check all plate-managed namespaces
	namespaces := []string{"dev", "staging", "production"}
	appStatus := make(map[string]interface{})

	for _, namespace := range namespaces {
		status, err := s.services.Kubernetes.GetDeploymentStatus(namespace, appName)
		if err != nil {
			// Skip if deployment doesn't exist in this namespace
			continue
		}
		appStatus[namespace] = status
	}

	if len(appStatus) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found in any environment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"application": appName,
		"environments": appStatus,
	})
}

func (s *Server) handleScaleApp(c *gin.Context) {
	appName := c.Param("name")
	environment := c.Query("env")
	
	if environment == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Environment parameter 'env' is required"})
		return
	}

	var req struct {
		Replicas int32 `json:"replicas" binding:"required,min=0"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.services.Kubernetes.ScaleDeployment(environment, appName, req.Replicas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application scaled successfully"})
}

func (s *Server) handleStopApp(c *gin.Context) {
	appName := c.Param("name")
	environment := c.Query("env")
	
	if environment == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Environment parameter 'env' is required"})
		return
	}

	if err := s.services.Kubernetes.StopDeployment(environment, appName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application stopped successfully"})
}

func (s *Server) handleStartApp(c *gin.Context) {
	appName := c.Param("name")
	environment := c.Query("env")
	
	if environment == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Environment parameter 'env' is required"})
		return
	}

	var req struct {
		Replicas int32 `json:"replicas"`
	}
	
	// Optional body for specifying replicas
	c.ShouldBindJSON(&req)
	if req.Replicas <= 0 {
		req.Replicas = 1
	}

	if err := s.services.Kubernetes.StartDeployment(environment, appName, req.Replicas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application started successfully"})
}

func (s *Server) handleRestartApp(c *gin.Context) {
	appName := c.Param("name")
	environment := c.Query("env")
	
	if environment == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Environment parameter 'env' is required"})
		return
	}

	if err := s.services.Kubernetes.RestartDeployment(environment, appName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application restarted successfully"})
}