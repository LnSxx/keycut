package entities

import "keycut/keycut/validators"

type SignUpCredentials struct {
	Username string
	Password string
}

func NewSignUpCredentials(
	username string,
	password string,
) (SignUpCredentials, error) {
	err := validators.UsernameValidationError(username)

	if err != nil {
		return SignUpCredentials{}, err
	}

	err = validators.PasswordValidationError(password)

	if err != nil {
		return SignUpCredentials{}, err
	}

	return SignUpCredentials{
		Username: username,
		Password: password,
	}, nil
}
