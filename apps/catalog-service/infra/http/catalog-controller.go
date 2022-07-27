package restcontroller

import (
	"jrobic/lawn-mower/catalog-service/domain"
	"jrobic/lawn-mower/catalog-service/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var (
	JSONContentType = "application/json"
)

type CreateMowerInputDTO struct {
	Name string `json:"name"`
}

type UpdateMowerInputDTO struct {
	Name string `json:"name,omitempty"`
}

type CatalogHTTPServer struct {
	App     *fiber.App
	repo    domain.CatalogRepository
	service usecase.CatalogService
}

func NewCatalogHTTPServer(repo domain.CatalogRepository) (*CatalogHTTPServer, error) {
	s := new(CatalogHTTPServer)

	s.repo = repo
	s.service = usecase.NewCatalogService(repo)

	app := fiber.New()

	app.Use(requestid.New())
	app.Use(compress.New())
	app.Use(etag.New())

	app.Get("/", s.GetCatalog)
	app.Post("/mowers", s.CreateMower)
	app.Get("/mowers/:id", s.GetMower)
	app.Patch("/mowers/:id", s.UpdateMower)

	s.App = app

	return s, nil
}

func (serv *CatalogHTTPServer) CreateMower(c *fiber.Ctx) error {
	c.Append("content-type", JSONContentType)

	mowerToCreate := new(CreateMowerInputDTO)

	err := c.BodyParser(mowerToCreate)

	if err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	mower, err := serv.service.CreateMower(domain.CreateMowerDTO{Name: mowerToCreate.Name})

	if err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	c.Status(http.StatusAccepted)

	return c.JSON(mower)
}

func (serv *CatalogHTTPServer) GetMower(c *fiber.Ctx) error {
	c.Append("content-type", JSONContentType)

	id := c.Params("id")

	mower, err := serv.service.GetMower(id)

	if err != nil {
		return c.Status(http.StatusNotFound).Send([]byte(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(mower)
}

func (serv *CatalogHTTPServer) UpdateMower(c *fiber.Ctx) error {
	c.Append("content-type", JSONContentType)

	id := c.Params("id")

	mowerToUpdate := new(UpdateMowerInputDTO)
	err := c.BodyParser(&mowerToUpdate)

	if err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	mower, err := serv.service.UpdateMower(id, domain.UpdateMowerDTO{
		Name: mowerToUpdate.Name,
	})

	if err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(mower)

}

func (serv *CatalogHTTPServer) GetCatalog(c *fiber.Ctx) error {
	c.Append("content-type", JSONContentType)

	mowers, err := serv.service.GetAvailableMowers()

	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	return c.JSON(mowers)
}
