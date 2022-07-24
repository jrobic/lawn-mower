package repository

import (
	"fmt"
	"jrobic/lawn-mower/catalog-service/domain"
	"sync"
)

type InMemoryRepo struct {
	Mowers []*domain.Mower
	lock   sync.RWMutex
}

func NewInMemoryRepo(initialMowers []*domain.Mower) *InMemoryRepo {
	mowers := []*domain.Mower{}

	mowers = append(mowers, initialMowers...)

	return &InMemoryRepo{Mowers: mowers}
}

func (r *InMemoryRepo) Add(input domain.CreateMowerDTO) (*domain.Mower, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	id := fmt.Sprint(len(r.Mowers) + 1)

	mower := &domain.Mower{
		ID:   id,
		Name: input.Name,
	}

	r.Mowers = append(r.Mowers, mower)

	return mower, nil
}

func (r *InMemoryRepo) Patch(id string, input domain.UpdateMowerDTO) (mower *domain.Mower, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i, mower := range r.Mowers {
		if mower.ID == id {
			r.Mowers[i].Name = input.Name
			mower = r.Mowers[i]
			return mower, err
		}
	}

	return mower, err
}

func (r *InMemoryRepo) Find(id string) (mower *domain.Mower, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i, mower := range r.Mowers {
		if mower.ID == id {
			return r.Mowers[i], nil
		}
	}
	return nil, nil
}

func (r *InMemoryRepo) FindAvailableMowers() (mowers []*domain.Mower, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	return r.Mowers, nil
}
