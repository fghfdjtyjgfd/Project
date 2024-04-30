package utils

import (
	"math"

	"gorm.io/gorm"

	m "hexTest/model"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
}

var db *gorm.DB

func FindUserOne(db *gorm.DB, email string, id int) (*m.User, error) {
	user := m.User{}
	err := db.Where("email = ? or id = ?", email, id).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

func GetPaginationData(page, perPage int, model interface{}) PaginationData {

	var totalRows int64
	db.Model(model).Count(&totalRows)
	totalPages := math.Ceil(float64(totalRows / int64(perPage)))

	return PaginationData{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		CurrentPage:  page,
		TotalPages:   int(totalPages),
	}
}
