package service

import (
	"fmt"
	"kube-deploy/web/reqBody"
	"k8s.io/client-go/kubernetes"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
	"kube-deploy/web/config"
	"encoding/json"
	"time"
	"strings"
)
func createDeployment(clientset *kubernetes.Clientset,request reqBody.ServiceRequest) error {


	deploymentsClient := clientset.AppsV1beta1().Deployments(request.Namespace)



	fmt.Println("Creating deployment...")
	_, err := deploymentsClient.Create(getDeployment(request))

	return err;
}


func deleteDeployment(clientset *kubernetes.Clientset,request reqBody.ServiceRequest) error {
	println("namespace:",request.Namespace)
	deploymentsClient := clientset.AppsV1beta1().Deployments(request.Namespace)

	deletePolicy := metav1.DeletePropagationForeground

	err := deploymentsClient.Delete(request.ServiceName,&metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	});

	return err;
}

func updateDeployment(clientset *kubernetes.Clientset,request reqBody.ServiceRequest) error {
	deploymentsClient := clientset.AppsV1beta1().Deployments(request.Namespace)
	_,err := deploymentsClient.Update(getDeployment(request))
	return err;
}

func restartDeployment(clientset *kubernetes.Clientset,request reqBody.ServiceRequest) error {
	deploymentsClient := clientset.AppsV1beta1().Deployments(request.Namespace)
	//_,err := deploymentsClient.Update(getDeployment(request))
	//metav1.
	println(request.ServiceName)
	deployment,err := deploymentsClient.Get(request.ServiceName,metav1.GetOptions{})
	if err!=nil{
		return err
	}
	/*
	labels := deployment.ObjectMeta.Labels
	labels["timestamp"] = time.Now().Format("2006-01-02-15-04-05")//1月2日3时4分5秒6年
	*/
	podLabels :=deployment.Spec.Template.ObjectMeta.Labels
	podLabels["timestamp"] = time.Now().Format("2006-01-02-15-04-05")//1月2日3时4分5秒6年

	_,err = deploymentsClient.Update(deployment)
	return err;
}

func getDeployment(request reqBody.ServiceRequest) *appsv1beta1.Deployment{
	var r apiv1.ResourceRequirements
	//资源分配会遇到无法设置值的问题，故采用json反解析
	//j := `{"limits": {"cpu":"2000m", "memory": "1Gi"}, "requests": {"cpu":"2000m", "memory": "1Gi"}}`
	j := `{"limits": {"cpu":"`+request.LimitCpu+`", "memory": "`+request.LimitMemory+`"}, "requests": {"cpu":"`+request.RequestCpu+`", "memory": "`+request.RequestMemory+`"}}`

	println(j)

	envMap := request.Env;
	//length:=len(envMap)
	var envs  = []apiv1.EnvVar{}

	for entity := range envMap {
		envs = append(envs,apiv1.EnvVar{
			Name:entity,
			Value:envMap[entity],
		} )
	}

	json.Unmarshal([]byte(j), &r)
	deployment := &appsv1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: request.ServiceName,
			Labels: map[string]string{
				"app": request.ServiceName,
			},
		},
		Spec: appsv1beta1.DeploymentSpec{
			Replicas: int32Ptr2(request.InstanceNum),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": request.ServiceName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{   Name:               request.ServiceName,
							Image:           request.Image,
							Env: envs,
							Resources: r,
						},
					},
					//下载镜像密钥
					//ImagePullSecrets: []apiv1.LocalObjectReference{
					//	{Name:config.Get("imagePullSecret")},
					//},
					ImagePullSecrets: getImagePullSecrets(),
					NodeSelector: request.Nodes,

				},

			},
		},
	}
	return deployment;
}


func getImagePullSecrets()[]apiv1.LocalObjectReference{
	imagePullSecrets := config.Get("imagePullSecrets");
	secretArr := strings.Split(imagePullSecrets,",")
	var result = []apiv1.LocalObjectReference{};
	for index := range secretArr {
		result = append(result, apiv1.LocalObjectReference{Name:secretArr[index]})
	}
	return result;
}