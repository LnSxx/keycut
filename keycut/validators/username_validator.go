package validators

import "errors"

func UsernameValidationError(username string) error {
	if len(username) < 4 {
		return errors.New("username must be at least 4 characters long")
	}
	return nil
}
