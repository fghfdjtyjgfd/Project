package core

import (
	"log"

	m "hexTest/model"
)

func (s *beerService) GetDis() ([]m.Distributer, error) {
	distributers, err := s.beerRepo.GetAllDis()
	if err != nil {
		log.Println(err)
	}

	return distributers, nil
}

func (s *beerService) UpdateDis(distributer m.Distributer) error {
	err := s.beerRepo.UpdateOneDis(distributer)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *beerService) DeleteDis(id int) error {
	err := s.beerRepo.DeleteOneDis(id)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *beerService) CreateDis(distributer m.Distributer) error {
	err := s.beerRepo.CreateDis(distributer)
	if err != nil {
		log.Println(err)
	}
	return nil
}
