package service

import (
	task "test_task_golang"
	"test_task_golang/pkg/repository"
)

type TaskService struct {
	repos repository.Task
}

func NewTaskService(r *repository.Repository) *TaskService {
	return &TaskService{repos: r}
}

func (s *TaskService) Create(task task.Task) (string, error) {
	return s.repos.Create(task)
}

func (s *TaskService) GetById(taskId string) (*task.TaskStatus, error) {
	return s.repos.GetById(taskId)
}
