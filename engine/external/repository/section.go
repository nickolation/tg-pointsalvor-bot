package repository

import (
	"context"
	"encoding/json"
	"log"

	sdk "github.com/nickolation/pointsalvor"
	"github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"
)


const (
	sectionKey = "section!!chat_id:%d"
)

var (
)


type SectionRepo struct {
	db *storage.BotStorage
}

func NewSectionRepo(db *storage.BotStorage) *SectionRepo {
	return &SectionRepo{
		db: db,
	}
}

func (sr *SectionRepo) SelectSection() {

}

func (sr *SectionRepo) RemoveSection() {

}

func (sr *SectionRepo) SelectAllSection() {

}

func (sr *SectionRepo) SelectLastSection() {

}

func (sr *SectionRepo) CasheSection(ctx context.Context, chatId int64, s *sdk.Section) error {
	jsonData, err :=  json.Marshal(s)
	if err != nil {
		return err
	}

	key, err := makeKey(sectionKey, chatId)
	if err != nil {
		return err
	}

	//		??? []byte
	_, err = sr.db.Client.Set(ctx, key, string(jsonData), 0).Result()
	if err != nil {
		//		test-log
		log.Print("error with set section key - [%s]", err.Error())
		return err
	}
	
	return nil
}
