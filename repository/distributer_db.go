package repository

import (
	m "hexTest/model"
)

func (r *beerRepositoryDB) GetAllDis() ([]m.Distributer, error) {
	var distributers []m.Distributer

	result := r.db.Find(&distributers)
	if result.Error != nil {
		return nil, result.Error
	}
	return distributers, nil
}

func (r *beerRepositoryDB) UpdateOneDis(distributer m.Distributer) error {
	err := r.db.Save(&distributer)
	if err != nil {
		return nil
	}
	return nil
}

func (r *beerRepositoryDB) DeleteOneDis(id int) error {
	var distributers m.Distributer
	err := r.db.Delete(&distributers, id)
	if err != nil {
		return nil
	}
	return nil
}

func (r *beerRepositoryDB) CreateDis(distributers m.Distributer) error {
	err := r.db.Create(&distributers)
	if err != nil {
		return nil
	}
	return nil
}
