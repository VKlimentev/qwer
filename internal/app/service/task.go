package service

import (
	"test_task_golang/internal/app/repository"
	"test_task_golang/internal/model"
)

type service struct {
	taskRepository repository.TaskRepository
}

func New(taskRepository repository.TaskRepository) *service {
	return &service{taskRepository: taskRepository}
}

func (s *service) Create(task *model.Task) (string, error) {
	return s.taskRepository.Create(task)
}

func (s *service) Get(taskID string) (*model.TaskStatus, error) {
	return s.taskRepository.Get(taskID)
}
