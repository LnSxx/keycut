package services_test

import (
	"keycut/keycut/services"
	"keycut/keycut/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUpService(t *testing.T) {
	// GIVEN
	db, err := test.OpenTestDBConnection()

	assert.Nil(t, err)

	defer db.Close()

	tx, err := db.Begin()

	defer tx.Rollback()

	assert.Nil(t, err)

	service := services.SignUpService{DBConnection: tx}

	username := "testuser"
	password := "testpassword"

	// WHEN
	token, err := service.SignUp(username, password)

	// THEN
	if err != nil {
		t.Errorf("unexpected error during sign up: %v", err)
	}

	if len(token) == 0 {
		t.Errorf("empty token returned by sign up")
	}
}
