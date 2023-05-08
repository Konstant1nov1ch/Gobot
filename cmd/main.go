package main

import (
	"algoru/internal/adapters"
	"algoru/internal/controller"
	"encoding/json"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
	if err != nil {
		log.Panic("UpdatesChanError")
	}

	// Создаем три канала
	commandChan := make(chan tgbotapi.Update)
	messageChan := make(chan tgbotapi.Update)
	callbackQueryChan := make(chan tgbotapi.Update)

	// Запускаем обработку команд в первом канале
	go func() {
		for update := range commandChan {
			if update.Message != nil && update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					// запускаем обработку сообщения в новой горутине
					go adapters.Start(update, bot)
				case "feedback":
					// запускаем обработку сообщения в новой горутине
					go adapters.Feedback(update, bot)
				default:
					continue
				}
			}
		}
	}()

	// Запускаем обработку сообщений во втором канале
	go func() {
		for update := range messageChan {
			// запускаем обработку сообщения в новой горутине
			go adapters.Message(update, bot, d.Key2)
		}
	}()

	// Запускаем обработку кнопок в третьем канале
	go func() {
		for update := range callbackQueryChan {
			buttonText := update.CallbackQuery.Data
			resp, err := controller.GetAnswerFromOpenAI(buttonText, d.Key2)
			if err != "" {
				resp = "Что-то пошло не так, попробуйте позже..."
			}
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, resp)
			_, err1 := bot.Send(msg)
			if err1 != nil {
				log.Printf("Error sending callback query response: %v", err)
			}
		}
	}()
	// Запускаем бесконечный цикл чтения обновлений
	for update := range updates {
		if update.Message != nil {
			// если сообщение содержит команду, отправляем его в первый канал
			if update.Message.IsCommand() {
				commandChan <- update
			} else {
				// иначе отправляем второму каналу
				messageChan <- update
			}
			//если сообщения нет но нажата кнопка то 3 канал
		} else if update.CallbackQuery != nil {
			// если обновление содержит callback query, отправляем его в третий канал
			callbackQueryChan <- update
		}
	}
}
