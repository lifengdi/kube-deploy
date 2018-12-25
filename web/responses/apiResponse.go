package responses

import (
	"github.com/gin-gonic/gin"
	"kube-deploy/web/exceptions"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode , errCode int, data interface{}) {
	g.C.JSON(httpCode,gin.H{
		"code": httpCode,
		"msg":  exceptions.GetMsg(errCode),
		"data": data,
	})

}
