package controller

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func GetAnswerFromOpenAI(userText string, token string) (string, string) {
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0301,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userText,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "Что-то пошло не так, попробуйте позже... .", ""
	}
	return resp.Choices[0].Message.Content, ""
}
