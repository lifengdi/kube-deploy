package service

import (
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

func listPods(clientset *kubernetes.Clientset) error{
	pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	println(pods)
	return err;
}