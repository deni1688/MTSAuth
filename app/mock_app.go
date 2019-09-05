package app

import (
	"errors"

	"github.com/deni1688/motusauth/models"
	"github.com/stretchr/testify/mock"
)

type MockApp struct {
	mock.Mock
}

func (a *MockApp) AuthenticateUser(u models.User) (string, error) {
	testPassword := hashAndSalt("testing123")

	if err := comparePasswords(testPassword, []byte(u.Password)); err != nil {
		return "", errors.New("Access Denied")
	}

	return "mockToken123", nil
}

// RegisterUser ...
func (a *MockApp) RegisterUser(u *models.User) error {
	return validateUser(u)
}
