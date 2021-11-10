package repository

import "github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"

type Tasks interface {
	CasheTask()
	//?
	SelectTask()
	RemoveTask()
	CasheCloseTask()
	SelectAllTasks()
	SelectFavoriteTasks()
	//other
}

type Sections interface {
	//?
	SelectSection()
	CasheSection()
	RemoveSection()
	SelectAllSection()
	SelectLastSection()
	//other
}

type Tables interface {
	//?
	RewriteTable()
	//other
}

type Repository struct {
	Tasks
	Sections
	Tables
}

func NewRepository(store *storage.BotStorage) *Repository {
	return &Repository{
		Tasks:    NewTaskRepo(store),
		Sections: NewSectionRepo(store),
		Tables:   NewTableRepo(store),
	}
}
