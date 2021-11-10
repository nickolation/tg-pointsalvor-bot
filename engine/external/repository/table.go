package repository

import "github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"

type TableRepo struct {
	db *storage.BotStorage
}

func NewTableRepo(db *storage.BotStorage) *TableRepo {
	return &TableRepo{
		db: db,
	}
}

func (tr *TableRepo) RewriteTable() {

}
