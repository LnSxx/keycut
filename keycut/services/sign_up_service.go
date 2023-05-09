package services

import "keycut/keycut/storage"

type SignUpService struct {
	DBConnection storage.DatabaseConnection
}
type SignUpServiceInterface interface {
	SignUp(username string, password string) (string, error)
}

func (service SignUpService) SignUp(username string, password string) (string, error) {
	return "", nil
}
