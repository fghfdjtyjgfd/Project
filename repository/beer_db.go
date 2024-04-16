package repository

import "gorm.io/gorm"
import "fmt"
import "github.com/bxcodec/faker/v3"
import "hexTest/core"


////repositiry adabter/////
type beerRepositoryDB struct {
	db *gorm.DB
}

func NewBeerDB(db *gorm.DB) *beerRepositoryDB {
	return &beerRepositoryDB{db: db}
}

func (r *beerRepositoryDB) GetAll() ([]core.Beer, error){
	var beers []core.Beer
	result := r.db.Find(&beers)
	if result.Error != nil {
		return nil, result.Error
	}
	return beers, nil
}

func (r *beerRepositoryDB) UpdateOne(beer core.Beer) error {
	err := r.db.Save(&beer)
	if err != nil{
		return nil
	}
	return nil
}

func (r *beerRepositoryDB) DeleteOne(id int) error {
	var beer core.Beer
	err := r.db.Delete(&beer, id)
	if err != nil {
		return nil
	}
	return nil
}

func (r *beerRepositoryDB) CreateAll(beer core.Beer) error {
	for i := 0; i < 1; i++ {
		r.db.Create(&core.Beer{
			Name:     faker.Word(),
			Type:     faker.Word(),
			Detail:   faker.Paragraph(),
			ImageURL: fmt.Sprintf("http://test.com/%s", faker.UUIDDigit()),
		})
	}
	return nil
}