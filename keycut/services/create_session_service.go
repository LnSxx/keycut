package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"keycut/keycut/entities"
	"keycut/keycut/storage"
	"keycut/keycut/values"
	"time"
)

type CreateSessionServiceInterface interface {
	CreateSession(values.SessionConfiguration) (entities.Session, error)
}

type CreateSessionService struct {
	DBConnection storage.DatabaseConnection
}

func (service CreateSessionService) CreateSession(config values.SessionConfiguration) (entities.Session, error) {
	createdSession := entities.Session{}

	createdAt := time.Now().UTC()
	expiresAt := createdAt.Add(config.Duration).UTC()

	randomBytes := make([]byte, 32)

	if _, err := rand.Read(randomBytes); err != nil {
		return entities.Session{}, err
	}

	hash := sha256.Sum256(randomBytes)

	token := base64.URLEncoding.EncodeToString(hash[:])

	err := service.DBConnection.QueryRow(`
		INSERT INTO sessions(user_id, name, created_at, expires_at, token) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, name, created_at, expires_at, token`,
		config.UserId, config.Name, createdAt, expiresAt, token).
		Scan(
			&createdSession.Id,
			&createdSession.UserId,
			&createdSession.Name,
			&createdSession.CreatedAt,
			&createdSession.ExpiresAt,
			&createdSession.Token,
		)

	if err != nil {
		return createdSession, err
	}

	return createdSession, nil
}
