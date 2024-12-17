package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
	userhttp "user-service/internal/adapter/http" // Import the HTTP handler package
	"user-service/internal/app"
	"user-service/internal/config"
	grpcServer "user-service/internal/grpc" // Import the gRPC server package
	"user-service/internal/initializers"
	"user-service/internal/repository"
	pb "user-service/proto/task" // Import your generated proto files here
)

func main() {
	// Load configuration
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Loaded configuration:", configuration)

	// Initialize the database
	db := initializers.GetDB()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Set up a connection to the server
	con, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials())) // Change the address as needed
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer con.Close()
	// Create a new client
	clientcon := pb.NewTaskServiceClient(con)

	// Initialize services
	userService := app.NewUserService(userRepo, clientcon)

	// Set up HTTP and gRPC handlers
	httpHandler := userhttp.NewUserHandler(configuration, userService)
	grpcPort := configuration.GRPCPort // Assuming GRPCPort is defined in your config

	// Run HTTP and gRPC servers concurrently
	var wg sync.WaitGroup
	wg.Add(2)

	// Start HTTP server
	go func() {
		defer wg.Done()
		httpHandler.Server()
	}()
	// //start grpc client
	// go func() {
	// 	defer wg.Done()
	// 	RunGrpcClient()
	// }()

	// Start gRPC server
	go func() {
		defer wg.Done()
		if err := grpcServer.RunGRPCServer(grpcPort, userService); err != nil {
			log.Fatalf("Failed to run gRPC server: %v", err)
		}
	}()

	// Wait for both servers to exit
	wg.Wait()
}

// func RunGrpcClient() {
// 	// Set up a connection to the server
// 	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials())) // Change the address as needed
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()

// 	// Create a new client
// 	client := pb.NewTaskServiceClient(conn)

// 	// // Call GetUser method
// 	// getUserRequest := &pb.GetTaskRequest{Email: "shyamsundar@gmail.com"}
// 	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	// defer cancel()

// 	// userResponse, err := client.GetTask(ctx, getUserRequest)
// 	// if err != nil {
// 	// 	log.Fatalf("could not get user: %v", err)
// 	// }
// 	// log.Printf("User: %v", userResponse)
// 	return
// }
