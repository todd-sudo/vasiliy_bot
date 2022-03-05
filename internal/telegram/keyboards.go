package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var deleteKeyboard = tgbotapi.NewInlineKeyboardMarkup(

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Удалить", "del"),
	),
)
