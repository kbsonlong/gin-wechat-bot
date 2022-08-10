package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kbsonlong/gin-wechat-bot/pkg/setting"
	"github.com/kbsonlong/gin-wechat-bot/pkg/utils"
	"github.com/kbsonlong/gin-wechat-bot/pkg/webhook"
)

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary webhook
// @Description 机器人推送
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /api/webhook/ [post]
func Webhook(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)

	data := make(map[string]interface{})

	json.Unmarshal(jsonData, &data)

	// 钉钉告警
	if setting.Conf.BotConfig.DingTalkConfig.Enable {
		res := utils.Parse(setting.Conf.BotConfig.DingTalkConfig.TemplateFile, data)
		// fmt.Print(res)

		for _, secret := range setting.Conf.BotConfig.DingTalkConfig.Secrets {
			// fmt.Println(setting.Conf.BotConfig.DingTalkConfig.CallBackUrl, secret)
			webhook_url := fmt.Sprint(setting.Conf.BotConfig.DingTalkConfig.CallBackUrl, secret)
			// fmt.Println(webhook_url)
			webhook.JsonPost(webhook_url, res)
		}
	}
	// 企微告警
	if setting.Conf.BotConfig.WxChatConfig.Enable {
		res := utils.Parse(setting.Conf.BotConfig.WxChatConfig.TemplateFile, data)
		// fmt.Print(res)

		for _, secret := range setting.Conf.BotConfig.WxChatConfig.Secrets {
			// fmt.Println(setting.Conf.BotConfig.WxChatConfig.CallBackUrl, secret)
			webhook_url := fmt.Sprint(setting.Conf.BotConfig.WxChatConfig.CallBackUrl, secret)
			// fmt.Println(webhook_url)
			webhook.JsonPost(webhook_url, res)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
