package models

type Task struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Status  *TaskStatus       `json:"status"`
}

type TaskStatus struct {
	ID             string            `json:"id"`
	Status         string            `json:"status"`
	HTTPStatusCode int               `json:"httpStatusCode"`
	Headers        map[string]string `json:"headers"`
	Length         int               `json:"length"`
}
