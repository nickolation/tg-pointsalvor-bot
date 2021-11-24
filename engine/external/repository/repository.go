package repository

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/nickolation/pointsalvor"
	"github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"
)

const (
	modelMatcher = "%s!!chat_id:%d*"
	keyTemplate = "%s!!chat_id:%d"
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


type Commits interface {
	CommitWrite(ctx context.Context, chatId int64, model string) error 

	CommitClose(ctx context.Context, chatId int64) error
}


type Sections interface {
	//?
	SelectSection()

	//cashe section json body to redis 
	CasheSection(ctx context.Context, chatId int64, s *sdk.Section) error

	//CheckStateSection(ctx context.Context, chatId int64) 

	RemoveSection()

	//getting all section for unmarshling inner body data and validate sections state 
	SelectAllSection(ctx context.Context, chatId int64) ([]sdk.Section, error)

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
	Commits
}

func NewRepository(store *storage.BotStorage) *Repository {
	return &Repository{
		Tasks:    NewTaskRepo(store),
		Sections: NewSectionRepo(store),
		Tables:   NewTableRepo(store),
		Commits: NewCommiter(store),
	}
}

//make key allows the template - one the key structure 
func makeKey(template string, chatId int64) (string, error) {
	if chatId == 0 {
		return "", errNilChatId
	} 

	return fmt.Sprintf(keyTemplate, template, chatId), nil
}


//generate identity matcher for scanner to search keys in the redis
func getIdentityMatcher(model string, chatId int64) string {
	return fmt.Sprintf(modelMatcher, model, chatId)
}


