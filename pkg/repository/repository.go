package repository

import (
	task "test_task_golang"
)

type Task interface {
	Create(task task.Task) (string, error)
	GetById(taskId string) (*task.TaskStatus, error)
}

type Repository struct {
	Task
}

func NewRepository() *Repository {
	return &Repository{
		Task: NewTaskRepository(),
	}
}
