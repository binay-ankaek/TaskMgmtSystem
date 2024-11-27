package repository

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
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
	var taskid string

	err := t.db.QueryRowContext(ctx, query, req.Name, pq.Array(req.Assign_To)).Scan(&taskid)
	if err != nil {
		return nil, err
	}
	response := &domain.TaskCreate{
		TaskID:    taskid,
		Name:      req.Name,
		Assign_To: req.Assign_To,
	}
	return response, nil
}
