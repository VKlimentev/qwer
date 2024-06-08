package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	task "test_task_golang"
	"test_task_golang/pkg/service"
	mock_service "test_task_golang/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockTask(ctrl)

	services := &service.Service{Task: mockService}
	handler := &Handler{services}

	router := gin.Default()
	router.POST("/task", handler.CreateTask)

	t.Run("Create task successfully", func(t *testing.T) {
		task := task.Task{Method: "GET", URL: "http://example.com"}
		mockService.EXPECT().Create(task).Return("1", nil)

		body := strings.NewReader(`{"method":"GET","url":"http://example.com"}`)
		req, _ := http.NewRequest("POST", "/task", body)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, `{"id":"1"}`, resp.Body.String())
	})

	t.Run("Create task with invalid JSON body", func(t *testing.T) {
		body := strings.NewReader(`{"method":"GET","url":"http://example.com"`)
		req, _ := http.NewRequest("POST", "/task", body)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Equal(t, `{"error":"unexpected EOF"}`, resp.Body.String())
	})

	t.Run("Create task with service error", func(t *testing.T) {
		task := task.Task{Method: "GET", URL: "http://example.com"}
		mockService.EXPECT().Create(task).Return("", errors.New("service error"))

		body := strings.NewReader(`{"method":"GET","url":"http://example.com"}`)
		req, _ := http.NewRequest("POST", "/task", body)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Equal(t, `{"error":"service error"}`, resp.Body.String())
	})
}

func TestGetTaskStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockTask(ctrl)

	services := &service.Service{Task: mockService}
	handler := &Handler{services}

	router := gin.Default()
	router.GET("/task/:id", handler.GetTaskStatus)

	t.Run("Get task status successfully", func(t *testing.T) {
		taskStatus := &task.TaskStatus{ID: "1", Status: "done", HTTPStatusCode: 200}
		mockService.EXPECT().GetById("1").Return(taskStatus, nil)

		req, _ := http.NewRequest("GET", "/task/1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, `{"id":"1","status":"done","httpStatusCode":200,"headers":null,"length":0}`, resp.Body.String())
	})

	t.Run("Get task status with service error", func(t *testing.T) {
		mockService.EXPECT().GetById("1").Return(nil, errors.New("service error"))

		req, _ := http.NewRequest("GET", "/task/1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Equal(t, `{"error":"service error"}`, resp.Body.String())
	})
}
