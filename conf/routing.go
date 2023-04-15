package conf

import (
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/task"
	"github.com/labstack/echo/v4"
)

type ServerHandlers struct {
	TaskHandler *task.HandlerTask
}

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo) {
	//router.GET("/api/get_task", sh.TaskHandler.GetTask)
	router.GET("/api/get_task_by_id", sh.TaskHandler.SetData)
}
