package service

import (
	"test_task_golang/internal/model"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type TaskService interface {
	Create(task *model.Task) (string, error)
	Get(taskID string) (*model.TaskStatus, error)
}
