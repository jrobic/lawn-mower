package restcontroller

import (
	"bytes"
	"encoding/json"

	"net/http"
	"testing"

	lmTesting "jrobic/lawn-mower/catalog-service"
	"jrobic/lawn-mower/catalog-service/domain"
)

func TestCreateMowerCtrl(t *testing.T) {

	t.Run("CreateMowerCtrl return accepted on POST", func(t *testing.T) {
		wantedCatalog := []*domain.Mower{
			{ID: "1", Name: "M-90"},
			{ID: "2", Name: "M-150"},
			{ID: "3", Name: "M-480"},
		}
		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}
		server, _ := NewCatalogHTTPServer(repo)

		mower := &CreateMowerInputDTO{Name: "M-600"}
		wantedMower := domain.Mower{Name: "M-600", ID: "4"}

		request := NewCreateMowerRequest(mower)

		response, _ := server.App.Test(request, -1)

		got := lmTesting.GetMowerFromResponse(t, response.Body)

		lmTesting.AssertStatus(t, response.StatusCode, http.StatusAccepted)
		lmTesting.AssertContentType(t, response, JSONContentType)

		lmTesting.AssertMowerEquals(t, got, wantedMower)
	})
}

func TestUpdateMowerCtrl(t *testing.T) {

	t.Run("UpdateMowerCtrl return accepted on PATCH", func(t *testing.T) {
		wantedMower := domain.Mower{ID: "3", Name: "M-480"}

		wantedUpdatedMower := wantedMower
		wantedUpdatedMower.Name = "M-380"

		wantedCatalog := []*domain.Mower{
			{ID: "1", Name: "M-90"},
			{ID: "2", Name: "M-150"},
			&wantedMower,
		}

		repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}
		server, _ := NewCatalogHTTPServer(repo)

		updateMower := domain.UpdateMowerDTO{Name: wantedUpdatedMower.Name}

		request := NewUpdateMowerRequest(wantedMower.ID, updateMower)

		response, _ := server.App.Test(request, -1)

		got := lmTesting.GetMowerFromResponse(t, response.Body)

		lmTesting.AssertStatus(t, response.StatusCode, http.StatusOK)
		lmTesting.AssertContentType(t, response, JSONContentType)

		lmTesting.AssertMowerEquals(t, got, wantedUpdatedMower)
	})
}

func TestGetMowerCtrl(t *testing.T) {
	wantedCatalog := []*domain.Mower{
		{ID: "1", Name: "M-90"},
		{ID: "2", Name: "M-150"},
		{ID: "3", Name: "M-480"},
	}

	repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}
	server, _ := NewCatalogHTTPServer(repo)

	t.Run("GetMowerCtrl return M-350 mower", func(t *testing.T) {
		request := NewGetMowerRequest("1")

		response, _ := server.App.Test(request, -1)

		wantedMower := domain.Mower{ID: "1", Name: "M-90"}
		got := lmTesting.GetMowerFromResponse(t, response.Body)

		lmTesting.AssertStatus(t, response.StatusCode, http.StatusOK)
		lmTesting.AssertContentType(t, response, JSONContentType)

		lmTesting.AssertMowerEquals(t, got, wantedMower)
	})

	t.Run("GetMowerCtrl return M-150 mower", func(t *testing.T) {
		request := NewGetMowerRequest("2")

		response, _ := server.App.Test(request, -1)

		got := lmTesting.GetMowerFromResponse(t, response.Body)
		wantedMower := domain.Mower{ID: "2", Name: "M-150"}

		lmTesting.AssertStatus(t, response.StatusCode, http.StatusOK)
		lmTesting.AssertContentType(t, response, JSONContentType)

		lmTesting.AssertMowerEquals(t, got, wantedMower)
	})

	t.Run("GetMowerCtrl return 404 on missing mower", func(t *testing.T) {
		request := NewGetMowerRequest("6")

		response, _ := server.App.Test(request, -1)

		lmTesting.AssertStatus(t, response.StatusCode, http.StatusNotFound)
	})
}

func TestGetCatalogCtrl(t *testing.T) {
	wantedCatalog := []*domain.Mower{
		{ID: "1", Name: "M-90"},
		{ID: "2", Name: "M-150"},
		{ID: "3", Name: "M-480"},
	}

	repo := &lmTesting.StubCatalogRepository{Mowers: wantedCatalog}
	server, _ := NewCatalogHTTPServer(repo)

	t.Run("GetCatalogCtrl return list of mowers", func(t *testing.T) {
		request := NewGetCatalogRequest()

		response, _ := server.App.Test(request, -1)

		got := lmTesting.GetCatalogFromResponse(t, response.Body)

		lmTesting.AssertCatalogEquals(t, got, wantedCatalog)
		lmTesting.AssertStatus(t, response.StatusCode, http.StatusOK)
		lmTesting.AssertContentType(t, response, JSONContentType)
	})
}

func NewCreateMowerRequest(body interface{}) *http.Request {
	jsonBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, "/mowers", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	return req
}

func NewUpdateMowerRequest(ID string, body interface{}) *http.Request {
	jsonBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPatch, "/mowers/"+ID, bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	return req
}

func NewGetMowerRequest(ID string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/mowers/"+ID, nil)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func NewGetCatalogRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	return req
}
