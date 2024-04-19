package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	m "hexTest/model"
)

func (h *beerHandler) GetDistributers(c *fiber.Ctx) error {
	distributers, err := h.beerServ.GetDis()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(distributers)
}

func (h *beerHandler) UpdateDistributer(c *fiber.Ctx) error {
	var distributer m.Distributer

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := c.BodyParser(&distributer); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	distributer.ID = uint(id)

	err = h.beerServ.UpdateDis(distributer)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "updated successful"})
}

func (h *beerHandler) DeleteDistributer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.beerServ.DeleteDis(id); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "deleted successful"})
}

func (h *beerHandler) CreateDistributer(c *fiber.Ctx) error {
	var distributer m.Distributer

	err := c.BodyParser(&distributer)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = h.beerServ.CreateDis(distributer)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"message": "created successful"})
}
