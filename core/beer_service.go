package core

import (
	"log"

 m "hexTest/model"
 repo "hexTest/repository"
)

////service port/////
type BeerService interface {
	GetBeers() ([]m.Beer, error)
	UpdateBeer(beer m.Beer) error
	DeleteBeer(int) error
	CreateBeer(beer m.Beer) error
}
////service adabter////
type beerService struct {
	beerRepo repo.BeerRepository
}

func NewBeerService(beerRepo repo.BeerRepository)  *beerService {
	return &beerService{beerRepo: beerRepo}
}

func (s *beerService) GetBeers() ([]m.Beer, error) {
	beers, err := s.beerRepo.GetAll()
	if err != nil {
		log.Println(err)
	}
	
	return beers, nil
}

func (s *beerService) UpdateBeer(beer m.Beer) error {
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

func (s *beerService) CreateBeer(beer m.Beer) error {
	err := s.beerRepo.CreateAll(beer)
	if err != nil {
		log.Println(err)
	}
	return nil
}
