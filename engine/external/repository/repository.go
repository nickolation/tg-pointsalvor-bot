package repository

import (
	"errors"
	"fmt"

	"github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"
)


var (
	errNilChatId = errors.New("nil chat id value - key generation isn't valid")
)


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

//make key allows the template - one the key structure 
func makeKey(template string, chatId int64) (string, error) {
	if chatId == 0 {
		return "", errNilChatId
	} 

	return fmt.Sprintf(sectionKey, chatId), nil
}


