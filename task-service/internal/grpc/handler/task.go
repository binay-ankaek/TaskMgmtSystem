package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"task-service/internal/app"
	"task-service/internal/domain"
	pb "task-service/proto/task"
	"time"
)

type TaskHandler struct {
	pb.UnimplementedTaskServiceServer
	hand app.TaskServiceGetter
}

func NewTaskHandler(hand app.TaskServiceGetter) *TaskHandler {
	return &TaskHandler{
		hand: hand,
	}
}

func (h *TaskHandler) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	// Call the task service to get the task
	// Get the task from the request body
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	task, err := h.hand.Task().GetTask(timeoutCtx, &domain.GetTaskRequest{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	if len(task) == 0 {
		return nil, status.Errorf(codes.NotFound, "task not found for email :%s", req.Email)
	}
	//map the task
	var pbTasks []*pb.GetTask
	// Return the task
	for _, task := range task {
		pbTasks = append(pbTasks, &pb.GetTask{
			Id:   task.TaskID,
			Name: task.Name,
		})

	}
	return &pb.GetTaskResponse{
		Tasks: pbTasks,
	}, nil

}
