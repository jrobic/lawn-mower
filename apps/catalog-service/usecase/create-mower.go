package usecase

import (
	"jrobic/lawn-mower/catalog-service/domain"
)

func (lm *LMCatalogService) CreateMower(input domain.CreateMowerDTO) (*domain.Mower, error) {
	mower, err := lm.repo.Add(input)

	if err != nil {
		return nil, err
	}

	return mower, nil
}
