package service

import (
	"path/filepath"
	"flag"
	"os"

	"fmt"
)

var configMap map[string]string = map[string]string{}

func getKubeConfig(configType string) *string{
	fmt.Println("getKubeconfig");
	if configMap[configType] != "" {
		config := configMap[configType]
		return &config;
	}
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "/src/b2c-deploy/web/config", "kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "../config", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	configMap[configType] = *kubeconfig;
	return kubeconfig;
}


func homeDir() string {
	if h := os.Getenv("GOPATH"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
