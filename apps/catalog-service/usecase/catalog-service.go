package usecase

import (
	"jrobic/lawn-mower/catalog-service/domain"
)

type CatalogService interface {
	CreateMower(input domain.CreateMowerDTO) (*domain.Mower, error)
	UpdateMower(id string, input domain.UpdateMowerDTO) (*domain.Mower, error)
	GetMower(id string) (*domain.Mower, error)
	GetAvailableMowers() ([]*domain.Mower, error)
}

type LMCatalogService struct {
	repo domain.CatalogRepository
}

func NewCatalogService(repo domain.CatalogRepository) *LMCatalogService {
	return &LMCatalogService{
		repo: repo,
	}
}
