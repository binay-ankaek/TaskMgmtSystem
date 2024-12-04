package app

import (
	"context"
	"errors"
	"google.golang.org/grpc/status"
	"task-service/internal/domain"
	"task-service/internal/ports"
	"task-service/internal/repository"
	pbUser "task-service/proto/user"
)

type TaskService struct {
	repo         repository.TaskGetter
	pbUserClient pbUser.UserServiceClient
}

type TaskServiceGetter interface {
	Task() ports.TaskService
}

func (t *TaskService) Task() ports.TaskService {
	return t
}

func NewTaskService(repo repository.TaskGetter, pbUserClient pbUser.UserServiceClient) *TaskService {
	return &TaskService{
		repo:         repo,
		pbUserClient: pbUserClient,
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
	if req.Email == nil {
		return nil, errors.New("empty email ")
	}

	// if len(req.Assign_To) == 0 {
	// 	return nil, errors.New("empty assign to")
	// }
	// user, err := t.pbUserClient.GetAllUser(ctx, &pbUser.GetAllUserRequest{})
	// if err != nil {
	// 	return nil, err
	// }
	// if len(user.Users) == 0 {
	// 	return nil, errors.New("no user")
	// }
	// //loop the user data to get userID and name
	// assign_to := []domain.UserDetails{}
	// for _, user := range user.Users {
	// 	assign_to = append(assign_to, domain.UserDetails{
	// 		ID:   user.Id,
	// 		Name: user.Name,
	// 	})
	// }
	assign_to := []domain.UserDetails{}

	for _, email := range req.Email {
		user, err := t.pbUserClient.GetUser(ctx, &pbUser.GetUserRequest{
			Email: email,
		})
		if err != nil {
			grpcStatus, ok := status.FromError(err)
			if ok {
				return nil, errors.New(grpcStatus.Message())
			}
			return nil, err
		}
		assign_to = append(assign_to, domain.UserDetails{
			ID:   user.Id,
			Name: user.Name,
		})
	}
	if len(assign_to) == 0 {
		return nil, errors.New("no user")
	}

	newTask := &domain.TaskCreate{
		Name:      req.Name,
		Assign_To: extractUserIds(assign_to),
	}

	//create task
	task, err := t.repo.Task().CreateTask(ctx, newTask)
	if err != nil {
		return nil, err
	}

	response := &domain.TaskCreateResponse{
		TaskID:    task.TaskID,
		Name:      task.Name,
		Assign_To: assign_to,
	}
	return response, nil
}
