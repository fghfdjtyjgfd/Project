package core

/////repository port/////
type BeerRepository interface {
	GetAll() ([]Beer, error)
	UpdateOne(beer Beer) error
	DeleteOne(int) error
	CreateAll(beer Beer) error
}