package repository

import (
	"context"
	"log"

	"github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"
)

const (
	Tmodel = "task"
	Smodel = "section"

	tempModel = "temp"
)


type Committer struct {
	db *storage.BotStorage
}


func NewCommiter(db *storage.BotStorage) *Committer {
	return &Committer{
		db: db,
	}
}


func (com *Committer) CommitWrite(ctx context.Context, chatId int64, model string) error {
	key, err := makeKey(tempModel, chatId)
	if err != nil {
		//		test-log
		log.Printf("error with make temp key - [%s]", err)
		return err
	}

	switch model{
	case Tmodel:
		if _, err := com.db.Client.Set(ctx, key, sectionModel, 0).Result(); err != nil {
			//		test-log
			log.Printf("error with make temp row [S] in redis - [%s]", err)
			return err
		}

	case Smodel:
		if _, err := com.db.Client.Set(ctx, key, taskModel, 0).Result(); err != nil {
			//		test-log
			log.Printf("error with make temp row [T] in redis - [%s]", err)
			return err
		}
	}

	return nil
}


func (com *Committer)  CommitClose(ctx context.Context, chatId int64) error {


	return nil
}


