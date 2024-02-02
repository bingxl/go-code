package feishu

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/webhooks/v6/github"
)

// ---------------------飞书机器人处理
func FeishuHandle(token string, payload github.PushPayload) {
	config := getFeishuConfig(token)

	hook := config["hook"]
	body := getFeishuBody(payload)

	_, err := http.Post(
		hook,
		"application/json",
		body,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func getFeishuConfig(token string) map[string]string {
	// TODO 依据token 获取到飞书机器人配置
	return map[string]string{
		"hook": "https://open.feishu.cn/open-apis/bot/v2/hook/1f983a38-c25a-4319-8f28-7da73e4adbed",
	}
}

func getFeishuBody(payload github.PushPayload) io.Reader {
	template := `
	{
		"msg_type":"post",
		"content": {
			"post": {
				"zh_cn": {
					"title": "仓库: %s 发生了 PUSH 事件",
					"content": [
						[
							{
								"tag": "text",
								"text": "%s"
							},
							{
								"tag": "a",
								"text": "点击查看",
								"href": "%s"
							}
						]
					]
				}
			}
		}
	}
	`

	result := fmt.Sprintf(
		template,
		payload.Repository.FullName,
		payload.HeadCommit.Message,
		payload.Repository.URL,
	)

	result = strings.ReplaceAll(result, "\n", "")
	result = strings.ReplaceAll(result, "\t", "")
	return strings.NewReader(result)
}

// ######################### end
