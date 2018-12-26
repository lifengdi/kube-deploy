package service



import (
	"github.com/gin-gonic/gin"
	"kube-deploy/web/responses"
	"kube-deploy/web/exceptions"
	"net/http"
	"kube-deploy/web/reqBody"
	"kube-deploy/web/service"
	"kube-deploy/web/logger"
)

/**
	url:http://localhost:8080/deploy/service
	body:{
		"serviceName":"nginx",
		"image":"registry.cn-beijing.aliyuncs.com/kevin-public/nginx:1.0.0",
		"port":80,
		"targetPort":80,
		"instanceNum":1
	}
 */
func CreateService(c *gin.Context)  {
	req := reqBody.InitServiceRequest()
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("响应失败")
		return
	}
	resp := responses.Gin{C: c}

	if req.Image == ""{
		resp.Response(http.StatusOK, exceptions.ERROR,"镜像不能为空",nil )
		return
	}
	if req.ServiceName == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "服务名不能为空",nil)
		return
	}
	println(req.Namespace)

	_,err := service.Create(req)
	if err != nil {
		logger.Error(err.Error())
		resp.Response(http.StatusOK, exceptions.ERROR, err.Error(),nil)
		return
	}
	resp.Response(http.StatusOK, exceptions.SUCCESS, "SUCCESS",nil)

}

func DeleteService(c *gin.Context){
	req := reqBody.InitServiceRequest()
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("响应失败")
		return
	}
	resp := responses.Gin{C: c}
	if req.ServiceName == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "服务名不能为空",nil)
		return
	}

	result,err:=service.Delete(req);

	if err != nil {
		logger.Error(err.Error())
		resp.Response(http.StatusOK, exceptions.ERROR, err.Error(),nil)
		return
	}
	resp.Response(http.StatusOK, exceptions.SUCCESS, result,nil)
}


/**
	{
    "serviceName":"nginx",
    "image":"nginx",
    "port":80,
    "targetPort":80,
    "requestCpu":"0.5"
}
 */
func UpdateService(c *gin.Context){
	req := reqBody.InitServiceRequest()
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("响应失败")
		return
	}
	resp := responses.Gin{C: c}
	if req.ServiceName == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "服务名不能为空",nil)
		return
	}
	if req.Image == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "镜像不能为空",nil)
		return
	}
	result,err:=service.Update(req);

	if err != nil {
		logger.Error(err.Error())
		resp.Response(http.StatusOK, exceptions.ERROR, err.Error(),nil)
		return
	}
	resp.Response(http.StatusOK, exceptions.SUCCESS, result,nil)
}

func Restart(c *gin.Context){
	req := reqBody.InitServiceRequest()
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("响应失败")
		return
	}
	resp := responses.Gin{C: c}

	if req.ServiceName == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "服务名不能为空",nil)
		return
	}

	result,err:=service.Restart(req);

	if err != nil {
		logger.Error(err.Error())
		resp.Response(http.StatusOK, exceptions.ERROR, err.Error(),nil)
		return
	}
	resp.Response(http.StatusOK, exceptions.SUCCESS, result,nil)
}

func Get(c *gin.Context){
	req := reqBody.InitServiceRequest()
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("响应失败")
		return
	}
	resp := responses.Gin{C: c}

	if req.ServiceName == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "服务名不能为空",nil)
		return
	}

	result,err:=service.Get(req);
	if err != nil {
		logger.Error(err.Error())
		resp.Response(http.StatusOK, exceptions.ERROR, err.Error(),nil)
		return
	}
	resp.Response(http.StatusOK, exceptions.SUCCESS,"SUCCESS", result)
}