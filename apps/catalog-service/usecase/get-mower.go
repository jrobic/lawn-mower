package usecase

import (
	"fmt"
	"jrobic/lawn-mower/catalog-service/domain"
)

func (lm *LMCatalogService) GetMower(id string) (*domain.Mower, error) {
	mower, _ := lm.repo.Find(id)

	if mower == nil {
		return nil, fmt.Errorf(domain.ErrMowerNotFound, id)
	}

	return mower, nil
}
