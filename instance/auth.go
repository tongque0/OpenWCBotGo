package instance

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"

	"github.com/eatmoreapple/openwechat"
)

var (
	allowedSenders    = make(map[string]bool)
	currentVerifyCode = ""
	mu                sync.Mutex
)

const verifyCommand = "!8888" // 设置一个特殊的指令来触发验证码生成

// generateVerificationCode 生成随机验证码
func generateVerificationCode() string {
	bytes := make([]byte, 4) // 生成一个4字节的随机数
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

// AuthMiddleware 处理用户认证逻辑
func AuthMiddleware(msg *openwechat.Message) bool {
	if msg.IsSendBySelf() {
		return false
	}

	// 获取消息的接收者
	user, _ := msg.Sender()
	receiver := user.String()
	mu.Lock()
	defer mu.Unlock()

	// 检查是否是验证请求
	if strings.TrimSpace(msg.Content) == verifyCommand {
		currentVerifyCode = generateVerificationCode()
		msg.ReplyText(fmt.Sprintf(currentVerifyCode))
		return false
	}

	// 检查验证码
	if strings.TrimSpace(msg.Content) == currentVerifyCode {
		allowedSenders[receiver] = true
		currentVerifyCode = "" // 重置验证码，确保每个验证码只能使用一次
		msg.ReplyText("可以开始聊天了")
		return false
	}

	if msg.IsSendByGroup() && allowedSenders[receiver] && containsWakeWord(msg.Content) {
		return true
	}
	if msg.IsSendByFriend() && allowedSenders[receiver] {
		return true
	}
	return false
}

func containsWakeWord(content string) bool {
	wakeWords := []string{"小喜", "喜羊羊"} // 这里可以定义更多的唤醒词

	for _, word := range wakeWords {
		if strings.Contains(content, word) {
			return true
		}
	}
	return false
}
