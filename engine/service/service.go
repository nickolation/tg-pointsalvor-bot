package service

import (
	"github.com/nickolation/tg-pointsalvor-bot/auth"
	"github.com/nickolation/tg-pointsalvor-bot/engine/external/repository"
)

type Tasks interface {
	AddTask()
	DeleteTask()
	CloseTask()
	GetAllTasks()
	GetAllFavoriteTasks()
	//other
}

type Sections interface {
	AddSection()
	DeleteSection()
	GetAllSections()
	GetLastSection()
	//other
}

type Tables interface {
	ReopenTable()
	//other
}

type Service struct {
	Tasks
	Sections
	Tables
}

func NewService(auth *auth.AuthEngine, repo *repository.Repository) *Service {
	return &Service{
		Tasks:    NewTaskService(&repo.Tasks),
		Sections: NewSectionService(&repo.Sections),
		Tables:   NewTableService(&repo.Tables),
	}
}
