package handler

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	// "gorm.io/gorm"

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

	page := 1
	pageStr := c.Params("page")
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid page number")
		}
	}

	beers, err := h.beerServ.GetBeers(page)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{
		"Beers": beers,
		"PageDetail": m.PaginationData{
			NextPage:     page + 1,
			PreviousPage: page - 1,
			CurrentPage:  page,
		},
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

	beer.ID = uint(id)

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

	file, err := c.FormFile("upload")
	if err != nil {
		return c.JSON(fiber.Map{"message": "created successful, but no file uploaded"})
	}
	err = c.SaveFile(file, "images/"+file.Filename)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "created beer successful"})
}

func (h *beerHandler) Register(c *fiber.Ctx) error {
	user := new(m.User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := h.beerServ.CreateUser(*user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "registed user successful"})
}

func (h *beerHandler) Login(c *fiber.Ctx) error {
	user := new(m.User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	token, err := h.beerServ.LoginUser(*user)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "login successful"})
}

func (h *beerHandler) AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	jwtSecretKey := os.Getenv("jwtSecretKey")

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
