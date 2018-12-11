package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

var cfg *ini.File

func Setup() {

	var err error
	cfg, err = ini.Load("web/resource/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'web/appconfig/app.ini': %v", err)
	}

	mapTo("server", ServerSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
