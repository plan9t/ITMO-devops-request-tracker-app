package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

// Функция для запуска бота
func startBot() {
	var err error
	bot, err = tgbotapi.NewBotAPI("8084015321:AAErWTJ9gqu2wmUHGt43OXvKDvPoq1YUTSA") // Замените на ваш токен
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	// Устанавливаем параметры для long polling
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 // Таймаут в секундах

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // Если сообщение не пустое
			handleMessage(update.Message.Text, update.Message.Chat.ID)
		}
	}
}

func handleMessage(message string, chatID int64) {
	if message == "/start" {
		sendMessage(chatID, "Введите ID заказа:")
	} else {
		orderID := message
		response := getOrderByUID(orderID) // Вызов функции для получения заказа
		sendMessage(chatID, response)
	}
}

func sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
	}
}

func getOrderByUID(orderID string) string {
	resp, err := http.Get(fmt.Sprintf("http://localhost:4444/api/orders/%s", orderID))
	if err != nil {
		return "Ошибка при получении заказа."
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var order Order // Используем вашу структуру Order
		if err := json.NewDecoder(resp.Body).Decode(&order); err != nil {
			return "Ошибка при декодировании ответа."
		}

		// Форматируем ответ в виде JSON
		orderJSON, err := json.MarshalIndent(order, "", "  ")
		if err != nil {
			return "Ошибка при формировании ответа."
		}

		return string(orderJSON) // Возвращаем JSON как строку
	}

	return "Заказ не найден."
}
