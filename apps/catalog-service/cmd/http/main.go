package main

import (
	"jrobic/lawn-mower/catalog-service/domain"
	httpController "jrobic/lawn-mower/catalog-service/infra/http"
	"jrobic/lawn-mower/catalog-service/infra/repository"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewInMemoryRepo([]*domain.Mower{
		{Id: "1", Name: "M-90"},
		{Id: "2", Name: "M-150"},
		{Id: "3", Name: "M-480"},
	})

	server, err := httpController.NewCatalogHTTPServer(repo)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	log.Println("Listen on port 5001")

	if err := http.ListenAndServe(":5001", server); err != nil {
		log.Fatalf("could not listen on port 5001 %v", err)
	}

}
