package main

import (
	"context"
	"log"
	"sync"
	"task-service/internal/adapter/http"
	"task-service/internal/app"
	"task-service/internal/initializers"
	"task-service/internal/repository"
	pb "task-service/proto/user" // Import your generated proto files here
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// init database
	db := initializers.GetDB()
	//get repo
	repo := repository.NewTaskRepository(db)
	// Set up a connection to the server
	con, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // Change the address as needed
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer con.Close()

	// Create a new client
	clientcon := pb.NewUserServiceClient(con)
	//get service
	service := app.NewTaskService(repo, clientcon)
	//get handler
	handler := http.NewTaskHandler(service)
	var wg sync.WaitGroup
	// Start the server
	wg.Add(2)
	go func() {
		defer wg.Done()
		log.Println("starting the http server....")
		//run the server
		if err := handler.Server(); err != nil {
			log.Println(err)
		}

	}()
	// Start the gRPC client in another goroutine
	go func() {
		defer wg.Done() // Mark this goroutine as done when it completes
		RunGrpcClient()
	}()

	// Wait for both tasks to complete
	wg.Wait()
	log.Println("All tasks completed.")

}
func RunGrpcClient() {
	// Set up a connection to the server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // Change the address as needed
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewUserServiceClient(conn)

	// Call GetUser method
	getUserRequest := &pb.GetUserRequest{Email: "shyamsundar@gmail.com"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userResponse, err := client.GetUser(ctx, getUserRequest)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %v", userResponse)

	// Call GetAllUser method
	allUserRequest := &pb.GetAllUserRequest{}
	allUsersResponse, err := client.GetAllUser(ctx, allUserRequest)
	if err != nil {
		log.Fatalf("could not get all users: %v", err)
	}
	log.Printf("All Users: %v", allUsersResponse.Users)
}
