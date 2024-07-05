package main

import (
	"test_task_golang/configs"
	"test_task_golang/internal/app"

	_ "test_task_golang/docs"

	"github.com/sirupsen/logrus"
)

// @title 	Tasks API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath /
func main() {
	// Configuration
	err := configs.Init()
	if err != nil {
		logrus.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run()
}
