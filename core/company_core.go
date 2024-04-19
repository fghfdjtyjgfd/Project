package core

import (
	"log"

	m "hexTest/model"
)

func (s *beerService) GetCom() ([]m.Company, error) {
	companys, err := s.beerRepo.GetAllCom()
	if err != nil {
		log.Println(err)
	}

	return companys, nil
}

func (s *beerService) UpdateCom(company m.Company) error {
	err := s.beerRepo.UpdateOneCom(company)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *beerService) DeleteCom(id int) error {
	err := s.beerRepo.DeleteOneCom(id)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *beerService) CreateCom(company m.Company) error {
	err := s.beerRepo.CreateCom(company)
	if err != nil {
		log.Println(err)
	}
	return nil
}
