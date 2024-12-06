package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection" // Import the reflection package
	"log"
	"net"
	"task-service/internal/app"
	"task-service/internal/grpc/handler"
	pb "task-service/proto/task"
)

// RunGRPCServer starts the gRPC server
func RunGRPCServer(port string, taskService app.TaskServiceGetter) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}
	// Create a new gRPC server
	srv := grpc.NewServer()
	// Register the task service handler
	taskhandler := handler.NewTaskHandler(taskService)
	// Register the task service handler
	pb.RegisterTaskServiceServer(srv, taskhandler)

	// Register the reflection service
	reflection.Register(srv)

	// Start the gRPC server
	log.Printf("gRPC server is listening on port %s", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return err
	}
	return nil

}
