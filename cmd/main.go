package main

import (
	"test_task_golang/pkg/handler"
	"test_task_golang/pkg/repository"
	"test_task_golang/pkg/service"
)

func main() {
	r := repository.NewRepository()
	s := service.NewService(r)
	h := handler.NewHandler(s)
	router := h.InitRoutes()
	router.Run(":8080")
}
