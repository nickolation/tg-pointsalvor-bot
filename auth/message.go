package auth

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type authMessageInterface interface {
	t() error
}

type authMessage struct {
	tgapi *tgbotapi.BotAPI
}

func newAuthMessage(tgapi *tgbotapi.BotAPI) *authMessage {
	return &authMessage{
		tgapi: tgapi,
	}
}

func (msg *authMessage) t() error {
	return nil
}
