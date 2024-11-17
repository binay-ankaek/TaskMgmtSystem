package repository

import (
	"context"
	"database/sql"
	"fmt"
	"task-service/internal/domain"
	"task-service/internal/ports"
)

type TaskRepository struct {
	db *sql.DB
}

type TaskGetter interface {
	Task() ports.TaskRepository
}

func (t *TaskRepository) Task() ports.TaskRepository {
	return t
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db,
	}
}

func (t *TaskRepository) CreateTask(ctx context.Context, req *domain.TaskCreate) (*domain.TaskCreate, error) {
	// Create a new task
	query := "INSERT into tasks (name,assign_to) VALUES($1,$2)  RETURNING id"

	err := t.db.QueryRowContext(ctx, query, req.Name, req.Assign_To).Scan(req.TaskID)
	if err != nil {
		return nil, err
	}
	fmt.Println(req)
	return req, nil
}
