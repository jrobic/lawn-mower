package httpController

import (
	"encoding/json"
	"jrobic/lawn-mower/catalog-service/domain"
	"jrobic/lawn-mower/catalog-service/usecase"
	"net/http"
	"strings"
)

var (
	JSONContentType = "application/json"
)

type CreateMowerInputDTO struct {
	Name string `json:"name"`
}

type UpdateMowerInputDTO struct {
	Name string `json:"name,omitempty"`
}

type CatalogHTTPServer struct {
	http.Handler
	repo    domain.CatalogRepository
	service usecase.CatalogService
}

func NewCatalogHTTPServer(repo domain.CatalogRepository) (*CatalogHTTPServer, error) {
	s := new(CatalogHTTPServer)

	s.repo = repo
	s.service = usecase.NewCatalogService(repo)

	router := http.NewServeMux()
	router.HandleFunc("/mowers", http.HandlerFunc(s.CreateMower))
	router.HandleFunc("/mowers/", http.HandlerFunc(s.GetOrUpdateMower))
	router.HandleFunc("/", http.HandlerFunc(s.GetCatalog))

	s.Handler = router

	return s, nil
}

func (serv *CatalogHTTPServer) CreateMower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JSONContentType)

	mowerToCreate := &CreateMowerInputDTO{}

	err := json.NewDecoder(r.Body).Decode(&mowerToCreate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mower, err := serv.service.CreateMower(domain.CreateMowerDTO{Name: mowerToCreate.Name})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(mower)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (serv *CatalogHTTPServer) GetOrUpdateMower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JSONContentType)

	id := strings.TrimPrefix(r.URL.Path, "/mowers/")

	mower := &domain.Mower{}
	var err error

	switch r.Method {
	case http.MethodPatch:
		{

			mowerToUpdate := &UpdateMowerInputDTO{}
			err = json.NewDecoder(r.Body).Decode(&mowerToUpdate)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			mower, err = serv.service.UpdateMower(id, domain.UpdateMowerDTO{
				Name: mowerToUpdate.Name,
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	case http.MethodGet:
		{
			mower, err = serv.service.GetMower(id)

			if err != nil {
				http.NotFound(w, r)
				return
			}
		}
	}

	err = json.NewEncoder(w).Encode(mower)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (serv *CatalogHTTPServer) GetCatalog(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("content-type", JSONContentType)

	mowers, err := serv.service.GetAvailableMowers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(mowers)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
