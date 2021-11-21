package repository

import "github.com/nickolation/tg-pointsalvor-bot/engine/external/storage"


const (
	taskKey = "task!!chat_id:%s"
)

type TaskRepo struct {
	db *storage.BotStorage
}

func NewTaskRepo(db *storage.BotStorage) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (tsr *TaskRepo) CasheTask() {

}

func (tsr *TaskRepo) RemoveTask() {

}

func (tsr *TaskRepo) SelectTask() {

}

func (tsr *TaskRepo) CasheCloseTask() {

}

func (tsr *TaskRepo) SelectAllTasks() {

}

func (tsr *TaskRepo) SelectFavoriteTasks() {

}
