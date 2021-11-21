package handler

import (
	"context"
	"log"
)

const (
	authTemp = "auth"
	sectionTemp = "section"
	taskTemp = "task"
)

func (hnd *Handler) HandleStart(ctx context.Context, chatId int64) error {
	_, err := hnd.auth.SignIn(ctx, chatId)
	if err != nil {
		if err := hnd.auth.MakeTemp(ctx, chatId, authTemp); err != nil {
			//		test-log
			log.Printf("error with message - [%s]", err.Error())
			return err
		}
	}

	return nil
}