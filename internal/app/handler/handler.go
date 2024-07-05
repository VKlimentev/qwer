package handler

import (
	"test_task_golang/internal/app/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	taskService service.TaskService
}

func New(taskService service.TaskService) *Handler {
	return &Handler{taskService: taskService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	// add swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/task", h.CreateTask)
	r.GET("/task/:id", h.GetTaskStatus)

	return r
}
