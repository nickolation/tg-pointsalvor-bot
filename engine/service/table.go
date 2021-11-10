package service

import "github.com/nickolation/tg-pointsalvor-bot/engine/external/repository"

type TableService struct {
	repo *repository.Tables
}

func NewTableService(repo *repository.Tables) *TableService {
	return &TableService{
		repo: repo,
	}
}

func (ts *TableService) ReopenTable() {

}
