package service

import (
	"kube-deploy/web/reqBody"

	"errors"
	"kube-deploy/web/responses"
)

func Create(request reqBody.ServiceRequest)(string,error){

	clientset,err := getClientset(request.KubeType)
	if err!=nil{
		//panic(err.Error())
		return "false",err;
	}

	//获取deployment是否存在，如果存在则部署失败
	deployment,err := getK8sDeployment(clientset,request)
	if err ==  nil{
		return "服务已存在",errors.New(deployment.Name+"服务已存在")
	}


	//删除服务，忽略异常
	//_,err = Delete(request)

	//time.Sleep(time.Duration(5)*time.Second)
	err = createDeployment(clientset,request);
	if err != nil {
		return "false",err
	}
	err = createService(clientset,request);
	if err != nil {
		return "false",err
	}
	return "true",nil
}


func Delete(request reqBody.ServiceRequest)(string,error){

	clientset,err := getClientset(request.KubeType)
	if err!=nil{
		//panic(err.Error())
		return "false",err;
	}
	err = deleteDeployment(clientset,request);
	if err != nil {
		//panic(err.Error())
		return "false",err
	}
	err = deleteService(clientset,request);
	if err != nil {
		//panic(err.Error())
		return "false",err
	}
	return "true",nil
}


func Update(request reqBody.ServiceRequest)(string,error){
	clientset,err := getClientset(request.KubeType)
	if err!=nil{
		//panic(err.Error())
		return "false",err;
	}
	err = updateDeployment(clientset,request)
	if err!=nil{
		//panic(err.Error())
		return "false",err;
	}
	return "true",nil
}


func Restart(request reqBody.ServiceRequest)(string,error){
	clientset,err := getClientset(request.KubeType)
	if err!=nil{
		//panic(err.Error())
		return "false",err;
	}
	err = restartDeployment(clientset,request)
	if err!=nil{
		//panic(err.Error())
		return "false",err;
	}
	return "true",nil
}


func Get(request reqBody.ServiceRequest)(responses.DeploymentResp,error){
	result := responses.DeploymentResp{}

	clientset,err := getClientset(request.KubeType)
	if err!=nil{
		//panic(err.Error())
		return result,err;
	}
	deployment,err := getK8sDeployment(clientset,request)
	if err!=nil{
		//panic(err.Error())
		return result,err;
	}

	result.ServiceName = deployment.Name
	result.Namespace = deployment.Namespace
	result.InstanceNum = *deployment.Spec.Replicas
	result.Image = deployment.Spec.Template.Spec.Containers[0].Image
	result.Running = result.InstanceNum == deployment.Status.AvailableReplicas

	return result,nil
}


