package service

import (
	"b2c-deploy/web/reqBody"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
)

func Create(request reqBody.CreateRequest)(string,error){
	println(request.Image)

	var kubeconfig *string = getKubeConfig("")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	err = createDeployment(clientset,request);
	if err != nil {
		panic(err.Error())
	}
	err = createService(clientset,request);
	if err != nil {
		panic(err.Error())
	}


	return "ok",nil
}



func int32Ptr2(i int32) *int32 { return &i }


