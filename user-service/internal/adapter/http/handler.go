package http

import (
	"fmt"
	"log"
	"user-service/internal/app"
	"user-service/internal/config"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	// Service is the business logic layer
	config *config.Config
	hand   app.UserServiceGetter
	router *gin.Engine
}

func NewUserHandler(config *config.Config, hand app.UserServiceGetter) *UserHandler {
	return &UserHandler{
		config: config,
		hand:   hand,
		router: gin.New(),
	}
}

func (h *UserHandler) Server() {
	err := h.SetupRoute()
	if err != nil {
		log.Fatal(err)
	}
	listenaddr := fmt.Sprintf("%s:%s", "127.0.0.1", h.config.Port)
	h.router.Run(listenaddr)
}
