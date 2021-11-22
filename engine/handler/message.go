package handler

import (
	"context"
	"log"

	"github.com/nickolation/tg-pointsalvor-bot/auth"
)

func (hnd *Handler) HandleMessage(ctx context.Context, chatId int64, data string) error {
	temp, err := hnd.auth.SearchTemp(ctx, chatId)
	log.Printf("temp is - [%s]", temp)
	if err != nil {
		//		test-log 
		log.Printf("error with search temp - [%s]", err)

		return err
	}
	
	//delete temp value for updating by new row next time
	if err := hnd.auth.DeleteTemp(ctx, chatId); err != nil {
		return err
	}

	switch temp {
	case authTemp:
		hnd.auth.SignUp(ctx, &auth.KeyTokenOpt{
			ChatId: chatId,
			Token: data,
		})
	
	//other cases
	case sectionTemp:
		///
	case taskTemp:
		///
	}

	return nil
}


func (hnd *Handler) HandleForeignCommand(chatId int64) error {
	if err := hnd.ui.ErrorCommand(chatId); err != nil {
		return err
	}
	
	return nil
}
