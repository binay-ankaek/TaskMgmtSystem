package domain

import ()

type TaskCreateRequest struct {
	Name      string   `json:"name"`
	Assign_To []string `json:"assign_to"`
}

type TaskCreate struct {
	TaskID    string   `json:"task_id,omitempty"`
	Name      string   `json:"name"`
	Assign_To []string `json:"assign_to"`
}
type TaskCreateResponse struct {
	TaskID    string   `json:"task_id"`
	Name      string   `json:"name"`
	Assign_To []string `json:"assign_to"`
}
