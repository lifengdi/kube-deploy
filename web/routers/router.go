package routers

import (
	"b2c-deploy/web/api/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv := r.Group("/api/v1")

	{
		apiv.POST("/hello")
	}

	deploys := r.Group("/deploy")

	{
		deploys.POST("/service",service.Create)
		deploys.DELETE("/service")
		deploys.PUT("/service")
	}

	return r
}
