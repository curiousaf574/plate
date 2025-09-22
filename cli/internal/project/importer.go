package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Importer struct {
	// Add any dependencies here
}

type ProjectConfig struct {
	Name        string            `yaml:"name"`
	Runtime     string            `yaml:"runtime"`
	Environment string            `yaml:"environment"`
	Port        int               `yaml:"port"`
	BuildCmd    string            `yaml:"build_cmd,omitempty"`
	StartCmd    string            `yaml:"start_cmd,omitempty"`
	EnvVars     map[string]string `yaml:"env_vars,omitempty"`
}

func NewImporter() *Importer {
	return &Importer{}
}

func (i *Importer) Import(projectPath, name, env, runtime string) error {
	// Detect project name if not provided
	if name == "" {
		name = filepath.Base(projectPath)
	}

	// Auto-detect runtime if not provided
	if runtime == "" {
		detectedRuntime, err := i.detectRuntime(projectPath)
		if err != nil {
			return fmt.Errorf("failed to detect runtime: %w", err)
		}
		runtime = detectedRuntime
	}

	// Create project configuration
	config := ProjectConfig{
		Name:        name,
		Runtime:     runtime,
		Environment: env,
		Port:        i.getDefaultPort(runtime),
		EnvVars:     make(map[string]string),
	}

	// Set runtime-specific commands
	i.setRuntimeCommands(&config)

	// Create .plate directory
	plateDir := filepath.Join(projectPath, ".plate")
	if err := os.MkdirAll(plateDir, 0755); err != nil {
		return fmt.Errorf("failed to create .plate directory: %w", err)
	}

	// Write configuration file
	configPath := filepath.Join(plateDir, "config.yaml")
	if err := i.writeConfig(configPath, &config); err != nil {
		return fmt.Errorf("failed to write configuration: %w", err)
	}

	// Generate Dockerfile if it doesn't exist
	dockerfilePath := filepath.Join(projectPath, "Dockerfile")
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		if err := i.generateDockerfile(dockerfilePath, runtime); err != nil {
			return fmt.Errorf("failed to generate Dockerfile: %w", err)
		}
	}

	fmt.Printf("Project configured successfully:\n")
	fmt.Printf("  Name: %s\n", config.Name)
	fmt.Printf("  Runtime: %s\n", config.Runtime)
	fmt.Printf("  Environment: %s\n", config.Environment)
	fmt.Printf("  Port: %d\n", config.Port)

	return nil
}

func (i *Importer) detectRuntime(projectPath string) (string, error) {
	// Check for various project files to detect runtime
	checks := map[string]string{
		"package.json":   "nodejs",
		"requirements.txt": "python",
		"Pipfile":        "python",
		"go.mod":         "go",
		"Cargo.toml":     "rust",
		"pom.xml":        "java",
		"build.gradle":   "java",
		"composer.json":  "php",
		"Gemfile":        "ruby",
	}

	for file, runtime := range checks {
		if _, err := os.Stat(filepath.Join(projectPath, file)); err == nil {
			return runtime, nil
		}
	}

	return "generic", nil
}

func (i *Importer) getDefaultPort(runtime string) int {
	ports := map[string]int{
		"nodejs": 3000,
		"python": 8000,
		"go":     8080,
		"rust":   8080,
		"java":   8080,
		"php":    8000,
		"ruby":   3000,
	}

	if port, exists := ports[runtime]; exists {
		return port
	}
	return 8080
}

func (i *Importer) setRuntimeCommands(config *ProjectConfig) {
	switch config.Runtime {
	case "nodejs":
		config.BuildCmd = "npm install"
		config.StartCmd = "npm start"
	case "python":
		config.BuildCmd = "pip install -r requirements.txt"
		config.StartCmd = "python app.py"
	case "go":
		config.BuildCmd = "go build -o app ."
		config.StartCmd = "./app"
	case "rust":
		config.BuildCmd = "cargo build --release"
		config.StartCmd = "./target/release/app"
	case "java":
		config.BuildCmd = "mvn clean package"
		config.StartCmd = "java -jar target/*.jar"
	case "php":
		config.StartCmd = "php -S 0.0.0.0:8000"
	case "ruby":
		config.BuildCmd = "bundle install"
		config.StartCmd = "bundle exec ruby app.rb"
	}
}

func (i *Importer) writeConfig(configPath string, config *ProjectConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func (i *Importer) generateDockerfile(dockerfilePath, runtime string) error {
	var dockerfile string

	switch runtime {
	case "nodejs":
		dockerfile = `FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD ["npm", "start"]`

	case "python":
		dockerfile = `FROM python:3.9-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
EXPOSE 8000
CMD ["python", "app.py"]`

	case "go":
		dockerfile = `FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]`

	case "rust":
		dockerfile = `FROM rust:1.70 AS builder
WORKDIR /app
COPY Cargo.toml Cargo.lock ./
COPY src ./src
RUN cargo build --release

FROM debian:bookworm-slim
COPY --from=builder /app/target/release/app /usr/local/bin/app
EXPOSE 8080
CMD ["app"]`

	case "java":
		dockerfile = `FROM openjdk:17-jdk-slim AS builder
WORKDIR /app
COPY pom.xml .
COPY src ./src
RUN ./mvnw clean package

FROM openjdk:17-jre-slim
COPY --from=builder /app/target/*.jar app.jar
EXPOSE 8080
CMD ["java", "-jar", "app.jar"]`

	default:
		dockerfile = `FROM alpine:latest
WORKDIR /app
COPY . .
EXPOSE 8080
CMD ["./start.sh"]`
	}

	return os.WriteFile(dockerfilePath, []byte(strings.TrimSpace(dockerfile)+"\n"), 0644)
}