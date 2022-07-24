package usecase

import (
	lmTesting "jrobic/lawn-mower/catalog-service"
	domain "jrobic/lawn-mower/catalog-service/domain"
	"reflect"
	"strings"

	"testing"
)

func TestGetMower(t *testing.T) {
	t.Run("catalog: return a mower", func(t *testing.T) {
		want := &domain.Mower{ID: "1", Name: "M-350"}

		wantedCatalog := []*domain.Mower{
			want,
		}
		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}

		service := NewCatalogService(repo)

		got, err := service.GetMower("1")

		lmTesting.AssertNoError(t, err)
		lmTesting.AssertMowerEquals(t, *got, *want)
	})

	t.Run("catalog: return error when mower not found", func(t *testing.T) {
		want := &domain.Mower{ID: "1", Name: "M-350"}
		wantedCatalog := []*domain.Mower{
			want,
		}

		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}

		IDNotFound := "2"

		service := NewCatalogService(repo)

		_, err := service.GetMower(IDNotFound)

		got := strings.Replace(domain.ErrMowerNotFound, "%v", IDNotFound, 1)

		lmTesting.AssertError(t, err, got)
	})
}

func TestCreateMower(t *testing.T) {
	t.Run("catalog: create new mower", func(t *testing.T) {
		wantedCatalog := []*domain.Mower{}
		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}
		service := NewCatalogService(repo)

		newMower := domain.CreateMowerDTO{Name: "M-150"}

		insertedMower, err := service.CreateMower(newMower)

		lmTesting.AssertNoError(t, err)

		got, _ := service.GetMower(insertedMower.ID)

		if got == nil {
			t.Errorf("could not find new created mower")
		}
	})
}

func TestUpdateMower(t *testing.T) {
	t.Run("catalog: update mower", func(t *testing.T) {
		wantedMower := domain.Mower{ID: "1", Name: "M-90"}

		wantedUpdatedMower := wantedMower
		wantedUpdatedMower.Name = "M-150"

		wantedCatalog := []*domain.Mower{
			&wantedMower,
		}

		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}
		service := NewCatalogService(repo)

		updateMower := domain.UpdateMowerDTO{Name: wantedUpdatedMower.Name}

		_, err := service.UpdateMower(wantedCatalog[0].ID, updateMower)

		lmTesting.AssertNoError(t, err)

		got, _ := service.GetMower(wantedCatalog[0].ID)

		if got == nil {
			t.Errorf("could not find updated mower")
			return
		}

		lmTesting.AssertMowerEquals(t, wantedUpdatedMower, *got)
	})

	t.Run("catalog: update mower with empty 'Name'", func(t *testing.T) {
		wantedMower := domain.Mower{ID: "1", Name: "M-90"}

		wantedCatalog := []*domain.Mower{
			&wantedMower,
		}

		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}
		service := NewCatalogService(repo)

		updateMower := domain.UpdateMowerDTO{}

		_, err := service.UpdateMower(wantedCatalog[0].ID, updateMower)

		lmTesting.AssertNoError(t, err)

		got, _ := service.GetMower(wantedCatalog[0].ID)

		if got == nil {
			t.Errorf("could not find updated mower")
			return
		}

		lmTesting.AssertMowerEquals(t, wantedMower, *got)
	})
}

func TestGetAvailableMowers(t *testing.T) {
	t.Run("catalog: find all mowers", func(t *testing.T) {
		wantedCatalog := []*domain.Mower{
			{ID: "1", Name: "M-90"},
			{ID: "2", Name: "M-150"},
			{ID: "3", Name: "M-480"},
		}
		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}

		wantedMowers := []*domain.Mower{
			{ID: "1", Name: "M-90"},
			{ID: "2", Name: "M-150"},
			{ID: "3", Name: "M-480"},
		}

		service := NewCatalogService(repo)

		got, err := service.GetAvailableMowers()

		lmTesting.AssertNoError(t, err)

		if len(got) != len(wantedMowers) {
			t.Errorf("got %v want %v", len(got), len(wantedMowers))
		}

		if !reflect.DeepEqual(got, wantedMowers) {
			t.Errorf("got %v want %v", got, wantedMowers)
		}
	})
}
