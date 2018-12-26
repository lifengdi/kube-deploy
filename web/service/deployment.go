package service

import (
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

	_, err := deploymentsClient.Create(getDeployment(request))

	return err;
}


func deleteDeployment(clientset *kubernetes.Clientset,request reqBody.ServiceRequest) error {
	deploymentsClient := clientset.AppsV1beta1().Deployments(request.Namespace)

	deletePolicy := metav1.DeletePropagationForeground

	err := deploymentsClient.Delete(request.ServiceName,&metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	});

	return err;
}

func getK8sDeployment(clientset *kubernetes.Clientset,request reqBody.ServiceRequest)(*appsv1beta1.Deployment,error){
	deploymentsClient := clientset.AppsV1beta1().Deployments(request.Namespace)
	return deploymentsClient.Get(request.ServiceName,metav1.GetOptions{})
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
						{
							Name: request.ServiceName,
							Image: request.Image,
							Env: envs,
							Resources: r,
							VolumeMounts: getVolumeMount(request),
							Args: request.Args,
						},
					},
					ImagePullSecrets: getImagePullSecrets(),
					NodeSelector: request.Nodes,
					Volumes: getVolume(request),

				},

			},
		},
	}
	return deployment;
}

/**
	获取镜像下载私钥
 */
func getImagePullSecrets()[]apiv1.LocalObjectReference{
	imagePullSecrets := config.Get("imagePullSecrets");
	secretArr := strings.Split(imagePullSecrets,",")
	var result = []apiv1.LocalObjectReference{};
	for index := range secretArr {
		result = append(result, apiv1.LocalObjectReference{Name:secretArr[index]})
	}
	return result;
}

/**
	获取挂载卷
 */
func getVolume(request reqBody.ServiceRequest) []apiv1.Volume{
	var result = []apiv1.Volume{};
	for index := range request.Volume {
		result = append(result, apiv1.Volume{
			Name:request.Volume[index].Name,
			VolumeSource: apiv1.VolumeSource{
				HostPath: &apiv1.HostPathVolumeSource{
					Path:request.Volume[index].HostPath,
				},
			},
		})
	};

	return result;
}

/**
	将挂载卷挂载到容器的某目录
 */
func getVolumeMount(request reqBody.ServiceRequest) []apiv1.VolumeMount{
	var result = []apiv1.VolumeMount{};
	for index := range request.VolumeMount {
		result = append(result, apiv1.VolumeMount{
			Name:request.VolumeMount[index].Name,
			MountPath: request.VolumeMount[index].MountPath,
		})
	};

	return result;
}