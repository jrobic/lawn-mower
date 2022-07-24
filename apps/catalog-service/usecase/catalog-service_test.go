package usecase

import (
	lmTesting "jrobic/lawn-mower/catalog-service"
	domain "jrobic/lawn-mower/catalog-service/domain"
	"reflect"
	"strings"

	"testing"
)

func TestFindMower(t *testing.T) {
	t.Run("catalog: return a mower", func(t *testing.T) {
		want := &domain.Mower{Id: "1", Name: "M-350"}

		wantedCatalog := []*domain.Mower{
			want,
		}
		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}

		service := NewCatalogService(repo)

		got, err := service.Find("1")

		lmTesting.AssertNoError(t, err)
		lmTesting.AssertMowerEquals(t, *got, *want)
	})

	t.Run("catalog: return error when mower not found", func(t *testing.T) {
		want := &domain.Mower{Id: "1", Name: "M-350"}
		wantedCatalog := []*domain.Mower{
			want,
		}

		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}

		idNotFound := "2"

		service := NewCatalogService(repo)

		_, err := service.Find(idNotFound)

		got := strings.Replace(domain.ErrMowerNotFound, "%v", idNotFound, 1)

		lmTesting.AssertError(t, err, got)
	})
}

func TestAddMower(t *testing.T) {
	t.Run("catalog: add mower", func(t *testing.T) {
		wantedCatalog := []*domain.Mower{}
		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}
		service := NewCatalogService(repo)

		newMower := domain.AddMowerDTO{Name: "M-150"}

		insertedMower, err := service.Add(newMower)

		lmTesting.AssertNoError(t, err)

		got, _ := service.Find(insertedMower.Id)

		if got == nil {
			t.Errorf("could not find new added mower")
		}
	})
}

func TestFindAllMowers(t *testing.T) {
	t.Run("catalog: find all mowers", func(t *testing.T) {
		wantedCatalog := []*domain.Mower{
			{Id: "1", Name: "M-90"},
			{Id: "2", Name: "M-150"},
			{Id: "3", Name: "M-480"},
		}
		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}

		wantedMowers := []*domain.Mower{
			{Id: "1", Name: "M-90"},
			{Id: "2", Name: "M-150"},
			{Id: "3", Name: "M-480"},
		}

		service := NewCatalogService(repo)

		got, err := service.FindAll()

		lmTesting.AssertNoError(t, err)

		if len(got) != len(wantedMowers) {
			t.Errorf("got %v want %v", len(got), len(wantedMowers))
		}

		if !reflect.DeepEqual(got, wantedMowers) {
			t.Errorf("got %v want %v", got, wantedMowers)
		}
	})
}
