package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"user-service/internal/domain/model"
	"user-service/internal/ports"
	"user-service/internal/repository"
	"user-service/internal/utils"
	pbTask "user-service/proto/task"
)

type UserService struct {
	repo         repository.UserGetter
	pbTaskClient pbTask.TaskServiceClient
}

type UserServiceGetter interface {
	User() ports.UserService
}

func (r *UserService) User() ports.UserService {
	return r
}

func NewUserService(repo repository.UserGetter, pbTaskClient pbTask.TaskServiceClient) *UserService {
	return &UserService{
		repo:         repo,
		pbTaskClient: pbTaskClient,
	}
}

// CreateUser handles the creation of a new user.
func (r *UserService) CreateUser(ctx context.Context, user *model.CreateRequest) (*model.CreateResponse, error) {
	// Creating a new user model from the request data
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err

	}

	newUser := &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Password: password,
	}

	// Creating the user in the repository
	createdUser, err := r.repo.User().CreateUser(ctx, newUser)
	if err != nil {
		// Return nil for the response and the error if creation fails
		return nil, err
	}

	// Creating a response based on the created user data
	response := &model.CreateResponse{
		ID:      createdUser.ID,
		Name:    createdUser.Name,
		Email:   createdUser.Email,
		Address: createdUser.Address,
	}

	// Return the response and nil for the error if successful
	return response, nil
}

func (r *UserService) GetAllUsers(ctx context.Context) ([]model.CreateResponse, error) {
	// Get all users from the repository
	users, err := r.repo.User().GetAllUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//loop in users
	var response []model.CreateResponse
	for _, user := range users {
		//create a new response
		response = append(response, model.CreateResponse{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Address: user.Address,
		})

	}
	return response, nil

}

func (r *UserService) GetUserByEmail(ctx context.Context, email string) (*model.PasswordResponse, error) {
	// Get user by email from the repository
	user, err := r.repo.User().GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// If user is not found, return a nil user and error
	if user == nil {
		return nil, fmt.Errorf("%s not found", email)
	}

	// Return the full response with the password included
	response := &model.PasswordResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Password: user.Password, // Ensure this is set
	}
	return response, nil
}

// get user by id
func (r *UserService) GetUserById(ctx context.Context, id string) (*model.CreateResponse, error) {
	//get user by id

	user, err := r.repo.User().GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	response := &model.CreateResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}
	return response, nil

}

func (r *UserService) UpdateUser(ctx context.Context, user *model.CreateResponse) (*model.UpdateUser, error) {
	if user.ID == "" {
		return nil, errors.New("user id is required")
	}
	// Update user in the repository
	updatednewdata := &model.User{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}
	updateduser, err := r.repo.User().UpdateUser(ctx, updatednewdata)
	if err != nil {
		return nil, err
	}
	updatedUserResponse := &model.UpdateUser{
		Name:    updateduser.Name,
		Email:   updateduser.Email,
		Address: updateduser.Address,
	}
	return updatedUserResponse, nil

}

func (r *UserService) GetTask(ctx context.Context, email string) (*model.TaskResponses, error) {
	//get email user
	if email == "" {
		return nil, errors.New("email is required")
	}
	//get task from grpc
	task, err := r.pbTaskClient.GetTask(ctx, &pbTask.GetTaskRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	var TaskResponse []model.TaskResponse
	for _, t := range task.Tasks {
		TaskResponse = append(TaskResponse, model.TaskResponse{
			TaskID:   t.Id,
			TaskName: t.Name,
		})
	}
	return &model.TaskResponses{
		TaskResponses: TaskResponse,
	}, nil

}
