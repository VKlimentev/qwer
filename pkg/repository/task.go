package repository

import (
	"fmt"
	"io"
	"net/http"

	task "test_task_golang"

	"github.com/google/uuid"
)

type TaskRepository struct {
	tasks map[string]*task.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]*task.Task),
	}
}

func (r *TaskRepository) Create(t task.Task) (string, error) {
	id := uuid.New().String()
	t.Status = &task.TaskStatus{
		ID:     id,
		Status: "in_process",
	}
	r.tasks[id] = &t

	go r.executeTask(&t)

	return id, nil
}

func (r *TaskRepository) GetById(taskID string) (*task.TaskStatus, error) {
	task, ok := r.tasks[taskID]
	if !ok {
		return nil, fmt.Errorf("task not found")
	}
	return task.Status, nil
}

func (r *TaskRepository) executeTask(task *task.Task) {
	client := &http.Client{}

	request, err := http.NewRequest(task.Method, task.URL, nil)
	if err != nil {
		task.Status.Status = "error"
		return
	}

	for key, value := range task.Headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		task.Status.Status = "error"
		return
	}
	defer response.Body.Close()

	task.Status.HTTPStatusCode = response.StatusCode
	headers := make(map[string]string)
	for key, value := range response.Header {
		headers[key] = value[0]
	}
	task.Status.Headers = headers

	body, err := io.ReadAll(response.Body)
	if err != nil {
		task.Status.Status = "error"
		return
	}

	task.Status.Length = int(len(body))
	task.Status.Status = "done"
}
