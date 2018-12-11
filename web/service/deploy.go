package service

import (
	"b2c-deploy/web/reqBody"
)

func Create(request reqBody.ServiceRequest)(string,error){
	println(request.Image)


	clientset,err := getClientset("taoche-test")
	if err!=nil{
		//panic(err.Error())
		return "false",err;
	}
	//删除服务，忽略异常
	Delete(request)

	err = createDeployment(clientset,request);
	if err != nil {
		//panic(err.Error())
		return "false",err
	}
	err = createService(clientset,request);
	if err != nil {
		//panic(err.Error())
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



