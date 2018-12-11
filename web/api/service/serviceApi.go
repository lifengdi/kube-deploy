package service



import (
	"fmt"
	"github.com/gin-gonic/gin"
	"b2c-deploy/web/responses"
	"b2c-deploy/web/exceptions"
	"net/http"
	"b2c-deploy/web/reqBody"
	"b2c-deploy/web/service"
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
		fmt.Println("响应失败")
		return
	}
	resp := responses.Gin{C: c}

	if req.Image == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "镜像不能为空")
		return
	}
	if req.ServiceName == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "服务名不能为空")
		return
	}
	println(req.Namespace)

	data,err := service.Create(req)
	if err != nil {
		resp.Response(http.StatusOK, exceptions.ERROR, nil)
		return
	}
	resp.Response(http.StatusOK, exceptions.SUCCESS, data)

}

func DeleteService(c *gin.Context){
	req := reqBody.InitServiceRequest()
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("响应失败")
		return
	}
	resp := responses.Gin{C: c}
	if req.ServiceName == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "服务名不能为空")
		return
	}

	result,err:=service.Delete(req);

	if err != nil {
		resp.Response(http.StatusOK, exceptions.ERROR, nil)
		panic(err)
		return
	}
	resp.Response(http.StatusOK, exceptions.SUCCESS, result)
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
		fmt.Println("响应失败")
		return
	}
	resp := responses.Gin{C: c}
	if req.ServiceName == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "服务名不能为空")
		return
	}
	if req.Image == ""{
		resp.Response(http.StatusOK, exceptions.ERROR, "镜像不能为空")
		return
	}
	result,err:=service.Update(req);

	if err != nil {
		resp.Response(http.StatusOK, exceptions.ERROR, nil)
		panic(err)
		return
	}
	resp.Response(http.StatusOK, exceptions.SUCCESS, result)
}