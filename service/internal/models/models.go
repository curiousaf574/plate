package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description"`
	Repository  string    `json:"repository"`
	Runtime     string    `json:"runtime"`
	BuildCmd    string    `json:"build_cmd"`
	StartCmd    string    `json:"start_cmd"`
	Port        int       `json:"port"`
	EnvVars     string    `json:"env_vars" gorm:"type:text"` // JSON string
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	Deployments []Deployment `json:"deployments,omitempty" gorm:"foreignKey:ProjectID"`
}

type Environment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"uniqueIndex;not null"`
	Namespace string    `json:"namespace"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	Deployments []Deployment `json:"deployments,omitempty" gorm:"foreignKey:EnvironmentID"`
}

type Deployment struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	ProjectID     uint        `json:"project_id"`
	EnvironmentID uint        `json:"environment_id"`
	Version       string      `json:"version"`
	Status        string      `json:"status"` // pending, running, success, failed
	URL           string      `json:"url"`
	LogURL        string      `json:"log_url"`
	ArgoAppName   string      `json:"argo_app_name"`
	HelmRelease   string      `json:"helm_release"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	
	Project     Project     `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
	Environment Environment `json:"environment,omitempty" gorm:"foreignKey:EnvironmentID"`
}

type DeploymentLog struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	DeploymentID uint      `json:"deployment_id"`
	Level        string    `json:"level"` // info, warning, error
	Message      string    `json:"message" gorm:"type:text"`
	Timestamp    time.Time `json:"timestamp"`
	
	Deployment Deployment `json:"deployment,omitempty" gorm:"foreignKey:DeploymentID"`
}

type Repository struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ProjectID   uint      `json:"project_id"`
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	CloneURL    string    `json:"clone_url"`
	Branch      string    `json:"branch"`
	GiteaRepoID int       `json:"gitea_repo_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	Project Project `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
}