package handler

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hexTest/core"
	m "hexTest/model"
)

// //handler adabter/////
type beerHandler struct {
	beerServ core.BeerService
	
}

func NewBeerHandler(beerServ core.BeerService) *beerHandler {
	return &beerHandler{beerServ: beerServ}
}

func (h *beerHandler) GetBeers(c *fiber.Ctx) error {
	var beers []m.Beer

	sql := "SELECT * FROM testdb.beers"

	if name := c.Query("name"); name != "" {
		sql = fmt.Sprintf("%s WHERE Name LIKE '%%%s%%' ", sql, name)
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 10
	var total int64

	db.Raw(sql).Count(&total)

	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

	db.Raw(sql).Scan(&beers)

	return c.JSON(fiber.Map{
		"data":     beers,
		"total":    total,
		"page":     page,
		"lastPage": math.Ceil(float64(total / int64(perPage))),
	})
}

func (h *beerHandler) UpdateBeer(c *fiber.Ctx) error {
	var beer m.Beer

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
	var beer m.Beer

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
