package service

import (
	task "test_task_golang"
	"test_task_golang/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Task interface {
	Create(task task.Task) (string, error)
	GetById(taskId string) (*task.TaskStatus, error)
}

type Service struct {
	Task
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Task: NewTaskService(r),
	}
}
