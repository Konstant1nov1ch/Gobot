package adapters

import (
	"algoru/internal/controller"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func Message(update tgbotapi.Update, bot *tgbotapi.BotAPI, apiKey string) {
	// Получаем ответ от openAi
	response, err := controller.GetAnswerFromOpenAI(strings.TrimSpace(update.Message.Text), apiKey)
	if err != "" {
		log.Printf("Error generating text - %v", err)
		return
	}
	// Отправляем ответ пользователю
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	_, err2 := bot.Send(msg)
	if err2 != nil {
		log.Printf("Error sending message to chat - %v", err2)
		return
	}
}
