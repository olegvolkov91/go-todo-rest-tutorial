package main

import (
	"log"

	"github.com/olegvolkov91/go-todo-rest-tutorial"
	"github.com/olegvolkov91/go-todo-rest-tutorial/package/handler"
	"github.com/olegvolkov91/go-todo-rest-tutorial/package/repository"
	"github.com/olegvolkov91/go-todo-rest-tutorial/package/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
