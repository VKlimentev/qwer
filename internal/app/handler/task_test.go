package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"test_task_golang/internal/app/service"
	"test_task_golang/internal/model"
	"testing"

	mock_service "test_task_golang/internal/app/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockTaskService(ctrl)

	services := service.New(mockService)
	handler := Handler{services}

	router := gin.Default()
	router.POST("/task", handler.CreateTask)

	t.Run("Create task successfully", func(t *testing.T) {
		task := &model.Task{Method: "GET", URL: "http://example.com"}
		mockService.EXPECT().Create(task).Return("1", nil)

		body := strings.NewReader(`{"method":"GET","url":"http://example.com"}`)
		req, _ := http.NewRequestWithContext(context.Background(), "POST", "/task", body)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		expected := `{"id":"1"}`

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, expected, resp.Body.String())
	})

	t.Run("Create task with invalid JSON body", func(t *testing.T) {
		body := strings.NewReader(`{"method":"GET","url":"http://example.com"`)
		req, _ := http.NewRequestWithContext(context.Background(), "POST", "/task", body)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		expected := `{"error":"unexpected EOF"}`

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Equal(t, expected, resp.Body.String())
	})

	t.Run("Create task with service error", func(t *testing.T) {
		task := &model.Task{Method: "GET", URL: "http://example.com"}
		mockService.EXPECT().Create(task).Return("", errors.New("service error"))

		body := strings.NewReader(`{"method":"GET","url":"http://example.com"}`)
		req, _ := http.NewRequestWithContext(context.Background(), "POST", "/task", body)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		expected := `{"error":"service error"}`

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Equal(t, expected, resp.Body.String())
	})
}

func TestGetTaskStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockTaskService(ctrl)
	services := service.New(mockService)
	handler := Handler{services}

	router := gin.Default()
	router.GET("/task/:id", handler.GetTaskStatus)

	t.Run("Get task status successfully", func(t *testing.T) {
		taskStatus := &model.TaskStatus{ID: "1", Status: "done", HTTPStatusCode: 200}
		mockService.EXPECT().Get("1").Return(taskStatus, nil)

		req, _ := http.NewRequestWithContext(context.Background(), "GET", "/task/1", http.NoBody)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		expected := `{"id":"1","status":"done","httpStatusCode":200,"headers":null,"length":0}`

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, expected, resp.Body.String())
	})

	t.Run("Get task status with service error", func(t *testing.T) {
		mockService.EXPECT().Get("1").Return(nil, errors.New("service error"))

		req, _ := http.NewRequestWithContext(context.Background(), "GET", "/task/1", http.NoBody)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		expected := `{"error":"service error"}`

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Equal(t, expected, resp.Body.String())
	})
}
