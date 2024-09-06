package services

import (
	"context"
	"go-test/internal/models"
	"go-test/internal/modules/database"

	"github.com/jackc/pgx/v5"
)

func GetUserById(id int) (models.User, error) {
	rows, err := database.Dbpool.Query(context.Background(), `SELECT * FROM "user" WHERE id = $1`, id)
	if err != nil {
		return models.User{}, err
	}
	rows.Next()

	return scanUser(rows)
}

func GetUserByEmail(email string) (models.User, error) {
	rows, err := database.Dbpool.Query(context.Background(), `SELECT * FROM "user" WHERE email = $1`, email)
	if err != nil {
		return models.User{}, err
	}
	rows.Next()

	return scanUser(rows)
}

func scanUser(row pgx.Row) (models.User, error) {
	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CountryNumCode, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
