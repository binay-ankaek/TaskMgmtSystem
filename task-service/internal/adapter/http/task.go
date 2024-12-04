package http

import (
	"context"

	"github.com/gin-gonic/gin"
	"net/http"
	"task-service/internal/domain"
	"time"
)

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	// Get the task from the request body
	timeoutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	var task domain.TaskCreateRequest
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if len(task.Assign_To) == 0 {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Assign to is required."})
	// 	return
	// }
	if task.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name is required."})
		return
	}
	if len(task.Email) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required."})
		return
	}
	// Create a new task
	response, err := h.svc.Task().CreateTask(timeoutCtx, &task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
