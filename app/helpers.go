package app

import "golang.org/x/crypto/bcrypt"

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
