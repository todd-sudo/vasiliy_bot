package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/todd-sudo/vasiliy_bot/pkg/logger"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

// NewBot constructor
func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

// Start bot
func (b *Bot) Start() error {
	log.Infof("Authorized on account %s", b.bot.Self.UserName)
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}
	b.handlerUpdates(updates)

	return nil
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil
}

func (b *Bot) handlerUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {
		if update.Message != nil {
			go b.handleMessage(&update)

		} else {
			log.Infof("Updating")
		}
	}
}
