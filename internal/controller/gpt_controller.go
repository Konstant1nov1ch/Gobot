package controller

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func GetAnswerFromOpenAI(userText string, token string) (string, string) {
	client := openai.NewClient(token)
	resp, err3 := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:       openai.GPT3TextDavinci003,
			Prompt:      userText,
			MaxTokens:   1024,
			Temperature: 0.6,
		},
	)
	if err3 != nil {
		fmt.Printf("Completion error: %v\n", err3)
		return "Что-то пошло не так, попробуйте позже... .", ""
	}

	return resp.Choices[0].Text, ""
}
