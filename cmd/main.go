package main

import (
	"gitgub.com/diploma-mppr/backend_mppr/conf"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/task"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func main() {
	pgxManager, err := db.NewPostgresqlX()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	taskRepo := task.NewRepositoryTask(pgxManager)
	taskUcase := task.NewUcaseTask(taskRepo)
	taskHandler := task.NewHandlerTask(taskUcase)

	router := echo.New()

	serverRouting := conf.ServerHandlers{
		TaskHandler: taskHandler,
	}

	serverRouting.ConfigureRouting(router)

	httpServ := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	if err := router.StartServer(&httpServ); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
