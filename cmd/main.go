package main

import (
	handlers2 "algoru/internal/adapters"
	"algoru/internal/controller"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

// Data Структура json файла с токенами (в гитигноре)
type Data struct {
	Key1 string `json:"keyTg"`
	Key2 string `json:"keyAi"`
}

func main() {
	//Читаем файл
	var data, err = os.ReadFile("internal/resources/tokens.json")
	if err != nil {
		log.Panic("FileNotFoundException")
	}
	var d Data
	err = json.Unmarshal(data, &d)
	if err != nil {
		log.Panic("IOException")
	}
	//создаем сущность бота
	bot, err := tgbotapi.NewBotAPI(d.Key1)
	if err != nil {
		log.Panic("BotInitError")
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	//следим за изменениями в чате
	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}
		//обработчик команд
		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				handlers2.Start(update, bot)
			case "feedback":
				handlers2.Feedback(update, bot)
			default:
				continue
			}
			continue
		}

		if update.CallbackQuery != nil {
			// Обработка нажатия кнопки
			buttonText := update.CallbackQuery.Data
			resp, err1 := controller.GetAnswerFromOpenAI(buttonText, d.Key2)
			if err1 != "" {
				resp = "Что-то пошло не так, попробуйте позже... ."
			}
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, resp)
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message to chat - %v", err)
				return
			}
		} else {
			handlers2.MainHandler(update, bot, d.Key2)
		}
	}
}