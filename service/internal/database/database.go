package database

import (
	"fmt"

	"github.com/plate/service/internal/config"
	"github.com/plate/service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Initialize(cfg config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate models
	err = db.AutoMigrate(
		&models.Project{},
		&models.Environment{},
		&models.Deployment{},
		&models.DeploymentLog{},
		&models.Repository{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	// Create default environments if they don't exist
	environments := []models.Environment{
		{Name: "development", Namespace: "plate-dev", Domain: "dev.plate.local"},
		{Name: "staging", Namespace: "plate-staging", Domain: "staging.plate.local"},
		{Name: "production", Namespace: "plate-prod", Domain: "plate.local"},
	}

	for _, env := range environments {
		var existing models.Environment
		if err := db.Where("name = ?", env.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&env).Error; err != nil {
				return nil, fmt.Errorf("failed to create environment %s: %w", env.Name, err)
			}
		}
	}

	return db, nil
}