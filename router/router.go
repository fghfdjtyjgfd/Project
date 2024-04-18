package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"hexTest/core"
	"hexTest/handler"
	"hexTest/repository"
)

func NewRouter() {
	app := fiber.New()
	
	db, err := NewDB()
	if err != nil {
		panic(err)
	}

	beerRepository := repository.NewBeerDB(db)
	beerService := core.NewBeerService(beerRepository)
	beerHandler := handler.NewBeerHandler(beerService)

	app.Use("/beers", beerHandler.AuthRequired)
	app.Get("/beers", beerHandler.GetBeers)
	app.Put("/beers/:id", beerHandler.UpdateBeer)
	app.Delete("/beers/:id", beerHandler.DeleteBeer)
	app.Post("/beers", beerHandler.CreateBeer)
	app.Post("/register", beerHandler.Register)
	app.Post("/login", beerHandler.Login)

	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}
