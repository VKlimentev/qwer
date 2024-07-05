package repository

import (
	"test_task_golang/internal/model"
)

type TaskRepository interface {
	Create(task *model.Task) (string, error)
	Get(taskId string) (*model.TaskStatus, error)
}
