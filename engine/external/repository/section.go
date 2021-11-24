package repository

import (
	"context"
	"encoding/json"
	"log"

	sdk "github.com/nickolation/pointsalvor"
	"github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"
)


const (
	sectionModel = "section"
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


func (sr *SectionRepo) SelectAllSection(ctx context.Context, chatId int64) ([]sdk.Section, error) {
	matcher := getIdentityMatcher(sectionModel, chatId)
	iter := sr.db.Client.Scan(ctx, 0, matcher, 0).Iterator()
	if err := iter.Err(); err == nil {
		//		test-log
		log.Printf("error with the scanner - [%s]", err.Error())

		return nil, err
	}

	sectionList := []sdk.Section{}

	for iter.Next(ctx) {
		key := iter.Val()
		val, err := sr.db.Client.Get(ctx, key).Result()
		if err != nil {
			//		test-log
			log.Printf("error with getting value - [%s]", err.Error())
			
			return nil, err
		}

		input := &sdk.Section{}

		//		[]byte???
		if err = json.Unmarshal([]byte(val), input); err != nil {
			//		test-log
			log.Printf("unmarshal problem - [%s]", err.Error())

			return nil, err
		}

		sectionList = append(sectionList, *input)
	}

	return sectionList, nil
}

func (sr *SectionRepo) SelectLastSection() {

}

func (sr *SectionRepo) CasheSection(ctx context.Context, chatId int64, s *sdk.Section) error {
	jsonData, err :=  json.Marshal(s)
	if err != nil {
		return err
	}

	key, err := makeKey(sectionModel, chatId)
	if err != nil {
		return err
	}

	//		??? []byte
	_, err = sr.db.Client.Set(ctx, key, string(jsonData), 0).Result()
	if err != nil {
		//		test-log
		log.Printf("error with set section key - [%s]", err.Error())
		return err
	}
	
	return nil
}
