package http

import ()

func (h *TaskHandler) SetupRoute() error {
	api := h.router.Group("api/v1/task")
	{
		// Create a new task
		api.POST("", h.CreateTask)

	}
	return nil

}
