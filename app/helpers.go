package app

import "golang.org/x/crypto/bcrypt"

func hashAndSalt(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)

	if err != nil {
		return ""
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)

	return bcrypt.CompareHashAndPassword(byteHash, plainPwd)
}
