package repository

import "github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"

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

func (sr *SectionRepo) CasheSection() {

}
