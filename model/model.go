package model

type Beer struct {
	ID uint
	Name     string `json:"name"`
	Type     string `json:"type"`
	Detail   string `json:"detail"`
	ImageURL string `json:"imageurl"`
	CompanyID uint
	Company Company 
	Distributer []Distributer `gorm:"many2many:Distributer_Beer;"`
}

type Company struct {
	ID uint
	Name string `gorm:"unique"`
}

type Distributer struct {
	ID uint
	Name string `gorm:"unique"`
	Beer []Beer `gorm:"many2many:Distributer_Beer;"`
}

type DistributerBeer struct {
	BeerID uint
	Beer Beer
	DistributerID uint
	Distributer Distributer
}