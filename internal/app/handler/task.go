package handler

import (
	"net/http"
	"test_task_golang/internal/model"

	"github.com/gin-gonic/gin"
)

// @Summary Create a task to fetch content from a URL
// @Description Creates a task to fetch content from the given URL
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Task Request payload"
// @Success 200 {integer} integer 1
// @Router /task [post]
func (h *Handler) CreateTask(c *gin.Context) {
	var t model.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.taskService.Create(&t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get the status of a task
// @Description Fetches the status of a task based on the provided task ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param taskId path string true "Task ID"
// @Success 200 {object} model.TaskStatus
// @Router /task/{taskId} [get]
func (h *Handler) GetTaskStatus(c *gin.Context) {
	id := c.Param("id")

	task, err := h.taskService.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
