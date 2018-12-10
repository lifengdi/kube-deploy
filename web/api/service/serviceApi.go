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
		"serviceName":"test",
		"InstanceNum":4,
		"image":"consul"
	}
 */

func Create(c *gin.Context)  {
	var req reqBody.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("响应失败")
		return
	}

	resp := responses.Gin{C: c}


	data,err := service.Create(req)

	if err != nil {
		resp.Response(http.StatusOK, exceptions.ERROR, nil)
		return
	}

	resp.Response(http.StatusOK, exceptions.SUCCESS, data)


}

