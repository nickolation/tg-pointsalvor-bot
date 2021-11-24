package ui

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	sdk "github.com/nickolation/pointsalvor"
)

//		caption value for handling
var (	
	//		zero section list
	ActiveListZS = "Похоже, у вас нет ни одного раздела\n" + 
		"Нажмите на +, чтобы его добавить, и введите название\n"

	//		full section list 
	ActiveListFS = "Активные разделы\n" + 
		"Нажмите на раздел, чтобы посмотреть задачи\n"

	//		zero task list
	ActiveListZT = "Список задач пуст\n" + 
		"Чтобы добавить, нажмите на + и введите описание задачи\n"

	//		full task list
	ActiveListFT = "Активные задачи\n"
)

//action high value 
const (
	ViewAction = "view"
	WriteAction = "write"
	CloseAction = "close"
)


const (
	//may change to smile or text
	addText = "+"
)


//string contains the action and name of work model 
//constucted by two parts: "action:name"
func CallbackData(action string, name string) string {
	return action + ":" + name
}

//		untested
//makes keyboard with section button and then seng this to the chat
func (u *Ui) ActiveSections(chatId int64, list []sdk.Section) error  {
	//button matrix 
	listRows := make([][]tgbotapi.InlineKeyboardButton, len(list) + 1)

	for idx, val := range list {
		name := val.Name

		//simple section buttons with data in sort of "view:section_name"
		listRows[idx] = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(name, CallbackData(ViewAction, name)))
	}

	//last row is the add button
	//data of add button is "write:+"
	addRow := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(addText, CallbackData(WriteAction, addText)))
	listRows = append(listRows, addRow)


	//section active list keyboard 
	keyBoard := tgbotapi.NewInlineKeyboardMarkup(listRows...)



	//validate state and make new message with markup
	msg := tgbotapi.NewMessage(0, "")
	if len(list) == 0 {
		msg = tgbotapi.NewMessage(chatId, ActiveListFS)
	} else {
		msg = tgbotapi.NewMessage(chatId, ActiveListZS)
	}

	msg.ReplyMarkup = keyBoard

	if _, err := u.api.Send(msg); err != nil {
		//		test-log
		log.Printf("error with makae the markup - [%s]", err)
		return err
	}
	
	return nil
}

func (u *Ui) ActiveTasks(chatId int64) error {
	return nil
}