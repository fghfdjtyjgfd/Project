package model

type Beer struct {
	ID       int	`json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Detail   string `json:"detail"`
	ImageURL string `json:"imageurl"`
}