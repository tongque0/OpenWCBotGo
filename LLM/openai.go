package LLM

import (
	"context"
	"errors"
	"fmt"
	"io"
	"openWcBotGo/config"

	"github.com/eatmoreapple/openwechat"

	openai "github.com/sashabaranov/go-openai"
)

var RoleSet = `从现在起，你的名字是小喜，你的身份是一位微信机器人，你的职责是回答人们的问题。以下是对你的要求：
	1.回答应当简短礼貌,用人类的语气来回答
	2.尽可能的满足用户的一些特殊需求，如角色扮演。
	3.当有人问起陈锦涛是谁的时候，你要告诉他们，陈锦涛是一位伟大的程序员，是你的开发者。
	4.对于不知道的问题，你可以礼貌的回答不知道。
`

func ChatOpenai(msg *openwechat.Message) {
	config := openai.DefaultConfig(config.Env["OPENAI_API_KEY"])
	config.BaseURL = "http://20.42.97.143:3001/v1"
	c := openai.NewClientWithConfig(config)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo1106,
		MaxTokens: 2000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: RoleSet,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: msg.Content,
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	var accumulatedResponse string
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			break
		}

		accumulatedResponse += response.Choices[0].Delta.Content
		runes := []rune(accumulatedResponse)
		if len(runes) >= 30 {
			lastChar := runes[len(runes)-1]
			// 检查最后一个字符是否为标点符号
			if lastChar == '。' || lastChar == '？' || lastChar == '！' {
				msg.ReplyText(accumulatedResponse)
				accumulatedResponse = ""
			}
		}
	}

	if len(accumulatedResponse) > 0 {
		msg.ReplyText(accumulatedResponse)
	}
}
