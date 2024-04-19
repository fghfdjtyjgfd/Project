package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	m "hexTest/model"
)

func (h *beerHandler) GetCompanys(c *fiber.Ctx) error {
	companys, err := h.beerServ.GetCom()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(companys)
}

func (h *beerHandler) UpdateCompany(c *fiber.Ctx) error {
	var company m.Company

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := c.BodyParser(&company); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	company.ID = uint(id)

	err = h.beerServ.UpdateCom(company)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "updated successful"})
}

func (h *beerHandler) DeleteCompany(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.beerServ.DeleteCom(id); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "deleted successful"})
}

func (h *beerHandler) CreateCompany(c *fiber.Ctx) error {
	var company m.Company

	err := c.BodyParser(&company)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = h.beerServ.CreateCom(company)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"message": "created successful"})
}
