package main

import (
	"mail/config"
	"mail/routes"
)

func main() {
	config.Init()
	r:=routes.NewRouter()
	r.Run(config.HttpPort)
}

