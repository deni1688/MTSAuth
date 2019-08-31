package models

import (
	"errors"
	"time"

	"github.com/deni1688/motusauth/db"
)

type metadata struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

// User model extended by metadate
type User struct {
	metadata
	CompanyID string `json:"companyId,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `gorm:"unique;not null" json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	IsRoot    bool   `json:"isRoot,omitempty"`
	Roles     string `json:"roles,omitempty"`
}

// Save creates user in the db and returns it
func (u *User) Save() error {
	conn := db.Connect()

	defer conn.Close()

	return conn.Create(&u).Error
}

// Find can return a user based on a query or id
func (u *User) Find(query interface{}) (*User, error) {
	user := &User{}

	conn := db.Connect()

	defer conn.Close()

	if query == nil {
		conn.First(&user, u.ID)
	} else {
		conn.Where(query).First(&user)
	}

	if user == (&User{}) {
		return nil, errors.New("userNotFound")
	}

	return user, nil
}
