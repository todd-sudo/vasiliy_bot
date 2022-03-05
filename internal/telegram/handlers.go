package telegram

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/todd-sudo/vasiliy_bot/pkg/logger"
)

const ()

func (b *Bot) handleMessage(update *tgbotapi.Update) error {
	msgU := fmt.Sprintf("Здарова @%s", update.Message.From.UserName)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgU)
	if strings.Contains(update.Message.Text, "ривет") {
		msg.ReplyMarkup = deleteKeyboard
		msgBot, err := b.bot.Send(msg)
		if err != nil {
			log.Error(err)
			return err
		}

		time.Sleep(10 * time.Second)
		del := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, msgBot.MessageID)
		b.bot.Request(del)

		// msgConf.

	}
	return nil
}
