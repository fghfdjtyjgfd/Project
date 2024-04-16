package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"hexTest/core"
)

////handler adabter/////
type beerHandler struct {
	beerServ core.BeerService
}

func NewBeerHandler(beerServ core.BeerService) *beerHandler {
	return &beerHandler{beerServ: beerServ}
}

func (h *beerHandler) GetBeers(c *fiber.Ctx) error {
	beers, err := h.beerServ.GetBeers()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(beers)
}

func (h *beerHandler) UpdateBeer(c *fiber.Ctx) error {
	var beer core.Beer

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := c.BodyParser(&beer); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	beer.ID = id

	err = h.beerServ.UpdateBeer(beer)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "updated successful"})
}

func (h *beerHandler) DeleteBeer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.beerServ.DeleteBeer(id); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "deleted beer successful"})
}

func (h *beerHandler) CreateBeer(c *fiber.Ctx) error {
	var beer core.Beer

	err := c.BodyParser(&beer)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = h.beerServ.CreateBeer(beer)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"message": "created beer successful"})
}
