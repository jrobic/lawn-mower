package main

import (
	"jrobic/lawn-mower/catalog-service/domain"
	restcontroller "jrobic/lawn-mower/catalog-service/infra/http"
	"jrobic/lawn-mower/catalog-service/infra/repository"
	"log"
)

func main() {
	repo := repository.NewInMemoryRepo([]*domain.Mower{
		{ID: "1", Name: "M-90"},
		{ID: "2", Name: "M-150"},
		{ID: "3", Name: "M-480"},
	})

	server, err := restcontroller.NewCatalogHTTPServer(repo)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	log.Println("Listen on port 5001")

	if err := server.App.Listen(":5001"); err != nil {
		log.Fatalf("could not listen on port 5001 %v", err)
	}

}
