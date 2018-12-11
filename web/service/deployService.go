package service

import (
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
	"b2c-deploy/web/reqBody"
	"fmt"
	"k8s.io/apimachinery/pkg/util/intstr"
)


/**
	创建service
 */
func createService(clientset *kubernetes.Clientset,request reqBody.ServiceRequest) error {
	fmt.Println("createService")
	_, err:= clientset.CoreV1().Services(request.Namespace).Create(&apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: request.ServiceName,
		},
		Spec: apiv1.ServiceSpec{
			Type:     apiv1.ServiceTypeClusterIP,
			Selector: map[string]string{
				"app": request.ServiceName,
			},
			Ports: []apiv1.ServicePort{
				{
					Port: request.Port,
					TargetPort: intstr.FromInt(request.TargetPort),
				},
			},
		},
	})
	return err
}

func deleteService(clientset *kubernetes.Clientset,request reqBody.ServiceRequest) error {

	deletePolicy := metav1.DeletePropagationForeground

	err:=clientset.CoreV1().Services(request.Namespace).Delete(request.ServiceName,&metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	return err;
}
