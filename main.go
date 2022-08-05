package main

import (
	_ "github.com/kbsonlong/gin-wechat-bot/pkg/setting"
	"github.com/kbsonlong/gin-wechat-bot/routers"
)

func main() {

	routersInit := routers.InitRouter()
	routersInit.Run()
}
