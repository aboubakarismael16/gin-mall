package main

import (
	"gin-mall/conf"
	"gin-mall/routers"
)

func main() {
	conf.Init()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}
