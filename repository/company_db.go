package repository

import (
	m "hexTest/model"
)

func (r *beerRepositoryDB) GetAllCom() ([]m.Company, error) {
	var companys []m.Company

	result := r.db.Find(&companys)
	if result.Error != nil {
		return nil, result.Error
	}
	return companys, nil
}

func (r *beerRepositoryDB) UpdateOneCom(company m.Company) error {
	err := r.db.Save(&company)
	if err != nil {
		return nil
	}
	return nil
}

func (r *beerRepositoryDB) DeleteOneCom(id int) error {
	var company m.Company
	err := r.db.Delete(&company, id)
	if err != nil {
		return nil
	}
	return nil
}

func (r *beerRepositoryDB) CreateCom(company m.Company) error {
	err := r.db.Create(company)
	if err != nil {
		return nil
	}
	return nil
}
