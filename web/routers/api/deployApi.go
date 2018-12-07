package api

import (
	"github.com/gin-gonic/gin"
	"b2c-deploy/web/responses"
	"b2c-deploy/web/service"
	"net/http"
	"b2c-deploy/web/exceptions"
)

func Deploy(c *gin.Context) {


	appG := responses.Gin{C: c}

	deployService := service.Demo{
		SourceId: "",
		UserId: "",
	}

	data,err := deployService.GetDemo()

	if err != nil {
		appG.Response(http.StatusOK, exceptions.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, exceptions.SUCCESS, data)

}
