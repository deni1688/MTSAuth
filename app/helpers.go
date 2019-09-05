package app

import (
	"errors"

	"github.com/deni1688/motusauth/models"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(str string) string {
	if str == "" {
		return ""
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)

	return bcrypt.CompareHashAndPassword(byteHash, plainPwd)
}

// ValidateUser ...
func validateUser(u *models.User) error {
	if u.Email == "" {
		return errors.New("Email is required")
	}

	if u.Password == "" {
		return errors.New("Password is required")
	}

	if len(u.Password) < 8 {
		return errors.New("Password min 8 letters")
	}

	if u.FirstName == "" {
		return errors.New("Firstname is required")
	}

	if u.LastName == "" {
		return errors.New("Lastname is required")
	}

	return nil
}
