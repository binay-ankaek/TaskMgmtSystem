package ports

import (
	"context"
	"task-service/internal/domain"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, req *domain.TaskCreate) (*domain.TaskCreate, error)
	GetTask(ctx context.Context, req *domain.GetTaskRequest) ([]domain.GetTaskResponse, error)
}

type TaskService interface {
	// create task
	CreateTask(ctx context.Context, req *domain.TaskCreateRequest) (*domain.TaskCreateResponse, error)
	// get task
	GetTask(ctx context.Context, req *domain.GetTaskRequest) ([]domain.GetTaskResponse, error)
}
