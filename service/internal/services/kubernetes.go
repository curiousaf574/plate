package services

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/plate/service/internal/config"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type KubernetesService struct {
	config    config.Kubernetes
	clientset *kubernetes.Clientset
	restConfig *rest.Config
}

func NewKubernetesService(cfg config.Kubernetes) *KubernetesService {
	return &KubernetesService{
		config: cfg,
	}
}

func (s *KubernetesService) Initialize() error {
	var config *rest.Config
	var err error

	if s.config.InCluster {
		// Use in-cluster configuration
		config, err = rest.InClusterConfig()
	} else {
		// Use kubeconfig file
		configPath := s.config.ConfigPath
		if configPath == "" {
			if home := homedir.HomeDir(); home != "" {
				configPath = filepath.Join(home, ".kube", "config")
			}
		}
		config, err = clientcmd.BuildConfigFromFlags("", configPath)
	}

	if err != nil {
		return fmt.Errorf("failed to create Kubernetes config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	s.clientset = clientset
	s.restConfig = config

	// Test connection
	_, err = s.clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{Limit: 1})
	if err != nil {
		return fmt.Errorf("failed to connect to Kubernetes cluster: %w", err)
	}

	return nil
}

func (s *KubernetesService) GetClientset() *kubernetes.Clientset {
	return s.clientset
}

func (s *KubernetesService) GetNamespaces() ([]string, error) {
	namespaces, err := s.clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var names []string
	for _, ns := range namespaces.Items {
		names = append(names, ns.Name)
	}
	return names, nil
}

func (s *KubernetesService) CreateNamespace(name string) error {
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	_, err := s.clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("failed to create namespace %s: %w", name, err)
	}

	return nil
}

func (s *KubernetesService) GetPods(namespace string) ([]corev1.Pod, error) {
	pods, err := s.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods.Items, nil
}

func (s *KubernetesService) GetServices(namespace string) ([]corev1.Service, error) {
	services, err := s.clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return services.Items, nil
}

// Deployment management methods
func (s *KubernetesService) ScaleDeployment(namespace, name string, replicas int32) error {
	deployment, err := s.clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment %s/%s: %w", namespace, name, err)
	}

	deployment.Spec.Replicas = &replicas
	_, err = s.clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to scale deployment %s/%s to %d replicas: %w", namespace, name, replicas, err)
	}

	return nil
}

func (s *KubernetesService) StopDeployment(namespace, name string) error {
	return s.ScaleDeployment(namespace, name, 0)
}

func (s *KubernetesService) StartDeployment(namespace, name string, replicas int32) error {
	if replicas <= 0 {
		replicas = 1 // Default to 1 replica if not specified
	}
	return s.ScaleDeployment(namespace, name, replicas)
}

func (s *KubernetesService) RestartDeployment(namespace, name string) error {
	deployment, err := s.clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment %s/%s: %w", namespace, name, err)
	}

	// Add or update restart annotation to trigger restart
	if deployment.Spec.Template.Annotations == nil {
		deployment.Spec.Template.Annotations = make(map[string]string)
	}
	deployment.Spec.Template.Annotations["kubectl.kubernetes.io/restartedAt"] = metav1.Now().Format("2006-01-02T15:04:05Z")

	_, err = s.clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to restart deployment %s/%s: %w", namespace, name, err)
	}

	return nil
}

