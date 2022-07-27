package lmTesting

import (
	"encoding/json"
	"fmt"
	"io"
	"jrobic/lawn-mower/catalog-service/domain"
	"net/http"
	"reflect"
	"testing"
)

type StubCatalogRepository struct {
	Mowers []*domain.Mower
}

func (r *StubCatalogRepository) Find(id string) (*domain.Mower, error) {
	for i, mower := range r.Mowers {
		if mower.ID == id {
			return r.Mowers[i], nil
		}
	}
	return nil, nil
}

func (r *StubCatalogRepository) Add(input domain.CreateMowerDTO) (*domain.Mower, error) {
	id := fmt.Sprint(len(r.Mowers) + 1)

	mower := &domain.Mower{
		ID:   id,
		Name: input.Name,
	}

	r.Mowers = append(r.Mowers, mower)

	return mower, nil
}

func (r *StubCatalogRepository) Patch(id string, input domain.UpdateMowerDTO) (mower *domain.Mower, err error) {
	for i, mower := range r.Mowers {
		if mower.ID == id {
			r.Mowers[i].Name = input.Name
			mower = r.Mowers[i]
			return mower, err
		}
	}

	return mower, err
}

func (r *StubCatalogRepository) FindAvailableMowers() ([]*domain.Mower, error) {
	return r.Mowers, nil
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertMowerEquals(t testing.TB, got, want domain.Mower) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertCatalogEquals(t testing.TB, got, want []*domain.Mower) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func AssertContentType(t testing.TB, response *http.Response, want string) {
	t.Helper()
	if response.Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Header)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func GetMowerFromResponse(t testing.TB, body io.Reader) (mower domain.Mower) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&mower)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into Mower, '%v'", body, err)
	}

	return
}

func GetCatalogFromResponse(t testing.TB, body io.Reader) (mowers []*domain.Mower) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&mowers)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into Catalog, '%v'", body, err)
	}

	return
}
