package config

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// GetConfigPath fetches the path to the users kubeconfig
func GetConfigPath() string {
	if configPath := os.Getenv("K6_CHAOS_KUBECONFIG"); configPath != "" {
		return configPath
	}

	return filepath.Join(homedir.HomeDir(), ".kube", "config")
}

// GetConfig creates a new k8s config
func GetConfig() *rest.Config {
	configPath := GetConfigPath()
	config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	return config
}
