package grpc

import (
	"google.golang.org/grpc/reflection" // Import the reflection package
	"log"
	"net"
	"user-service/internal/app"
	"user-service/internal/grpc/handler"
	pb "user-service/proto/user"

	"google.golang.org/grpc"
)

// RunGRPCServer starts the gRPC server
func RunGRPCServer(port string, userService app.UserServiceGetter) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()

	// Create a new ProfileHandler with the injected ProfileService
	profileHandler := handler.NewUserHandler(userService)

	// Register Profile service
	pb.RegisterUserServiceServer(grpcServer, profileHandler)

	// Register reflection service on gRPC server
	reflection.Register(grpcServer)

	log.Printf("gRPC server listening on %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return err
	}

	return nil
}
