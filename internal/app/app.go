package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"test_task_golang/configs"
	"test_task_golang/internal/app/handler"
	"test_task_golang/internal/app/repository"
	"test_task_golang/internal/app/service"
	"test_task_golang/pkg/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() {
	if err := configs.Init(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	r := repository.New()
	s := service.New(r)
	h := handler.New(s)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("http.port"), h.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	r.Wg.Wait()

	logrus.Print("Server Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
