package services

import (
	"context"
	"go-test/internal/models"
	"go-test/internal/modules"
	"go-test/internal/modules/database"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(data *LoginData) (models.User, error) {
	user, err := GetUserByEmail(data.Email)
	if err != nil {
		return models.User{}, err
	}

	if !modules.CheckPasswordHash(data.Password, user.Password) {
		return models.User{}, nil
	}

	if user.DeletedAt != nil {
		return models.User{}, nil
	}

	return user, nil
}

type RegisterData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func Register(data *RegisterData) error {
	var user models.User
	hash, err := modules.HashPassword(data.Password)
	if err != nil {
		return err
	}

	err = database.Dbpool.QueryRow(
		context.Background(),
		`INSERT INTO "user" (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id`,
		data.FirstName, data.LastName, data.Email, hash,
	).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
