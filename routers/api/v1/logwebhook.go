package v1

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kbsonlong/gin-wechat-bot/pkg/setting"
	"github.com/kbsonlong/gin-wechat-bot/pkg/webhook"
)

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary opensearch告警推送
// @Description Post alter from opensearch
// @Tags webhooks
// @Accept  json
// @Produce  json
// @Param        string      data     string  false  "string valid"
// @Success      200         {string}  string  "answer"
// @Failure      400         {string}  string  "ok"
// @Failure      404         {string}  string  "ok"
// @Failure      500         {string}  string  "ok"
// @Router /api/v1/log_webhook/ [post]
func LogWebhook(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(data))
	t, err := template.ParseFiles("templates/log.tmpl")
	// Capture any error
	if err != nil {
		log.Fatalln(err)
	}
	// Print out the template to std
	var tpl bytes.Buffer
	t.Execute(&tpl, string(data))
	// fmt.Print(tpl.String())
	// 钉钉告警
	if setting.Conf.BotConfig.DingTalkConfig.Enable {
		for _, secret := range setting.Conf.BotConfig.DingTalkConfig.Secrets {
			// fmt.Println(setting.Conf.BotConfig.DingTalkConfig.CallBackUrl, secret)
			webhook_url := fmt.Sprint(setting.Conf.BotConfig.DingTalkConfig.CallBackUrl, secret)
			// fmt.Println(webhook_url)
			webhook.JsonPost(webhook_url, fmt.Sprint(tpl.String()))
		}
	}
	// 企微告警
	if setting.Conf.BotConfig.WxChatConfig.Enable {
		for _, secret := range setting.Conf.BotConfig.WxChatConfig.Secrets {
			// fmt.Println(setting.Conf.BotConfig.WxChatConfig.CallBackUrl, secret)
			webhook_url := fmt.Sprint(setting.Conf.BotConfig.WxChatConfig.CallBackUrl, secret)
			// fmt.Println(webhook_url)
			webhook.JsonPost(webhook_url, fmt.Sprint(tpl.String()))
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
