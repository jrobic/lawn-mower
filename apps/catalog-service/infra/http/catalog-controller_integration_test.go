package restcontroller

import (
	"fmt"
	lmTesting "jrobic/lawn-mower/catalog-service"
	"jrobic/lawn-mower/catalog-service/domain"
	"jrobic/lawn-mower/catalog-service/infra/repository"
	"net/http"
	"testing"
)

func TestCreateMowersAndRetrievingThemes(t *testing.T) {
	repo := repository.NewInMemoryRepo([]*domain.Mower{})
	server, _ := NewCatalogHTTPServer(repo)

	wantedMowers := []*domain.Mower{
		{ID: "1", Name: "M-90"},
		{ID: "2", Name: "M-150"},
		{ID: "3", Name: "M-480"},
	}

	for _, wantedMower := range wantedMowers {
		server.App.Test(NewCreateMowerRequest(&CreateMowerInputDTO{Name: wantedMower.Name}), -1)
	}

	for _, wantedMower := range wantedMowers {
		testCaseName := "get " + wantedMower.Name + " mower"

		t.Run(testCaseName, func(t *testing.T) {
			response, _ := server.App.Test(NewGetMowerRequest(wantedMower.ID), -1)

			got := lmTesting.GetMowerFromResponse(t, response.Body)

			lmTesting.AssertStatus(t, response.StatusCode, http.StatusOK)
			lmTesting.AssertMowerEquals(t, got, *wantedMower)
		})
	}

	t.Run("get list mowers", func(t *testing.T) {

		response, _ := server.App.Test(NewGetCatalogRequest(), -1)

		got := lmTesting.GetCatalogFromResponse(t, response.Body)

		lmTesting.AssertCatalogEquals(t, got, wantedMowers)
		lmTesting.AssertStatus(t, response.StatusCode, http.StatusOK)
		lmTesting.AssertContentType(t, response, JSONContentType)
	})
}

func TestUpdateMowersAndRetrievingThemes(t *testing.T) {
	wantedMowers := []*domain.Mower{
		{ID: "1", Name: "M-90"},
		{ID: "2", Name: "M-150"},
		{ID: "3", Name: "M-480"},
	}

	wantedUpdatedMowers := []*domain.Mower{
		{ID: "1", Name: "M-90"},
		{ID: "2", Name: "M-150"},
		{ID: "3", Name: "M-390"},
	}

	repo := repository.NewInMemoryRepo(wantedMowers)
	server, _ := NewCatalogHTTPServer(repo)

	for _, wantedMower := range wantedUpdatedMowers {
		server.App.Test(NewUpdateMowerRequest(wantedMower.ID, &UpdateMowerInputDTO{Name: wantedMower.Name}), -1)
	}

	for _, wantedMower := range wantedUpdatedMowers {
		testCaseName := "patch " + wantedMower.Name + " mower"

		t.Run(testCaseName, func(t *testing.T) {
			response, _ := server.App.Test(NewGetMowerRequest(wantedMower.ID), -1)

			got := lmTesting.GetMowerFromResponse(t, response.Body)

			lmTesting.AssertStatus(t, response.StatusCode, http.StatusOK)
			lmTesting.AssertMowerEquals(t, got, *wantedMower)
		})
	}
}

func BenchmarkCreateMower(b *testing.B) {
	repo := repository.NewInMemoryRepo([]*domain.Mower{})
	server, _ := NewCatalogHTTPServer(repo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := fmt.Sprintf("M-90 %v", i)
		server.App.Test(NewCreateMowerRequest(&CreateMowerInputDTO{Name: name}), -1)
	}
}

func BenchmarkGetMower(b *testing.B) {
	mowers := []*domain.Mower{}

	for i := 0; i < 200; i++ {
		ID := fmt.Sprintf("%d", i)
		name := fmt.Sprintf("M-%d-50", i)

		mowers = append(mowers, &domain.Mower{ID: ID, Name: name})
	}

	repo := repository.NewInMemoryRepo(mowers)
	server, _ := NewCatalogHTTPServer(repo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		server.App.Test(NewGetMowerRequest("120"), -1)
	}
}

func BenchmarkGetCatalog(b *testing.B) {
	mowers := []*domain.Mower{}

	for i := 0; i < 200; i++ {
		ID := fmt.Sprintf("%d", i)
		name := fmt.Sprintf("M-%d-50", i)

		mowers = append(mowers, &domain.Mower{ID: ID, Name: name})
	}

	repo := repository.NewInMemoryRepo(mowers)
	server, _ := NewCatalogHTTPServer(repo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		server.App.Test(NewGetCatalogRequest(), -1)
	}
}
