package usecase

import (
	"jrobic/lawn-mower/catalog-service/domain"
)

func (lm *LMCatalogService) UpdateMower(id string, input domain.UpdateMowerDTO) (*domain.Mower, error) {
	mower, err := lm.repo.Patch(id, input)

	if err != nil {
		return nil, err
	}

	return mower, nil
}
