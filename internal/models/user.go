package models

import "time"

type User struct {
	ID             int       `json:"id"`
	PassportSeries string    `json:"passport_series"`
	PassportNumber string    `json:"passport_number"`
	Surname        string    `json:"surname"`
	Name           string    `json:"name"`
	Patronymic     string    `json:"patronymic"`
	Address        string    `json:"address"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
