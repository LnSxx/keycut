package services

import (
	"keycut/keycut/entities"
	"keycut/keycut/storage"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	DBConnection storage.DatabaseConnection
}
type CreateUserServiceInterface interface {
	CreateUser(creds entities.SignUpCredentials) (string, error)
}

func (service CreateUserService) CreateUser(creds entities.SignUpCredentials) (entities.User, error) {
	createdUser := entities.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)

	if err != nil {
		return createdUser, err
	}

	err = service.DBConnection.QueryRow(`
		INSERT INTO users(username, password, created_at) 
		VALUES ($1, $2, $3) RETURNING id, username, created_at`,
		creds.Username, hashedPassword, time.Now().UTC()).
		Scan(&createdUser.Id,
			&createdUser.Username,
			&createdUser.CreatedAt)

	if err != nil {
		return createdUser, err
	}

	return createdUser, nil
}
