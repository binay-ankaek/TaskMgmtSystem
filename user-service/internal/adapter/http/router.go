package http

import (
	"user-service/internal/app"
)

func (h *UserHandler) SetupRoute() error {
	api := h.router.Group("api/v1")
	{
		//create user
		api.POST("/signup", h.SignUp)
		//get all user
		api.GET("/user", h.GetAllUsers)
		//login
		api.POST("/login", h.Login)
	}

	// Protected API group (requires JWT authentication)
	protected := h.router.Group("/api/v1/protected")
	protected.Use(app.AuthMiddleware()) // Apply middleware from `app`
	{
		protected.GET("/profile", h.GetProfile) // Get authenticated user's profile
		protected.PUT("/update-profile", h.UpdateProfile)
	}
	return nil
}
