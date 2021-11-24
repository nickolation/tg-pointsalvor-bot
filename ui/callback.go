package ui

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)


var (
	nameModelText = "Введите название /смайлик/"
)


func (u *Ui) NameModel (chatId int64) error {
	nameMsg := tgbotapi.NewMessage(chatId, nameModelText)


	if _, err := u.api.Send(nameMsg); err != nil {
		//	test-log 
		fmt.Printf("send message error - [%s]", err)
		return err
	}

	return nil
}