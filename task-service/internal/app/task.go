package app

import (
	"context"
	"errors"
	"task-service/internal/domain"
	"task-service/internal/ports"
	"task-service/internal/repository"
)

type TaskService struct {
	repo repository.TaskGetter
}

type TaskServiceGetter interface {
	Task() ports.TaskService
}

func (t *TaskService) Task() ports.TaskService {
	return t
}

func NewTaskService(repo repository.TaskGetter) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) CreateTask(ctx context.Context, req *domain.TaskCreateRequest) (*domain.TaskCreateResponse, error) {
	//check either empty or not

	if req == nil {
		return nil, errors.New("empty request")
	}
	if req.Name == "" {
		return nil, errors.New("empty name")
	}
	if len(req.Assign_To) == 0 {
		return nil, errors.New("empty assign to")
	}
	newTask := &domain.TaskCreate{
		Name:      req.Name,
		Assign_To: req.Assign_To,
	}

	//create task
	task, err := t.repo.Task().CreateTask(ctx, newTask)
	if err != nil {
		return nil, err
	}
	response := &domain.TaskCreateResponse{
		TaskID:    task.TaskID,
		Name:      task.Name,
		Assign_To: task.Assign_To,
	}
	return response, nil
}
