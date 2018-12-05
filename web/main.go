package main

import (
	"b2c-deploy/web/appsetting"
	"b2c-deploy/web/routers"
	"runtime"
	"net/http"
	"github.com/fvbock/endless"
	"log"
	"syscall"
	"fmt"
)

func main() {

	//初始化配置文件
	appsetting.Setup()
	//初始化路由表
	routersInit := routers.InitRouter()

	readTimeout := appsetting.ServerSetting.ReadTimeout
	writeTimeout := appsetting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", appsetting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	if runtime.GOOS == "windows" {
		server := &http.Server{
			Addr:           endPoint,
			Handler:        routersInit,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			MaxHeaderBytes: maxHeaderBytes,
		}

		server.ListenAndServe()
		return
	}

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, routersInit)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
