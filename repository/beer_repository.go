package repository


import m "hexTest/model"

/////repository port/////
type BeerRepository interface {
	GetAll() ([]m.Beer, error)
	UpdateOne(beer m.Beer) error
	DeleteOne(int) error
	CreateAll(beer m.Beer) error
}