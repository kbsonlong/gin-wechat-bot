package main

import (
	_ "github.com/kbsonlong/gin-wechat-bot/pkg/setting"
	"github.com/kbsonlong/gin-wechat-bot/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run()
}
