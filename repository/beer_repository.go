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

	GetAllDis() ([]m.Distributer, error)
	UpdateOneDis(distributer m.Distributer) error
	DeleteOneDis(id int) error
	CreateDis(distributers m.Distributer) error

	GetAllCom() ([]m.Company, error)
	UpdateOneCom(company m.Company) error
	DeleteOneCom(id int) error
	CreateCom(company m.Company) error

}