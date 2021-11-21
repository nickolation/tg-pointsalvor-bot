package engine

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	//sdk "github.com/nickolation/pointsalvor"
	"github.com/nickolation/tg-pointsalvor-bot/engine/handler"
	"github.com/nickolation/tg-pointsalvor-bot/ui"
)


var (
	//parse
	errNilMessageData = errors.New("nil message data: parse message is locked")
	errNilMessageChatId = errors.New("nil message chatId: parse message is locked")
)


type EngineBot struct {
	bot     *tgbotapi.BotAPI
	handler handler.HandlerAdapter	
}

func NewEngineBot(bot *tgbotapi.BotAPI, hnd handler.HandlerAdapter, ui *ui.Ui) *EngineBot {
	return &EngineBot{
		bot:     bot,
		handler: hnd,
	}
}

type MessageOpt struct {
	//inner data of telegram message
	Data string 

	//chatId of user 
	ChatId int64
}

//		???
func ParseMessage(msg *tgbotapi.Message) (*MessageOpt, error) {
	data := msg.Text
	if data == "" {
		return nil, errNilMessageData
	}

	chatId := msg.Chat.ID
	if chatId == 0 {
		return nil, errNilMessageChatId
	}

	return &MessageOpt{
		Data: data,
		ChatId: chatId,
	}, nil
}
