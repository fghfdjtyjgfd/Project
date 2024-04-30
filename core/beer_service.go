package core

import (
	m "hexTest/model"
)

// //service port/////
type BeerService interface {
	GetBeers(page int) ([]m.Beer, error)
	TotalPages() (int, error)
	UpdateBeer(beer m.Beer) error
	DeleteBeer(id int) error
	CreateBeer(beer m.Beer) error
	CreateUser(user m.User) error
	LoginUser(user m.User) (string, error)

	GetDis() ([]m.Distributer, error)
	UpdateDis(distributer m.Distributer) error
	DeleteDis(id int) error
	CreateDis(distributer m.Distributer) error

	GetCom() ([]m.Company, error)
	UpdateCom(company m.Company) error
	DeleteCom(id int) error
	CreateCom(company m.Company) error
}
