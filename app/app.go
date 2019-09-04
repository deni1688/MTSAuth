// Package app is the main business logic layer
package app

import (
	"encoding/hex"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/deni1688/motusauth/models"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	FirstName string   `json:"frstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	jwt.StandardClaims
}

type Domain interface {
	RegisterUser(u *models.User) error
	AuthenticateUser(u models.User) (string, error)
}

type App struct{}

var jwtKey = []byte(os.Getenv("MOTUS_JWT_SECRET"))

// RegisterUser ...
func (a *App) RegisterUser(u *models.User) error {
	if err := ValidateUser(u); err != nil {
		return err
	}

	u.CreatedAt = time.Now()
	u.CompanyID = hex.EncodeToString([]byte(u.Email))
	u.Password = hashAndSalt(u.Password)
	u.Roles = "ProductManager,OrderManager,UserManager"
	u.IsRoot = true

	return u.Save()
}

// AuthenticateUser ...
func (a *App) AuthenticateUser(u models.User) (string, error) {
	if u.Password == "" || u.Email == "" {
		return "", errors.New("Email and Password are required")
	}

	user, err := u.Find(&models.User{Email: u.Email})

	if err != nil {
		return "", err
	}

	if err := comparePasswords(user.Password, []byte(u.Password)); err != nil {
		return "", err
	}

	return CreateToken(user)
}

// CreateToken returns a signed JWT token with the users
// firstname, lastname, email, and roles
func CreateToken(u *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claimsExpiration := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	c := claims{
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Email:          u.Email,
		Roles:          strings.Split(u.Roles, ","),
		StandardClaims: claimsExpiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(jwtKey)
}

// ValidateUser ...
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
