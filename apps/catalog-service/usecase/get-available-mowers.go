package usecase

import (
	"jrobic/lawn-mower/catalog-service/domain"
)

func (lm *LMCatalogService) GetAvailableMowers() ([]*domain.Mower, error) {
	return lm.repo.FindAvailableMowers()
}
