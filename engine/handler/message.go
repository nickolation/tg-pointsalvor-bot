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
