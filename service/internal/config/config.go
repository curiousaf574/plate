package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port      string     `mapstructure:"port"`
	Database  Database   `mapstructure:"db"`
	Kubernetes Kubernetes `mapstructure:"kubernetes"`
	ArgoCD    ArgoCD     `mapstructure:"argocd"`
	Gitea     Gitea      `mapstructure:"gitea"`
	Helm      Helm       `mapstructure:"helm"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"sslmode"`
}

type Kubernetes struct {
	ConfigPath string `mapstructure:"config_path"`
	Namespace  string `mapstructure:"namespace"`
	InCluster  bool   `mapstructure:"in_cluster"`
}

type ArgoCD struct {
	Server   string `mapstructure:"server"`
	Token    string `mapstructure:"token"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Gitea struct {
	URL    string `mapstructure:"url"`
	Token  string `mapstructure:"token"`
	OrgName string `mapstructure:"org_name"`
}

type Helm struct {
	RepoURL   string `mapstructure:"repo_url"`
	ChartPath string `mapstructure:"chart_path"`
}

func Load() *Config {
	cfg := &Config{
		Port: viper.GetString("port"),
		Database: Database{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Name:     viper.GetString("db.name"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
		Kubernetes: Kubernetes{
			ConfigPath: viper.GetString("kubernetes.config_path"),
			Namespace:  viper.GetString("kubernetes.namespace"),
			InCluster:  viper.GetBool("kubernetes.in_cluster"),
		},
		ArgoCD: ArgoCD{
			Server:   viper.GetString("argocd.server"),
			Token:    viper.GetString("argocd.token"),
			Username: viper.GetString("argocd.username"),
			Password: viper.GetString("argocd.password"),
		},
		Gitea: Gitea{
			URL:     viper.GetString("gitea.url"),
			Token:   viper.GetString("gitea.token"),
			OrgName: viper.GetString("gitea.org_name"),
		},
		Helm: Helm{
			RepoURL:   viper.GetString("helm.repo_url"),
			ChartPath: viper.GetString("helm.chart_path"),
		},
	}

	// Set defaults
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.Database.SSLMode == "" {
		cfg.Database.SSLMode = "disable"
	}
	if cfg.Kubernetes.Namespace == "" {
		cfg.Kubernetes.Namespace = "plate-system"
	}

	return cfg
}