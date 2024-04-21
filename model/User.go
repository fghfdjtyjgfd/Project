package model

type User struct {
	ID       uint    `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
}
