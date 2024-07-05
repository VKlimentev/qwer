package repository

import (
	"test_task_golang/internal/model"
)

type TaskRepository interface {
	Create(task *model.Task) (string, error)
	Get(taskID string) (*model.TaskStatus, error)
}
