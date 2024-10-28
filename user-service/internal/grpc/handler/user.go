package handler

import (
	"context"
	"user-service/internal/app"
	pb "user-service/proto/user"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	hand app.UserServiceGetter
}

func NewUserHandler(hand app.UserServiceGetter) *UserHandler {
	return &UserHandler{
		hand: hand,
	}
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Get the user service instance

	user, err := h.hand.User().GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err

	}
	return &pb.GetUserResponse{
		Id:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}, nil
}

func (h *UserHandler) GetAllUser(ctx context.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	// Get the user service instance
	user, err := h.hand.User().GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	var users []*pb.User
	for _, user := range user {
		users = append(users, &pb.User{
			Id:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Address: user.Address,
		})
	}
	return &pb.GetAllUserResponse{
		Users: users,
	}, nil

}
