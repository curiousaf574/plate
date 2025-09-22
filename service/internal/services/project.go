package services

import (
	"github.com/plate/service/internal/models"
	"gorm.io/gorm"
)

type ProjectService struct {
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{db: db}
}

func (s *ProjectService) List() ([]models.Project, error) {
	var projects []models.Project
	err := s.db.Preload("Deployments").Find(&projects).Error
	return projects, err
}

func (s *ProjectService) GetByID(id uint) (*models.Project, error) {
	var project models.Project
	err := s.db.Preload("Deployments").First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (s *ProjectService) GetByName(name string) (*models.Project, error) {
	var project models.Project
	err := s.db.Where("name = ?", name).First(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (s *ProjectService) Create(project *models.Project) error {
	return s.db.Create(project).Error
}

func (s *ProjectService) Update(project *models.Project) error {
	return s.db.Save(project).Error
}

func (s *ProjectService) Delete(id uint) error {
	return s.db.Delete(&models.Project{}, id).Error
}

type EnvironmentService struct {
	db *gorm.DB
}

func NewEnvironmentService(db *gorm.DB) *EnvironmentService {
	return &EnvironmentService{db: db}
}

func (s *EnvironmentService) List() ([]models.Environment, error) {
	var environments []models.Environment
	err := s.db.Find(&environments).Error
	return environments, err
}

func (s *EnvironmentService) GetByID(id uint) (*models.Environment, error) {
	var environment models.Environment
	err := s.db.First(&environment, id).Error
	if err != nil {
		return nil, err
	}
	return &environment, nil
}

func (s *EnvironmentService) GetByName(name string) (*models.Environment, error) {
	var environment models.Environment
	err := s.db.Where("name = ?", name).First(&environment).Error
	if err != nil {
		return nil, err
	}
	return &environment, nil
}

func (s *EnvironmentService) Create(environment *models.Environment) error {
	return s.db.Create(environment).Error
}

func (s *EnvironmentService) Update(environment *models.Environment) error {
	return s.db.Save(environment).Error
}

func (s *EnvironmentService) Delete(id uint) error {
	return s.db.Delete(&models.Environment{}, id).Error
}