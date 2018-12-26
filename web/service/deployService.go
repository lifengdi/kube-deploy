package service

import (
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
	"kube-deploy/web/reqBody"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
	"strings"
)


/**
	创建service
 */
func createService(clientset *kubernetes.Clientset,request reqBody.ServiceRequest) error {
	_, err:= clientset.CoreV1().Services(request.Namespace).Create(&apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: request.ServiceName,
		},
		Spec: apiv1.ServiceSpec{
			Type:     apiv1.ServiceTypeClusterIP,
			Selector: map[string]string{
				"app": request.ServiceName,
			},
			Ports: getServicePort(request),
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

/**
	获取服务端口映射
 */
func getServicePort(request reqBody.ServiceRequest)[]apiv1.ServicePort{
	var result = []apiv1.ServicePort{}

	if len(request.Ports) > 0{
		for index := range request.Ports {

			portType := request.Ports[index].Type;

			if portType == ""{
				portType = "TCP"
			}

			result = append(result, apiv1.ServicePort{
				Port: request.Ports[index].Port,
				TargetPort: intstr.FromInt(request.Ports[index].TargetPort),
				Protocol: apiv1.Protocol(request.Ports[index].Type),
				Name: strings.ToLower(portType)+strconv.Itoa(int(request.Ports[index].Port)),
			})
		}
	}else {
		result = append(result, apiv1.ServicePort{
			Port: request.Port,
			TargetPort: intstr.FromInt(request.TargetPort),
			Protocol: apiv1.Protocol("TCP"),
			Name: "tcp"+strconv.Itoa(int(request.Port)),
		})
	}
	return result;
}