package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"hexTest/core"
	"hexTest/handler"
	"hexTest/repository"
)

func main() {
	app := fiber.New()
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	beerRepository := repository.NewBeerDB(db)
	beerService := core.NewBeerService(beerRepository)
	beerHandler := handler.NewBeerHandler(beerService)

	app.Get("/beers", beerHandler.GetBeers)
	app.Put("/beers/:id", beerHandler.UpdateBeer)
	app.Delete("/beers/:id", beerHandler.DeleteBeer)
	app.Post("/beers", beerHandler.CreateBeer)

	app.Listen(":8000")
}
