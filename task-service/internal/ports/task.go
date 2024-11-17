package ports

import (
	"context"
	"task-service/internal/domain"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, req *domain.TaskCreate) (*domain.TaskCreate, error)
}

type TaskService interface {
	// create task
	CreateTask(ctx context.Context, req *domain.TaskCreateRequest) (*domain.TaskCreateResponse, error)
}