func (s *KubernetesService) GetDeploymentStatus(namespace, name string) (*DeploymentStatus, error) {
	deployment, err := s.clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment %s/%s: %w", namespace, name, err)
	}

	pods, err := s.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", name),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get pods for deployment %s/%s: %w", namespace, name, err)
	}

	// Get ingress routes for this deployment
	routes, err := s.GetIngressRoutes(namespace, name)
	if err != nil {
		// Don't fail if ingress lookup fails, just log and continue
		routes = []IngressRoute{}
	}

	status := &DeploymentStatus{
		Name:              deployment.Name,
		Namespace:         deployment.Namespace,
		DesiredReplicas:   *deployment.Spec.Replicas,
		ReadyReplicas:     deployment.Status.ReadyReplicas,
		AvailableReplicas: deployment.Status.AvailableReplicas,
		Conditions:        deployment.Status.Conditions,
		Pods:              make([]PodStatus, 0, len(pods.Items)),
		Routes:            routes,
	}

	for _, pod := range pods.Items {
		podStatus := PodStatus{
			Name:    pod.Name,
			Phase:   string(pod.Status.Phase),
			Ready:   "0/0",
			Restarts: 0,
		}

		// Calculate ready containers
		readyCount := 0
		totalCount := len(pod.Status.ContainerStatuses)
		for _, container := range pod.Status.ContainerStatuses {
			if container.Ready {
				readyCount++
			}
			podStatus.Restarts += container.RestartCount
		}
		podStatus.Ready = fmt.Sprintf("%d/%d", readyCount, totalCount)

		status.Pods = append(status.Pods, podStatus)
	}

	return status, nil
}

// GetIngressRoutes retrieves ingress routes for a specific deployment
func (s *KubernetesService) GetIngressRoutes(namespace, deploymentName string) ([]IngressRoute, error) {
	// Get all ingresses in the namespace
	ingresses, err := s.clientset.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get ingresses in namespace %s: %w", namespace, err)
	}

	var routes []IngressRoute

	// Look for ingresses that reference services related to this deployment
	for _, ingress := range ingresses.Items {
		// Check if ingress is related to this deployment by labels or service names
		if s.isIngressRelatedToDeployment(ingress, deploymentName) {
			for _, rule := range ingress.Spec.Rules {
				host := rule.Host
				if host == "" {
					host = "localhost" // Default for cluster-internal access
				}

				if rule.HTTP != nil {
					for _, path := range rule.HTTP.Paths {
						route := IngressRoute{
							Host:        host,
							Path:        path.Path,
							ServiceName: path.Backend.Service.Name,
							ServicePort: int32(path.Backend.Service.Port.Number),
							PathType:    string(*path.PathType),
							URL:         fmt.Sprintf("http://%s%s", host, path.Path),
						}
						routes = append(routes, route)
					}
				}
			}
		}
	}

	return routes, nil
}

// isIngressRelatedToDeployment checks if an ingress is related to a deployment
func (s *KubernetesService) isIngressRelatedToDeployment(ingress networkingv1.Ingress, deploymentName string) bool {
	// Check labels
	if appLabel, exists := ingress.Labels["app"]; exists && appLabel == deploymentName {
		return true
	}

	// Check if any service in the ingress matches the deployment pattern
	for _, rule := range ingress.Spec.Rules {
		if rule.HTTP != nil {
			for _, path := range rule.HTTP.Paths {
				serviceName := path.Backend.Service.Name
				// Common patterns: app-service, app-svc, or just app
				if serviceName == deploymentName ||
					serviceName == deploymentName+"-service" ||
					serviceName == deploymentName+"-svc" ||
					strings.Contains(serviceName, deploymentName) {
					return true
				}
			}
		}
	}

	return false
}

type DeploymentStatus struct {
	Name              string                     `json:"name"`
	Namespace         string                     `json:"namespace"`
	DesiredReplicas   int32                      `json:"desired_replicas"`
	ReadyReplicas     int32                      `json:"ready_replicas"`
	AvailableReplicas int32                      `json:"available_replicas"`
	Conditions        []appsv1.DeploymentCondition `json:"conditions"`
	Pods              []PodStatus                `json:"pods"`
	Routes            []IngressRoute             `json:"routes"`
}

type PodStatus struct {
	Name     string `json:"name"`
	Phase    string `json:"phase"`
	Ready    string `json:"ready"`
	Restarts int32  `json:"restarts"`
}

type IngressRoute struct {
	Host        string `json:"host"`
	Path        string `json:"path"`
	ServiceName string `json:"service_name"`
	ServicePort int32  `json:"service_port"`
	PathType    string `json:"path_type"`
	URL         string `json:"url"`
}