package main

import (
	"algoru/controller"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Data struct {
	Key1 string `json:"keyTg"`
	Key2 string `json:"keyAi"`
}

func main() {
	var data, err = ioutil.ReadFile("./resources/tokens.json")
	if err != nil {
		log.Panic(err)
	}
	var d Data
	err = json.Unmarshal(data, &d)
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(d.Key1)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				start(update, bot)
			case "feedback":
				feedback(update, bot)
			default:
				continue
			}
			continue
		}

		if update.CallbackQuery != nil {
			// Обработка нажатия кнопки
			buttonText := update.CallbackQuery.Data
			resp, err1 := controller.Ask(buttonText, d.Key2)
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
			mainHandler(update, bot, d.Key2)
		}
	}
}

func start(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	message := "👋 Привет! Я искусственный интеллект, и я могу ответить на ваши вопросы! " +
		"Просто напишите запрос, например:\n\n👉 Расскажи про поиск в бинарном дереве?\n\n" +
		"——————————————————————————————\n" +
		"💬 Если хотите помочь развитию проекта, оставьте отзыв с помощью команды /feedback"

	// Создаем первый уровень кнопок
	row1 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Как меня зовут?", "What's is your name? (Answer in Russian)"),
		tgbotapi.NewInlineKeyboardButtonData("Сколько мне лет?", "How old of you (Answer in Russian)"),
		tgbotapi.NewInlineKeyboardButtonData("Какой сегодня день?", "What a day today? (Answer in Russian)"),
		tgbotapi.NewInlineKeyboardButtonData("Какой у меня город?", "What is your hometown ? (Answer in Russian)"),
	)

	row2 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Мой любимый цвет?", "What is your favorite color? (Answer in Russian)"),
		tgbotapi.NewInlineKeyboardButtonData("Апроксимация методом Ньютона?", "Can you write about newton method? (Answer in Russian)"),
	)

	row3 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Решить квадратное уравнение?", "Solve a quadratic equation? (Answer in Russian)"),
		tgbotapi.NewInlineKeyboardButtonData("Интеграл?", "Can you write about newton integral? (Answer in Russian)"),
	)

	row4 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Почему небо голубое?", "Why is the sky blue? (Answer in Russian)"),
		tgbotapi.NewInlineKeyboardButtonData("Почему трава зеленая?", "Why is the grass green? (Answer in Russian)"),
	)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

	msg.ReplyMarkup = keyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat - %v", err)
		return
	}
}

func feedback(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пройдите по ссылке, чтобы оставить подробный отзыв https://forms.gle/UkhZzFWPpipTvRiQA")
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat - %v", err)
		return
	}
}

func mainHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI, apiKey string) {
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
