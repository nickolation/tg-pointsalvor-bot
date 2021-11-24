package handler

import (
	"context"
	"log"

	"github.com/nickolation/tg-pointsalvor-bot/auth"
)

//		untested 
// check out the state of the user's lists -> generate need ui
func (hnd *Handler) HandleStateSList(ctx context.Context, chatId int64) error {
	list, err := hnd.svice.GetAllSections(ctx, chatId)
	if err != nil {
		return err
	}

	if err = hnd.ui.ActiveSections(chatId, list); err != nil {
		return err
	}

	return nil
}

//		same but with the task
func (hnd *Handler) HandleStateTList(ctx context.Context, chatId int64) error {
	
	
	return nil
}


//		untested 
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

		//send active list for as the first action
		//for zero list -> empty active list with caption 
		//for full list -> the simple acive list
		if err := hnd.HandleStateSList(ctx, chatId); err != nil {
			return err
		}
	
	//other cases
	case sectionTemp:
		//user seng message with section name to create new section 
		//add new section to api and send new markup
		if err := hnd.svice.AddSection(ctx, chatId, data); err != nil {
			return err
		}

		if err := hnd.HandleStateSList(ctx, chatId); err != nil {
			return err
		}
		
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
