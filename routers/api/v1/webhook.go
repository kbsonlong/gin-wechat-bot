package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	var data interface{}
	err1 := json.Unmarshal(jsonData, &data)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(jsonData)
	fmt.Println(data)
	fmt.Print()
	// 钉钉
	// wxchat.JsonPost(
	// 	"https://oapi.dingtalk.com/robot/send?access_token=9bedcb0e28f938ae0ba0840145c09488431af917f366b0db317957ecce0bd9be",
	// 	"{\"msgtype\": \"markdown\",\"markdown\": {\"title\":\"杭州天气gin\",\"text\": \"#### 杭州天气 @150XXXXXXXX \n > 9度,西北风1级,空气良89,相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n\"},\"at\": {\"atMobiles\": [\"150XXXXXXXX\"],\"atUserIds\": [\"user123\"],\"isAtAll\": false}}",
	// )
	// 企微
	// wxchat.JsonPost(
	// 	"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=22ff440f-9b76-4e33-8ba6-63a283c31ee1",
	// 	"{\"msgtype\": \"news\",\"news\": {\"articles\" : [{\"title\" : \"测试中秋节礼品领取\",\"description\" : \"今年中秋节公司有豪礼相送\",\"url\" : \"www.qq.com\",\"picurl\" : \"http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png\"}]}}",
	// )
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
