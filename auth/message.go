package auth

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	alreadyAuthText = "Вы уже зарегистрированы. Давайте приступим к работе"

	succesAuthText = "Успешная регистрация, ваша доска готова к работе"
)

type authMessageAdapter interface {
	AlreadyAuthorized(chatId int64) error
	SuccesAuthorized(chatId int64) error
}

type authMessage struct {
	api *tgbotapi.BotAPI
}

func newAuthMessage(api *tgbotapi.BotAPI) *authMessage {
	return &authMessage{
		api: api,
	}
}


//message to user caption-callback -- allready 
func (msg *authMessage) AlreadyAuthorized(chatId int64) error {
	authMsg := tgbotapi.NewMessage(chatId, alreadyAuthText)

	if _, err := msg.api.Send(authMsg); err != nil {
		return err
	}

	return nil
}


//message to user caption-callback -- succes 
func (msg *authMessage) SuccesAuthorized(chatId int64) error {
	authMsg := tgbotapi.NewMessage(chatId, succesAuthText)

	if _, err := msg.api.Send(authMsg); err != nil {
		return err
	}

	return nil
}
