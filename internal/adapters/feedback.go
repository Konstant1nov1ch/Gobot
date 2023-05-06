package adapters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func Feedback(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	//Отзыв - ссылка на мой тг
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пройдите по ссылке, чтобы оставить подробный отзыв https://t.me/kostyapin")
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat - %v", err)
		return
	}
}
