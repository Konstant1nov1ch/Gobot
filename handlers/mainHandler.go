package handlers

import (
	"algoru/controller"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func MainHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI, apiKey string) {
	// Получаем ответ от openAi
	response, err := controller.Ask(strings.TrimSpace(update.Message.Text), apiKey)
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
