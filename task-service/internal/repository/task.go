package repository

import (
	"context"
	"database/sql"
	"fmt"
	"task-service/internal/domain"
	"task-service/internal/ports"

	"github.com/lib/pq"
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

func (t *TaskRepository) GetTask(ctx context.Context, req *domain.GetTaskRequest) ([]domain.GetTaskResponse, error) {
	//step:1 get user id by using email first
	var userID string
	err := t.db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", req.Email).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email: %s", req.Email)
		}
		return nil, fmt.Errorf("error querying user ID for email: %w", err)
	}

	// Step 2: Use the user ID to query tasks assigned to the user
	query := "SELECT id, name FROM tasks WHERE $1 = ANY(assign_to)"
	rows, err := t.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	// Step 3: Collect the tasks from the query result
	var tasks []domain.GetTaskResponse
	for rows.Next() {
		var task domain.GetTaskResponse
		err = rows.Scan(&task.TaskID, &task.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating rows: %w", err)
	}

	// Step 4: Check if any tasks were found
	if len(tasks) == 0 {
		return nil, fmt.Errorf("no tasks found for email: %s", req.Email)
	}

	return tasks, nil
}
