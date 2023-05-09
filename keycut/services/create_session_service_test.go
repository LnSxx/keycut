package services_test

import (
	"keycut/keycut/entities"
	"keycut/keycut/services"
	"keycut/keycut/test"
	"keycut/keycut/values"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateSessionService(t *testing.T) {
	db, err := test.OpenTestDBConnection()

	assert.Nil(t, err)

	defer db.Close()

	tx, err := db.Begin()

	defer tx.Rollback()

	assert.Nil(t, err)

	service := services.CreateSessionService{DBConnection: tx}

	createUserService := services.CreateUserService{DBConnection: tx}

	creds := entities.SignUpCredentials{
		Username: "testuser",
		Password: "testpassword",
	}

	createdUser, err := createUserService.CreateUser(creds)

	assert.Nil(t, err)

	config := values.SessionConfiguration{
		UserId:   createdUser.Id,
		Name:     "macos",
		Duration: time.Duration(time.Duration.Hours(24)),
	}

	session, err := service.CreateSession(config)

	assert.Nil(t, err)

	assert.Equal(t, config.UserId, session.UserId)
	assert.Equal(t, config.Name, session.Name)
	assert.NotEmpty(t, session.Token)
}
