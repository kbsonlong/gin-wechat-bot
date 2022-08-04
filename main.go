package main

import "github.com/kbsonlong/gin-wechat-bot/routers"

func main() {
	routersInit := routers.InitRouter()
	routersInit.Run()
}
