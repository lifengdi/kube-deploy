package responses

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode , errCode int,errorMsg string, data interface{}) {
	g.C.JSON(httpCode,gin.H{
		"code": errCode,
		"msg":  errorMsg,
		"data": data,
	})

}

