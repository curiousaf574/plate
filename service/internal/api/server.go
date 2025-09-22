package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plate/service/internal/config"
	"github.com/plate/service/internal/services"
)

type Server struct {
	config  *config.Config
	services *services.Manager
	router  *gin.Engine
}

func NewServer(cfg *config.Config, serviceManager *services.Manager) *Server {
	server := &Server{
		config:   cfg,
		services: serviceManager,
		router:   gin.Default(),
	}

	server.setupRoutes()
	return server
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) setupRoutes() {
	// CORS middleware
	s.router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Health check
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// API v1 routes
	v1 := s.router.Group("/api/v1")
	{
		// Projects
		projects := v1.Group("/projects")
		{
			projects.GET("", s.handleListProjects)
			projects.POST("", s.handleCreateProject)
			projects.GET("/:id", s.handleGetProject)
			projects.PUT("/:id", s.handleUpdateProject)
			projects.DELETE("/:id", s.handleDeleteProject)
		}

		// Deployments
		deployments := v1.Group("/deployments")
		{
			deployments.GET("", s.handleListDeployments)
			deployments.POST("", s.handleCreateDeployment)
			deployments.GET("/:id", s.handleGetDeployment)
			deployments.DELETE("/:id", s.handleDeleteDeployment)
			deployments.GET("/:id/logs", s.handleGetDeploymentLogs)
		}

		// Environments
		environments := v1.Group("/environments")
		{
			environments.GET("", s.handleListEnvironments)
			environments.POST("", s.handleCreateEnvironment)
			environments.GET("/:id", s.handleGetEnvironment)
			environments.PUT("/:id", s.handleUpdateEnvironment)
		}

		// Deploy action
		v1.POST("/deploy", s.handleDeploy)
		
		// Status
		v1.GET("/status", s.handleGetStatus)
	}
}