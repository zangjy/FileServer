package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"gopkg.in/ini.v1"
	"log"
	"strconv"
)

func main() {
	cfg, err := ini.Load("./conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件: %v", err)
	}

	ip := cfg.Section("server").Key("ip").String()
	port := cfg.Section("server").Key("port").MustInt(8080)
	staticRoot := cfg.Section("server").Key("static_path").String()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = colorable.NewColorableStdout()
	gin.ForceConsoleColor()

	router := gin.Default()

	router.Static("/", staticRoot)

	err = router.Run(ip + ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
