package main

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MetaData - gorm.Model definition
type MetaData struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

// User - schema for users
type User struct {
	MetaData
	CompanyID string `json:"companyId,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `gorm:"unique;not null" json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	IsRoot    bool   `json:"isRoot,omitempty"`
	Roles     string `json:"roles,omitempty"`
}

var jwtKey = []byte(os.Getenv("MOTUS_JWT_SECRET"))

func (u *User) createUser() *User {
	db := connectDB()

	defer db.Close()

	db.Create(&u)

	return u
}

func (u *User) getUser(query interface{}) (*User, error) {

	user := &User{}

	db := connectDB()

	defer db.Close()

	if query == nil {
		db.First(&user, u.ID)
	} else {
		db.Where(query).First(&user)
	}

	if user != (&User{}) {
		return user, nil
	}

	return nil, errors.New("userNotFound")
}

func (u *User) getUsers() (*User, error) {
	return nil, errors.New("na")
}

func (u *User) updateUser() (*User, error) {
	return nil, errors.New("na")
}

func (u *User) deleteUser() (*User, error) {
	return nil, errors.New("na")
}

func (u *User) authenticate() (*User, error) {
	user, err := u.getUser(&User{Email: u.Email})

	if err != nil {
		return nil, err
	}

	if isValid := comparePasswords(user.Password, []byte(u.Password)); isValid {
		return user, nil
	}

	return nil, errors.New("passwordInvalid")
}

func (u *User) createTokenString() (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claimsExpiration := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	claims := &Claims{
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
