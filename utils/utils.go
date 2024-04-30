package utils

import (
	"gorm.io/gorm"

	m "hexTest/model"
)

var db *gorm.DB

func FindUserOne(db *gorm.DB, email string, id int) (*m.User, error) {
	user := m.User{}
	err := db.Where("email = ? or id = ?", email, id).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}
