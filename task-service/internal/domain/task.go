package domain

import ()

type TaskCreateRequest struct {
	Name  string   `json:"name"`
	Email []string `json:"emails"`
}

type TaskCreate struct {
	TaskID    string   `json:"task_id,omitempty"`
	Name      string   `json:"name"`
	Assign_To []string `json:"assign_to"`
}
type TaskCreateResponse struct {
	TaskID    string        `json:"task_id"`
	Name      string        `json:"name"`
	Assign_To []UserDetails `json:"assign_to"`
}

type UserDetails struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
type GetTaskRequest struct {
	Email string `json:"email"`
}

type GetTaskResponse struct {
	TaskID string `json:"task_id"`
	Name   string `json:"name"`
}
