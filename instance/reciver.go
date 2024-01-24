package instance

import (
	"fmt"
	"openWcBotGo/LLM"

	"github.com/eatmoreapple/openwechat"
)

func reciver(msg *openwechat.Message) {
	if !AuthMiddleware(msg) {
		return
	}
	go handleMessage(msg)
}

func handleMessage(msg *openwechat.Message) {
	if msg.IsText() {
		LLM.ChatOpenai(msg)
	} else if msg.IsPicture() {
		fmt.Println("图片消息")
	} else if msg.IsVoice() {
		fmt.Println("语音消息")
	} else if msg.IsVideo() {
		fmt.Println("视频消息")
	} else if msg.IsCard() {
		fmt.Println("名片消息")
	} else if msg.IsFriendAdd() {
		fmt.Println("好友认证")
	} else if msg.IsRecalled() {
		fmt.Println("撤回")
	} else {
		fmt.Println("其他类型消息")
	}
}
