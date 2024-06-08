package handler

import (
	"net/http"
	task "test_task_golang"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTask(c *gin.Context) {
	var task task.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Create(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) GetTaskStatus(c *gin.Context) {
	id := c.Param("id")

	task, err := h.services.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
