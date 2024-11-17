package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"task-service/internal/app"
)

type TaskHandler struct {
	svc    app.TaskServiceGetter
	router *gin.Engine
}

func NewTaskHandler(svc app.TaskServiceGetter) *TaskHandler {

	return &TaskHandler{
		svc:    svc,
		router: gin.New(),
	}
}

func (h *TaskHandler) Server() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Println("No .env file is exist there in our code!")

	}
	err = h.SetupRoute()
	if err != nil {
		fmt.Println(err)
	}
	port := os.Getenv("Port")
	if port == "" {
		port = "8081"
	}
	listenaddr := fmt.Sprintf("%s:%s", "127.0.0", port)
	h.router.Run(listenaddr)
}
