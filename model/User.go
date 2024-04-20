package model

type User struct {
	ID       int    `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
}
