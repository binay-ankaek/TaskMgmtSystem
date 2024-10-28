package ports

import (
	"context"
	"user-service/internal/domain/model"
)

type UserRepository interface {
	//create user
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	// Get all users
	GetAllUsers(ctx context.Context) ([]model.User, error)
	// Get user by id
	GetUserById(ctx context.Context, id string) (*model.User, error)
	//Get user by username
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	//Update user
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
}

type UserService interface {
	//create user
	CreateUser(ctx context.Context, user *model.CreateRequest) (*model.CreateResponse, error)
	// Get all users
	GetAllUsers(ctx context.Context) ([]model.CreateResponse, error)
	//Get user
	GetUserById(ctx context.Context, id string) (*model.CreateResponse, error)
	//Get User by name
	GetUserByEmail(ctx context.Context, email string) (*model.PasswordResponse, error)
	//update user
	UpdateUser(ctx context.Context, user *model.CreateResponse) (*model.UpdateUser, error)
}
