package models

type Student struct {
	Name    string  `json:"name"`
	ZipCode string  `json:"zipCode"`
	Age     int     `json:"age"`
	Height  float64 `json:"height"`
	Balance float64 `json:"balance"`
	Email   string  `json:"email"`
	ID      int     `json:"id"`
}
