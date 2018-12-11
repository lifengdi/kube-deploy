package service

import (
	"path/filepath"
	"flag"
	"os"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	restclient "k8s.io/client-go/rest"

	"k8s.io/client-go/kubernetes"
)

var configMap map[string]string = map[string]string{}

/**
	获取clientset
 */
func getClientset(configType string)(*kubernetes.Clientset, error){
	config,err:=getKubeClient(configType)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)

	return clientset,err;
}


/**
 	获取kubeconfig
 */
func getKubeClient(configType string)(*restclient.Config, error){

	var kubeconfig *string = getKubeConfig(configType)

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	return config,err;
}

/**
	根据集群类型获取kubeconfig
 */
func getKubeConfig(configType string) *string{
	fmt.Println("getKubeconfig");
	if configMap[configType] != "" {
		config := configMap[configType]
		return &config;
	}
	var kubeconfig *string
	if home := homeDir(); home != "" {
		configName := configType
		configName = configName + "-kubeconfig"
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "/src/b2c-deploy/web/resource", configName), "(optional) absolute path to the taoche-test-kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "../config", "absolute path to the taoche-test-kubeconfig file")
	}
	flag.Parse()
	configMap[configType] = *kubeconfig;
	return kubeconfig;
}

/**
 获取gopath路径
 */
func homeDir() string {
	if h := os.Getenv("GOPATH"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
