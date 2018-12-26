package service

import (
	"os"
	"k8s.io/client-go/tools/clientcmd"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	"kube-deploy/web/config"
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
	if configMap[configType] != "" {
		config := configMap[configType]
		return &config;
	}
	//var kubeconfig *string
	configName := configType
	configName = configName + "-kubeconfig"
	//kubeconfig = flag.String("kubeconfig", filepath.Join(home, "/src/kube-deploy/web/resource", configName), "(optional) absolute path to the taoche-test-kubeconfig file")
	//kubeconfig = flag.String("kubeconfig", filepath.Join(config.Get("kubeconfs"), configName), "(optional) absolute path to the taoche-test-kubeconfig file")
	//println(kubeconfig)
	//flag.Parse()
	kubeconfig := config.Get("kubeConfs")+configName
	configMap[configType] = kubeconfig;
	println("--------------"+kubeconfig)
	return &kubeconfig;
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
