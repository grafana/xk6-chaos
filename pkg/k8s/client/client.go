package client

import (
	"github.com/simskij/k6-extension-chaos/pkg/k8s/config"
	"k8s.io/client-go/kubernetes"
)

// New creates a new k8s client
func New() (*kubernetes.Clientset, error) {
	config := config.GetConfig()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
