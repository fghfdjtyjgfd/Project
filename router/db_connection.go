package router

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	m "hexTest/model"
)

var db *gorm.DB

func NewDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&m.Beer{}, &m.User{})

	return db, nil
}
