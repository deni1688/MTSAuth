package app

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/deni1688/motusauth/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type _claims struct {
	FirstName string   `json:"frstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	jwt.StandardClaims
}

var jwtKey = []byte(os.Getenv("MOTUS_JWT_SECRET"))

func CreateUser(u *models.User) (*models.User, error) {
	if err := ValidateUser(u); err != nil {
		return nil, err
	}

	companyID := HashAndSalt(u.Email)
	password := HashAndSalt(u.Password)

	u.CompanyID = companyID
	u.Password = password
	u.CreatedAt = time.Now()
	u.Roles = "ProductManager,OrderManager,UserManager"
	u.IsRoot = true

	return u.Save()
}

func AuthenticateUser(u *models.User) (*models.User, error) {
	user, err := u.Find(&models.User{Email: u.Email})

	if err != nil {
		return nil, err
	}

	if isValid := ComparePasswords(user.Password, []byte(u.Password)); isValid {
		return user, nil
	}

	return nil, errors.New("passwordInvalid")
}

func CreateToken(u *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claimsExpiration := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	claims := &_claims{
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Email:          u.Email,
		Roles:          strings.Split(u.Roles, ","),
		StandardClaims: claimsExpiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func ValidateUser(u *models.User) error {
	if u.Email == "" {
		return errors.New("Email is required")
	}

	if u.Password == "" {
		return errors.New("Password is required")
	}

	if len(u.Password) < 8 {
		return errors.New("Password to short")
	}

	if u.FirstName == "" {
		return errors.New("Firstname is required")
	}

	if u.LastName == "" {
		return errors.New("Lastname is required")
	}

	return nil
}

func HashAndSalt(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)

	if err != nil {
		return ""
	}

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	if err != nil {
		log.Println(err)

		return false
	}

	return true
}
