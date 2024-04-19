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
	///middleware///
	app.Use("/beers", beerHandler.AuthRequired)
	///beer///
	app.Get("/beers", beerHandler.GetBeers)
	app.Put("/beers/:id", beerHandler.UpdateBeer)
	app.Delete("/beers/:id", beerHandler.DeleteBeer)
	app.Post("/beers", beerHandler.CreateBeer)
	///register, login///
	app.Post("/register", beerHandler.Register)
	app.Post("/login", beerHandler.Login)
	///company///
	app.Get("/company", beerHandler.GetCompanys)
	app.Post("/company", beerHandler.CreateCompany)
	app.Put("/company/:id", beerHandler.UpdateCompany)
	app.Delete("/company/:id", beerHandler.DeleteCompany)
	///distributer///
	app.Get("/distributer", beerHandler.GetDistributers)
	app.Post("/distributer", beerHandler.CreateDistributer)
	app.Put("/distributer/:id", beerHandler.UpdateDistributer)
	app.Delete("/distributer/:id", beerHandler.DeleteDistributer)
	
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}
