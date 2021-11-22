package ui

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	greetingText = "Привет, я бот-помощник, который работает на основе Todoist\n" +
		"Я облегчу твою работу с задачами, испольуя простую логику разделов и доски\n" +
		"Распредели свои задачи по темам, а я размещу их на доске и помогу с ними работать\n" +
		"Всё намного проще, чем ты думаешь, просто следуй моим инструкциям \xF0\x9F\x98\x89"

	tokenText = "Чтобы продолжить работу, мне нужен доступ к твоему Todoist аккаунту\n" +
		"Для этого напиши свой токен интеграции, который можно найти в настройках\n" +
		"Например, ff16c7b22cdb650c4d14679b20cc4fdef0fc1264"

	errCommandText = "Я не знаю твою комманду :(\n" +
		"Может быть, ты имел в виду комманду /start ?\n" + 
		"Введи её с новой строки"
)


//message to start command handler
func (u *Ui) Greet(chatId int64) error {
	greetMsg := tgbotapi.NewMessage(chatId, greetingText)
	tokenMsg := tgbotapi.NewMessage(chatId, tokenText)

	if _, err := u.api.Send(greetMsg); err != nil {
		return err
	}

	if _, err := u.api.Send(tokenMsg); err != nil {
		return err
	}

	return nil
}


func (u *Ui) ActiveSections(chatId int64) error  {
	return nil
}

func (u *Ui) ActiveTasks(chatId int64) error {
	return nil
}

func (u *Ui) ErrorCommand(chatId int64) error {
	errMsg := tgbotapi.NewMessage(chatId, errCommandText)

	if _, err := u.api.Send(errMsg); err != nil {
		return err
	}
	
	return nil
}