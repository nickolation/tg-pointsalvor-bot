package ui

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	sdk "github.com/nickolation/pointsalvor"
)

type UiAdapter interface {
	Greet(chatId int64) error 
	ActiveSections(chatId int64, list []sdk.Section) error 
	ActiveTasks(chatId int64) error
	ErrorCommand(chatId int64) error
	NameModel(chaId int64) error
}


type Ui struct {
	api *tgbotapi.BotAPI
}

func NewUi(api *tgbotapi.BotAPI) *Ui {
	return &Ui{
		api: api,
	}
}
