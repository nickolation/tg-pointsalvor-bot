package ui

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Ui struct {
	api *tgbotapi.BotAPI
}

func NewUi(api *tgbotapi.BotAPI) *Ui {
	return &Ui{
		api: api,
	}
}
