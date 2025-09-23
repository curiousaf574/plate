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

// Project handlers
func (s *Server) handleListProjects(c *gin.Context) {
	// Get all namespaces managed by Plate
	namespaces, err := s.services.Kubernetes.GetNamespaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get namespaces"})
		return
	}

	// Filter namespaces managed by Plate
	plateNamespaces := []string{}
	for _, ns := range namespaces {
		if ns == "dev" || ns == "staging" || ns == "production" {
			plateNamespaces = append(plateNamespaces, ns)
		}
	}

	// Collect unique applications across all environments
	applicationMap := make(map[string]*ProjectResponse)

	for _, namespace := range plateNamespaces {
		deployments, err := s.services.Kubernetes.GetClientset().AppsV1().Deployments(namespace).List(c, metav1.ListOptions{
			LabelSelector: "managed-by=plate",
		})
		if err != nil {
			continue
		}

		for i, deployment := range deployments.Items {
			appName := deployment.Labels["app"]
			runtime := deployment.Labels["runtime"]
			if appName == "" {
				appName = deployment.Name
			}
			if runtime == "" {
				runtime = "unknown"
			}

			status := "active"
			if deployment.Status.ReadyReplicas == 0 {
				status = "inactive"
			} else if deployment.Status.ReadyReplicas < *deployment.Spec.Replicas {
				status = "building"
			}

			lastDeploy := "unknown"
			if deployment.Status.Conditions != nil && len(deployment.Status.Conditions) > 0 {
				lastUpdate := deployment.Status.Conditions[len(deployment.Status.Conditions)-1].LastUpdateTime
				lastDeploy = lastUpdate.Format("2006-01-02 15:04")
			}

			if existing, exists := applicationMap[appName]; exists {
				// Add environment to existing app
				existing.Environments = append(existing.Environments, namespace)
				// Update status if this environment has issues
				if status != "active" {
					existing.Status = status
				}
			} else {
				// Create new application entry
				applicationMap[appName] = &ProjectResponse{
					ID:          uint(i + 1),
					Name:        appName,
					Description: fmt.Sprintf("%s application running on Kubernetes", strings.Title(runtime)),
					Repository:  fmt.Sprintf("https://github.com/yourorg/%s.git", appName),
					Runtime:     runtime,
					Status:      status,
					LastDeploy:  lastDeploy,
					Environments: []string{namespace},
				}
			}
		}
	}

	// Convert map to slice
	var projects []ProjectResponse
	for _, project := range applicationMap {
		projects = append(projects, *project)
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