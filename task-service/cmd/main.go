package main

import (
	"context"
	"fmt"
	"log"
	"task-service/internal/initializers"
	pb "task-service/proto/user" // Import your generated proto files here
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
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
	//db init
	db := initializers.GetDB()
	fmt.Println(db)

}
