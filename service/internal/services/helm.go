package services

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/plate/service/internal/config"
	"github.com/plate/service/internal/models"
)

type HelmService struct {
	config config.Helm
}

func NewHelmService(cfg config.Helm) *HelmService {
	return &HelmService{
		config: cfg,
	}
}

func (s *HelmService) Initialize() error {
	// TODO: Initialize Helm client
	// This would typically involve:
	// 1. Setting up Helm configuration
	// 2. Adding chart repositories
	// 3. Updating repository cache
	
	fmt.Printf("Initializing Helm client\n")
	fmt.Printf("  Chart path: %s\n", s.config.ChartPath)
	
	// Create chart directory if it doesn't exist
	if err := os.MkdirAll(s.config.ChartPath, 0755); err != nil {
		return fmt.Errorf("failed to create chart directory: %w", err)
	}
	
	return nil
}

func (s *HelmService) GenerateChart(project *models.Project, environment *models.Environment) (string, error) {
	chartName := fmt.Sprintf("%s-%s", project.Name, environment.Name)
	chartDir := filepath.Join(s.config.ChartPath, chartName)
	
	// Create chart directory structure
	dirs := []string{
		chartDir,
		filepath.Join(chartDir, "templates"),
	}
	
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	
	// Generate Chart.yaml
	if err := s.generateChartYaml(chartDir, project); err != nil {
		return "", err
	}
	
	// Generate values.yaml
	if err := s.generateValues(chartDir, project, environment); err != nil {
		return "", err
	}
	
	// Generate deployment template
	if err := s.generateDeploymentTemplate(chartDir, project); err != nil {
		return "", err
	}
	
	// Generate service template
	if err := s.generateServiceTemplate(chartDir, project); err != nil {
		return "", err
	}
	
	// Generate ingress template if needed
	if environment.Domain != "" {
		if err := s.generateIngressTemplate(chartDir, project, environment); err != nil {
			return "", err
		}
	}
	
	return chartDir, nil
}

func (s *HelmService) InstallRelease(name, chartPath, namespace string) error {
	// TODO: Implement Helm release installation
	fmt.Printf("Installing Helm release: %s\n", name)
	fmt.Printf("  Chart: %s\n", chartPath)
	fmt.Printf("  Namespace: %s\n", namespace)
	
	// Placeholder implementation
	return nil
}

func (s *HelmService) UninstallRelease(name string) error {
	// TODO: Implement Helm release uninstallation
	fmt.Printf("Uninstalling Helm release: %s\n", name)
	return nil
}

func (s *HelmService) GetRelease(name string) (*HelmRelease, error) {
	// TODO: Implement Helm release retrieval
	return &HelmRelease{
		Name:      name,
		Status:    "deployed",
		Revision:  1,
		Namespace: "default",
	}, nil
}

func (s *HelmService) generateChartYaml(chartDir string, project *models.Project) error {
	chartYaml := `apiVersion: v2
name: {{.Name}}
description: {{.Description}}
type: application
version: 0.1.0
appVersion: "1.0.0"
`
	
	tmpl, err := template.New("chart").Parse(chartYaml)
	if err != nil {
		return err
	}
	
	file, err := os.Create(filepath.Join(chartDir, "Chart.yaml"))
	if err != nil {
		return err
	}
	defer file.Close()
	
	return tmpl.Execute(file, project)
}

func (s *HelmService) generateValues(chartDir string, project *models.Project, environment *models.Environment) error {
	valuesYaml := `replicaCount: 1

image:
  repository: {{.Name}}
  pullPolicy: IfNotPresent
  tag: "latest"

service:
  type: ClusterIP
  port: {{.Port}}

ingress:
  enabled: {{.IngressEnabled}}
  className: "nginx"
  annotations: {}
  hosts:
    - host: {{.Host}}
      paths:
        - path: /
          pathType: Prefix
  tls: []

resources:
  limits:
    cpu: 500m
    memory: 512Mi
  requests:
    cpu: 250m
    memory: 256Mi

env: {{.EnvVars}}
`
	
	data := struct {
		*models.Project
		IngressEnabled bool
		Host          string
		EnvVars       string
	}{
		Project:        project,
		IngressEnabled: environment.Domain != "",
		Host:          fmt.Sprintf("%s.%s", project.Name, environment.Domain),
		EnvVars:       project.EnvVars,
	}
	
	tmpl, err := template.New("values").Parse(valuesYaml)
	if err != nil {
		return err
	}
	
	file, err := os.Create(filepath.Join(chartDir, "values.yaml"))
	if err != nil {
		return err
	}
	defer file.Close()
	
	return tmpl.Execute(file, data)
}

func (s *HelmService) generateDeploymentTemplate(chartDir string, project *models.Project) error {
	deploymentTemplate := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "chart.fullname" . }}
  labels:
    {{- include "chart.labels" . | nindent 4 }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      {{- include "chart.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      labels:
        {{- include "chart.selectorLabels" . | nindent 8 }}
    spec:
      containers:
        - name: {{ .Chart.Name }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          ports:
            - name: http
              containerPort: {{ .Values.service.port }}
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /
              port: http
          readinessProbe:
            httpGet:
              path: /
              port: http
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
`
	
	return os.WriteFile(filepath.Join(chartDir, "templates", "deployment.yaml"), []byte(deploymentTemplate), 0644)
}

func (s *HelmService) generateServiceTemplate(chartDir string, project *models.Project) error {
	serviceTemplate := `apiVersion: v1
kind: Service
metadata:
  name: {{ include "chart.fullname" . }}
  labels:
    {{- include "chart.labels" . | nindent 4 }}
spec:
  type: {{ .Values.service.type }}
  ports:
    - port: {{ .Values.service.port }}
      targetPort: http
      protocol: TCP
      name: http
  selector:
    {{- include "chart.selectorLabels" . | nindent 4 }}
`
	
	return os.WriteFile(filepath.Join(chartDir, "templates", "service.yaml"), []byte(serviceTemplate), 0644)
}

func (s *HelmService) generateIngressTemplate(chartDir string, project *models.Project, environment *models.Environment) error {
	ingressTemplate := `{{- if .Values.ingress.enabled -}}
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: {{ include "chart.fullname" . }}
  labels:
    {{- include "chart.labels" . | nindent 4 }}
  {{- with .Values.ingress.annotations }}
  annotations:
    {{- toYaml . | nindent 4 }}
  {{- end }}
spec:
  {{- if .Values.ingress.className }}
  ingressClassName: {{ .Values.ingress.className }}
  {{- end }}
  {{- if .Values.ingress.tls }}
  tls:
    {{- range .Values.ingress.tls }}
    - hosts:
        {{- range .hosts }}
        - {{ . | quote }}
        {{- end }}
      secretName: {{ .secretName }}
    {{- end }}
  {{- end }}
  rules:
    {{- range .Values.ingress.hosts }}
    - host: {{ .host | quote }}
      http:
        paths:
          {{- range .paths }}
          - path: {{ .path }}
            {{- if .pathType }}
            pathType: {{ .pathType }}
            {{- end }}
            backend:
              service:
                name: {{ include "chart.fullname" $ }}
                port:
                  number: {{ $.Values.service.port }}
          {{- end }}
    {{- end }}
{{- end }}
`
	
	return os.WriteFile(filepath.Join(chartDir, "templates", "ingress.yaml"), []byte(ingressTemplate), 0644)
}

// HelmRelease represents a Helm release
type HelmRelease struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	Revision  int    `json:"revision"`
	Namespace string `json:"namespace"`
	Chart     string `json:"chart"`
}