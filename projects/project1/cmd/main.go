package main

import (
	"log"

	"github.com/kirigaikabuto/hero_preparation/projects/project1"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/handler"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/repository"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(project1.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running of server %s", err.Error())
	}
}
