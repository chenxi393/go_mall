package main

import (
	"mail/config"
	"mail/routes"
)

func main() {
	//config.Init() 这里走上面的init初始化
	r:=routes.NewRouter()
	r.Run(config.HttpPort)
}

