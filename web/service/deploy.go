package service

import (
	"kube-deploy/web/reqBody"

	"errors"
)

func Create(request reqBody.ServiceRequest)(string,error){
	println(request.Image)


	clientset,err := getClientset(request.KubeType)
	if err!=nil{
		//panic(err.Error())
		return "false",err;
	}

	//获取deployment是否存在，如果存在则部署失败
	deployment,err := getK8sDeployment(clientset,request)
	if err ==  nil{
		println(deployment.Name+" deployment not is null")
		return "服务已存在",errors.New("服务已存在")
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
		panic(err.Error())
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


