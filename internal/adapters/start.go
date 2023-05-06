package adapters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func Start(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	message := "👋 Привет! Я искусственный интеллект, и я могу ответить на ваши вопросы! " +
		"Просто напишите запрос, например:\n\n👉 Расскажи про поиск в бинарном дереве?\n\n" +
		"——————————————————————————————\n" +
		"💬 Если хотите помочь развитию проекта, оставьте отзыв с помощью команды /feedback"

	// Создаем слои кнопок
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
		tgbotapi.NewInlineKeyboardButtonData("Интеграл?", "Can you write about integral? (Answer in Russian)"),
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
