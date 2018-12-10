package service

import (
	"encoding/json"
	"fmt"
	"b2c-deploy/web/reqBody"
	"k8s.io/client-go/kubernetes"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
)

func createDeployment(clientset *kubernetes.Clientset,request reqBody.CreateRequest) error {


	deploymentsClient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)

	var r apiv1.ResourceRequirements
	//资源分配会遇到无法设置值的问题，故采用json反解析
	//j := `{"limits": {"cpu":"2000m", "memory": "1Gi"}, "requests": {"cpu":"2000m", "memory": "1Gi"}}`
	j := `{"limits": {"cpu":"0.1", "memory": "32Mi"}, "requests": {"cpu":"0.1", "memory": "32Mi"}}`
	json.Unmarshal([]byte(j), &r)
	deployment := &appsv1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: request.ServiceName,
			Labels: map[string]string{
				"app": request.ServiceName,
			},
		},
		Spec: appsv1beta1.DeploymentSpec{
			Replicas: int32Ptr2(1),
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
							Resources: r,
						},
					},
				},
			},
		},
	}

	fmt.Println("Creating deployment...")
	_, err := deploymentsClient.Create(deployment)

	return err;
}
