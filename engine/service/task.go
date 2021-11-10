package service

import "github.com/nickolation/tg-pointsalvor-bot/engine/external/repository"

type TaskService struct {
	repo *repository.Tasks
}

func NewTaskService(repo *repository.Tasks) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (tks *TaskService) AddTask() {

}

func (tks *TaskService) DeleteTask() {

}

func (tks *TaskService) CloseTask() {

}

func (tks *TaskService) GetAllTasks() {

}

func (tks *TaskService) GetAllFavoriteTasks() {

}
