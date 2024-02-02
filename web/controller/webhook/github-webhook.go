package webhook

import (
	"log"
	"net/http"
	"os"

	"gin-sites/controller/webhook/feishu"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/github"
)

var logger = log.Logger{}

type MsgHandle func(token string, payload github.PushPayload)

// Header github自定义的请求头
// X-GitHub-Event string  事件名
// X-GitHub-Deliver string 递送的GUID
// X-Hub-Signature
// X-Hub-Signature-256 Webhook 中设置了secret 时用于验证的数据
// User-Agent 含有前缀 GitHub-Hookshot/

// 接受 web hook 请求
func GithubWebhookHandle(c *gin.Context) {
	logFile, _ := os.OpenFile("./uploads/github.txt", os.O_CREATE|os.O_APPEND, 0766)

	logger.SetOutput(logFile)
	logger.Println("==========START=============")
	defer logFile.Close()
	defer logger.Println("=============END============")

	// 进行鉴权处理
	accessToken := c.Query("access_token")
	if !checkToken(accessToken) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "token 无效",
		})
		return
	}

	hook, _ := github.New()
	// parse 函数可以继续添加想处理的事件
	payload, err := hook.Parse(c.Request, github.PushEvent)

	if err != nil {
		c.JSON(400, gin.H{"msg": "failed to bind header" + err.Error()})
		return
	}
	switch request := payload.(type) {
	// push 事件处理
	case github.PushPayload:
		{
			handleWebhook(accessToken, request)
		}
	}
	c.JSON(200, gin.H{"msg": "成功接收到webhook 事件"})

}

// 平台名称和处理函数注册位置
var platformHandles = map[string]MsgHandle{
	"feishu": feishu.FeishuHandle,
}

func RegisterPlatform(name string, handle MsgHandle) {
	platformHandles[name] = handle
}

// 检测 access token 是否有效
func checkToken(token string) bool {
	// 检查token是否有效
	if token == "" {
		return false
	} else {
		// 做具体的检查
	}

	return true
}

// 依据 token 返回对应的 platform
func getPlatformsByToken(token string) []string {
	//

	return []string{"feishu"}
}

// 处理接收到的钩子事件
//
// 依据 token 查找需要处理方法及参数
func handleWebhook(token string, payload github.PushPayload) {

	for _, platform := range getPlatformsByToken(token) {

		if pl, ok := platformHandles[platform]; ok {
			pl(token, payload)
		} else {
			logger.Println("platform: ", platform, " 未注册")
		}

	}
}
