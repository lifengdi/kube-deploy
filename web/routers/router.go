package routers

import (
	"kube-deploy/web/api/service"

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
		deploys.POST("/service",service.CreateService)
		deploys.DELETE("/service",service.DeleteService)
		deploys.PUT("/service",service.UpdateService)
		deploys.PATCH("/service",service.Restart)
	}

	return r
}
