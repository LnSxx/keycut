package services_test

import (
	"keycut/keycut/entities"
	"keycut/keycut/services"
	"keycut/keycut/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserService(t *testing.T) {
	// GIVEN
	db, err := test.OpenTestDBConnection()

	assert.Nil(t, err)

	defer db.Close()

	tx, err := db.Begin()

	defer tx.Rollback()

	assert.Nil(t, err)

	service := services.CreateUserService{DBConnection: tx}

	creds := entities.SignUpCredentials{
		Username: "testuser",
		Password: "testpassword",
	}

	// WHEN
	createdUser, err := service.CreateUser(creds)

	// THEN
	assert.Nil(t, err)

	assert.Equal(t, creds.Username, createdUser.Username)

}
