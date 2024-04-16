package core

import (
	"log"
)

////service port/////
type BeerService interface {
	GetBeers() ([]Beer, error)
	UpdateBeer(beer Beer) error
	DeleteBeer(int) error
	CreateBeer(beer Beer) error
}
////service adabter////
type beerService struct {
	beerRepo BeerRepository
}

func NewBeerService(beerRepo BeerRepository)  *beerService {
	return &beerService{beerRepo: beerRepo}
}

func (s *beerService) GetBeers() ([]Beer, error) {
	beers, err := s.beerRepo.GetAll()
	if err != nil {
		log.Println(err)
	}
	
	return beers, nil
}

func (s *beerService) UpdateBeer(beer Beer) error {
	err := s.beerRepo.UpdateOne(beer)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *beerService) DeleteBeer(id int) error {
	err := s.beerRepo.DeleteOne(id)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *beerService) CreateBeer(beer Beer) error {
	err := s.beerRepo.CreateAll(beer)
	if err != nil {
		log.Println(err)
	}
	return nil
}
