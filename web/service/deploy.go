package service

import (
	"kube-deploy/web/reqBody"

	"time"
)

func Create(request reqBody.ServiceRequest)(string,error){
	println(request.Image)


	clientset,err := getClientset(request.KubeType)
	if err!=nil{
		//panic(err.Error())
		return "false",err;
	}
	//删除服务，忽略异常
	_,err = Delete(request)

	//todo 检查是否仍存在该deployment  如果存在，则休眠5s后再次检查 检查10次后若仍然存在，则返回错误
	time.Sleep(time.Duration(5)*time.Second)
	err = createDeployment(clientset,request);
	if err != nil {
		panic(err.Error())
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


