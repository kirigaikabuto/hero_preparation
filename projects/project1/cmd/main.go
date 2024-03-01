package main

import (
	"log"

	"github.com/kirigaikabuto/hero_preparation/projects/project1"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/handler"
)

func main() {
	srv := new(project1.Server)
	handlers := new(handler.Handler)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running of server %s", err.Error())
	}
}
