package repository


import m "hexTest/model"

/////repository port/////
type BeerRepository interface {
	GetAll() ([]m.Beer, error)
	UpdateOne(beer m.Beer) error
	DeleteOne(id int) error
	CreateAll(beer m.Beer) error
	CreateUser(user m.User) error
	LoginUser(user m.User) (string, error)
}