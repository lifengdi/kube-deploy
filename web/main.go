package main

import (
	"kube-deploy/web/setting"
	"kube-deploy/web/routers"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"syscall"
	"github.com/fvbock/endless"
	"flag"
	"kube-deploy/web/config"
	"kube-deploy/web/logger"
)

func main() {

	initConf();

	//初始化配置文件
	setting.Setup()
	//初始化路由表
	routersInit := routers.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
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

// todo 入参有问题
func initConf(){
	var kubeConfs string
	var appConf string
	var imagePullSecrets string
	var log string

	flag.StringVar(&kubeConfs, "kubeconfs", "/Users/liukai/go/src/kube-deploy/web/resource/", "kubeconfs dir")
	flag.StringVar(&appConf, "f", "/Users/liukai/go/src/kube-deploy/web/resource/app.ini", "app.ini path")
	flag.StringVar(&imagePullSecrets, "imagePullSecrets", "tencent-registry,kanche-registry", "docker image pull secret")
	flag.StringVar(&log, "log", "/var/log/kubebernetes/", "log dir")

	println("imagePullSecrets:"+imagePullSecrets)
	flag.Parse()
	config.Set("appConf",appConf)
	config.Set("kubeConfs",kubeConfs)
	config.Set("imagePullSecrets",imagePullSecrets)
	//config.Set("log",log)
	logger.LogDir = log;

}