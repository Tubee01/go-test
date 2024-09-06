package models

import "time"

type User struct {
	ID             uint       `json:"id"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Email          string     `json:"email"`
	Password       string     `json:"password"`
	CountryNumCode *string    `json:"country_num_code"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
