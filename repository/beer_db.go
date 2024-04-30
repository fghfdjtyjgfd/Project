package repository

import (
	"fmt"
	// "math/rand"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	m "hexTest/model"
)

// //repositiry adabter/////
type beerRepositoryDB struct {
	db   *gorm.DB
	repo Repo
}

func NewBeerDB(db *gorm.DB) *beerRepositoryDB {
	return &beerRepositoryDB{db: db}
}

func (r *beerRepositoryDB) GetAll() ([]m.Beer, error) {
	var beers []m.Beer

	result := r.db.Preload("Distributer").Preload("Company").Find(&beers)
	if result.Error != nil {
		return nil, result.Error
	}
	return beers, nil
}

func (r *beerRepositoryDB) GetAll2(page int) ([]m.Beer, error) {
	var beers []m.Beer

	perPage := 10
	offSet := (page - 1) * perPage

	result := r.db.Limit(perPage).Offset(offSet).Find(&beers)
	if result.Error != nil {
		return nil, result.Error
	}
	return beers, nil
}

func (r *beerRepositoryDB) UpdateOne(beer m.Beer) error {
	err := r.db.Save(&beer)
	if err != nil {
		return nil
	}
	return nil
}

func (r *beerRepositoryDB) DeleteOne(id int) error {
	var beer m.Beer
	err := r.db.Delete(&beer, id)
	if err != nil {
		return nil
	}
	return nil
}

func (r *beerRepositoryDB) CreateAll(beer m.Beer) error {
	// randomNumber := rand.Intn(101)
	company := m.Company{
		// ID:   uint(randomNumber),
		Name: faker.Word(),
	}
	err := r.repo.Create(r.db, &company)
	if err != nil {
		return err
	}

	distributer := m.Distributer{
		// ID:   uint(randomNumber),
		Name: faker.Word(),
	}
	err = r.repo.Create(r.db, &distributer)
	if err != nil {
		return err
	}

	for i := 0; i < 1; i++ {
		_ = r.repo.Create(r.db, &m.Beer{
			Name:        faker.Word(),
			Type:        faker.Word(),
			Detail:      faker.Paragraph(),
			ImageURL:    fmt.Sprintf("http://test.com/%s", faker.UUIDDigit()),
			CompanyID:   company.ID,
			Distributer: []m.Distributer{distributer},
		})
	}

	return nil
}

func (r *beerRepositoryDB) CreateUser(user m.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	err = r.repo.Create(r.db, &user)
	if err != nil {
		return err
	}
	return nil
}

func (r *beerRepositoryDB) LoginUser(user m.User) (string, error) {
	selectedUser := new(m.User)
	result := r.db.Where("email = ?", user.Email).First(selectedUser)
	if result.Error != nil {
		return "", result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = selectedUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(os.Getenv("jwtSecretKey")))
	if err != nil {
		return "", err
	}
	return t, nil
}
