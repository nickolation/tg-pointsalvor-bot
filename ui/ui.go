package ui

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type UiAdapter interface {
	Greet(chatId int64) error 
	ActiveSections(chatId int64) error 
	ActiveTasks(chatId int64) error
	ErrorCommand(chatId int64) error
}


type Ui struct {
	api *tgbotapi.BotAPI
}

func NewUi(api *tgbotapi.BotAPI) *Ui {
	return &Ui{
		api: api,
	}
}
