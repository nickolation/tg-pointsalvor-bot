package handler

import (
	"context"
	"log"
	"strings"

	"github.com/nickolation/tg-pointsalvor-bot/ui"
)

const (
	tModel = "task"
	sModel = "section"
)

type CallbackOpt struct {
	//for choice of the action - add/select
	Action string 

	//for validation model - section/task
	NameModel string
}

func parseToCallbackOpt(data string) CallbackOpt {
	//index of :
	i := strings.Index(data, ":")

	return CallbackOpt{
		Action: data[:i],
		NameModel: data[i + 1:],
	}
}

//		untested
func (hnd *Handler) HandleCallback(ctx context.Context, chatId int64, data string) error {
	//		test-log
	log.Printf("action and name mode (data) - [%s]", data)
	if data == "" {
		return errNilOpt
	}

	//parse opts 
	opt := parseToCallbackOpt(data)
	action := opt.Action 

	//		later
	// name := opt.NameModel

	//handling actions before user's message with model name or new markup
	switch action {
	case ui.WriteAction:
		//write temp - setion row in redis 
		if err := hnd.svice.Actions.ActWrite(ctx, chatId, sModel); err != nil {
			return err
		}

		//send tick to the chat 
		if err := hnd.ui.NameModel(chatId); err != nil {
			return err
		}

	case ui.CloseAction:
		///
	case ui.ViewAction:
		///
	}



	return nil
}
