package routers

import (
	"github.com/gin-gonic/gin"
	"b2c-deploy/web/routers/api"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv := r.Group("/api/v1")

	{
		apiv.POST("/hello")
		apiv.GET("/hello",api.Deploy)
	}

	return r
}
