package model

type TaskResponse struct {
	TaskID   string `json:"task_id"`
	TaskName string `json:"task_name"`
}
type TaskRequest struct {
	Email string `json:"email"`
}
type TaskResponses struct {
	TaskResponses []TaskResponse `json:"task_responses"`
}
