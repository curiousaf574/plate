package services

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/plate/service/internal/config"
	corev1 "k8s.io/api/core/v1"
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