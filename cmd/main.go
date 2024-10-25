package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	httpctrl "github.com/himmel520/practice2024/internal/controller/websocket"
	"github.com/himmel520/practice2024/internal/server"
	"github.com/himmel520/practice2024/internal/usecase"
)

func main() {
	log := server.SetupLogger()

	addr := os.Getenv("SERVER_ADDRESS")

	usecase := usecase.New()
	handler := httpctrl.New(usecase, log)

	app := server.New(handler.InitRoutes(), addr)
	go func() {
		log.Infof("the server is starting on %v", addr)

		if err := app.Run(); err != nil {
			log.Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	if err := app.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err)
	}

	log.Info("the server is shut down")
}
