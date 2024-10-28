// package main

// import (
// 	"log"
// 	userhttp "user-service/internal/adapter/http"
// 	"user-service/internal/app"
// 	"user-service/internal/config"
// 	"user-service/internal/initializers"
// 	"user-service/internal/repository"
// )

// func main() {
// 	configuration, error := config.LoadConfig()
// 	if error != nil {
// 		log.Fatal(error)
// 	}
// 	log.Println(configuration)
// 	db := initializers.GetDB()
// 	repoObj := repository.NewUserRepository(db)
// 	app_obj := app.NewUserService(repoObj)
// 	handler_obj := userhttp.NewUserHandler(configuration, app_obj)
// 	handler_obj.Server()

// }
package main

import (
	"log"
	"sync"
	userhttp "user-service/internal/adapter/http" // Import the HTTP handler package
	"user-service/internal/app"
	"user-service/internal/config"
	"user-service/internal/grpc" // Import the gRPC server package
	"user-service/internal/initializers"
	"user-service/internal/repository"
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

	// Initialize services
	userService := app.NewUserService(userRepo)

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

	// Start gRPC server
	go func() {
		defer wg.Done()
		if err := grpc.RunGRPCServer(grpcPort, userService); err != nil {
			log.Fatalf("Failed to run gRPC server: %v", err)
		}
	}()

	// Wait for both servers to exit
	wg.Wait()
}
