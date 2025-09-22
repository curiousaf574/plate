package services

import (
	"fmt"
	"github.com/plate/service/internal/config"
	"gorm.io/gorm"
)

type Manager struct {
	config     *config.Config
	db         *gorm.DB
	Project    *ProjectService
	Deployment *DeploymentService
	Environment *EnvironmentService
	Kubernetes *KubernetesService
	ArgoCD     *ArgoCDService
	Helm       *HelmService
	Gitea      *GiteaService
}

func NewManager(cfg *config.Config, db *gorm.DB) *Manager {
	manager := &Manager{
		config: cfg,
		db:     db,
	}

	// Initialize services (skip database-dependent services for development)
	if db != nil {
		manager.Project = NewProjectService(db)
		manager.Environment = NewEnvironmentService(db)
		manager.Deployment = NewDeploymentService(db, nil, nil, nil, nil)
	}
	
	manager.Kubernetes = NewKubernetesService(cfg.Kubernetes)
	manager.ArgoCD = NewArgoCDService(cfg.ArgoCD)
	manager.Helm = NewHelmService(cfg.Helm)
	manager.Gitea = NewGiteaService(cfg.Gitea)

	return manager
}

func (m *Manager) Initialize() error {
	// For development, skip actual service initialization
	// TODO: Enable when implementing real integrations
	
	fmt.Println("Initializing Plate service manager...")
	fmt.Println("Running in development mode with fake data")
	
	return nil
}