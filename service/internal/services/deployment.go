package services

import (
	"fmt"
	"time"

	"github.com/plate/service/internal/models"
	"gorm.io/gorm"
)

type DeploymentService struct {
	db         *gorm.DB
	kubernetes *KubernetesService
	argocd     *ArgoCDService
	helm       *HelmService
	gitea      *GiteaService
}

func NewDeploymentService(db *gorm.DB, k8s *KubernetesService, argo *ArgoCDService, helm *HelmService, gitea *GiteaService) *DeploymentService {
	return &DeploymentService{
		db:         db,
		kubernetes: k8s,
		argocd:     argo,
		helm:       helm,
		gitea:      gitea,
	}
}

func (s *DeploymentService) List() ([]models.Deployment, error) {
	var deployments []models.Deployment
	err := s.db.Preload("Project").Preload("Environment").Find(&deployments).Error
	return deployments, err
}

func (s *DeploymentService) GetByID(id uint) (*models.Deployment, error) {
	var deployment models.Deployment
	err := s.db.Preload("Project").Preload("Environment").First(&deployment, id).Error
	if err != nil {
		return nil, err
	}
	return &deployment, nil
}

func (s *DeploymentService) Create(deployment *models.Deployment) error {
	deployment.Status = "pending"
	return s.db.Create(deployment).Error
}

func (s *DeploymentService) Update(deployment *models.Deployment) error {
	return s.db.Save(deployment).Error
}

func (s *DeploymentService) Delete(id uint) error {
	// Get deployment to clean up resources
	deployment, err := s.GetByID(id)
	if err != nil {
		return err
	}

	// Clean up ArgoCD application
	if deployment.ArgoAppName != "" {
		if err := s.argocd.DeleteApplication(deployment.ArgoAppName); err != nil {
			// Log error but don't fail the deletion
			fmt.Printf("Warning: Failed to delete ArgoCD application %s: %v\n", deployment.ArgoAppName, err)
		}
	}

	// Clean up Helm release
	if deployment.HelmRelease != "" {
		if err := s.helm.UninstallRelease(deployment.HelmRelease); err != nil {
			// Log error but don't fail the deletion
			fmt.Printf("Warning: Failed to uninstall Helm release %s: %v\n", deployment.HelmRelease, err)
		}
	}

	return s.db.Delete(&models.Deployment{}, id).Error
}

func (s *DeploymentService) Deploy(projectID, environmentID uint, version string) (*models.Deployment, error) {
	// Get project and environment
	var project models.Project
	if err := s.db.First(&project, projectID).Error; err != nil {
		return nil, fmt.Errorf("project not found: %w", err)
	}

	var environment models.Environment
	if err := s.db.First(&environment, environmentID).Error; err != nil {
		return nil, fmt.Errorf("environment not found: %w", err)
	}

	// Set default version if not provided
	if version == "" {
		version = fmt.Sprintf("v%d", time.Now().Unix())
	}

	// Create deployment record
	deployment := &models.Deployment{
		ProjectID:     projectID,
		EnvironmentID: environmentID,
		Version:       version,
		Status:        "pending",
		ArgoAppName:   fmt.Sprintf("%s-%s", project.Name, environment.Name),
		HelmRelease:   fmt.Sprintf("%s-%s", project.Name, environment.Name),
	}

	if err := s.Create(deployment); err != nil {
		return nil, fmt.Errorf("failed to create deployment record: %w", err)
	}

	// Start deployment process asynchronously
	go s.performDeployment(deployment, &project, &environment)

	return deployment, nil
}

func (s *DeploymentService) performDeployment(deployment *models.Deployment, project *models.Project, environment *models.Environment) {
	// Update status to running
	deployment.Status = "running"
	s.db.Save(deployment)

	// Create repository if it doesn't exist
	if err := s.gitea.CreateRepository(project.Name, project.Description); err != nil {
		s.handleDeploymentError(deployment, fmt.Errorf("failed to create repository: %w", err))
		return
	}

	// Generate Helm chart
	chartPath, err := s.helm.GenerateChart(project, environment)
	if err != nil {
		s.handleDeploymentError(deployment, fmt.Errorf("failed to generate Helm chart: %w", err))
		return
	}

	// Create ArgoCD application
	if err := s.argocd.CreateApplication(deployment.ArgoAppName, project.Name, environment.Namespace, chartPath); err != nil {
		s.handleDeploymentError(deployment, fmt.Errorf("failed to create ArgoCD application: %w", err))
		return
	}

	// Deploy with Helm
	if err := s.helm.InstallRelease(deployment.HelmRelease, chartPath, environment.Namespace); err != nil {
		s.handleDeploymentError(deployment, fmt.Errorf("failed to install Helm release: %w", err))
		return
	}

	// Update deployment status and URL
	deployment.Status = "success"
	if environment.Domain != "" {
		deployment.URL = fmt.Sprintf("https://%s.%s", project.Name, environment.Domain)
	}
	s.db.Save(deployment)

	// Log successful deployment
	s.logDeployment(deployment.ID, "info", "Deployment completed successfully")
}

func (s *DeploymentService) handleDeploymentError(deployment *models.Deployment, err error) {
	deployment.Status = "failed"
	s.db.Save(deployment)
	s.logDeployment(deployment.ID, "error", err.Error())
}

func (s *DeploymentService) logDeployment(deploymentID uint, level, message string) {
	log := &models.DeploymentLog{
		DeploymentID: deploymentID,
		Level:        level,
		Message:      message,
		Timestamp:    time.Now(),
	}
	s.db.Create(log)
}

func (s *DeploymentService) GetLogs(deploymentID uint) ([]models.DeploymentLog, error) {
	var logs []models.DeploymentLog
	err := s.db.Where("deployment_id = ?", deploymentID).Order("timestamp desc").Find(&logs).Error
	return logs, err
}

func (s *DeploymentService) GetStatus(environment string, detailed bool) (interface{}, error) {
	query := s.db.Preload("Project").Preload("Environment")
	
	if environment != "" {
		query = query.Joins("JOIN environments ON deployments.environment_id = environments.id").
			Where("environments.name = ?", environment)
	}

	var deployments []models.Deployment
	if err := query.Find(&deployments).Error; err != nil {
		return nil, err
	}

	if !detailed {
		// Return simple status
		result := make(map[string]interface{})
		for _, d := range deployments {
			result[fmt.Sprintf("%s-%s", d.Project.Name, d.Environment.Name)] = d.Status
		}
		return result, nil
	}

	// Return detailed status
	return deployments, nil
}