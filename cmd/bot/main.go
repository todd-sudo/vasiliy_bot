package main

import (
	"github.com/todd-sudo/vasiliy_bot/internal/telegram"
	log "github.com/todd-sudo/vasiliy_bot/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const token = "5207365782:AAFlMPp0i42AroASs6J381fdq49TjkT5BXs"

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Error(err)
	}
	bot.Debug = false

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Error(err)
	}
	log.Info("Бот запущен")
}
