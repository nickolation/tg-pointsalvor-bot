package service

import (
	"context"

	sdk "github.com/nickolation/pointsalvor"
	"github.com/nickolation/tg-pointsalvor-bot/auth"
	"github.com/nickolation/tg-pointsalvor-bot/engine/external/repository"
)

//action wich can be performed under the callback-signal from user 
//list -> write model, view model, close task model
type Actions interface {
	//section / task 
	ActWrite(ctx context.Context, chatId int64, model string) error
	
	//section only 
	ActView(ctx context.Context, chatId int64) error

	//clsoe task only
	ActClose(ctx context.Context, chatId int64) error
}


type Tasks interface {
	AddTask()
	DeleteTask()
	CloseTask()
	GetAllTasks()
	GetAllFavoriteTasks()
	//other
}

type Sections interface {
	AddSection(ctx context.Context, chatId int64, name string) error
	DeleteSection()
	GetAllSections(ctx context.Context, chatId int64) ([]sdk.Section, error)
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
	Actions
}

func NewService(a auth.Auth, r repository.Repository) *Service {
	return &Service{
		Tasks:    NewTaskService(&r.Tasks),
		Sections: NewSectionService(r.Sections, a),
		Tables:   NewTableService(&r.Tables),
		Actions: NewActor(&r.Commits),
	}
}
