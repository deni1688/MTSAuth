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

func (u *User) Save() (*User, error) {
	conn := db.Connect()

	defer conn.Close()

	err := conn.Create(&u).Error

	return u, err
}

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

func (u *User) FindAll() (*User, error) {
	return nil, errors.New("na")
}

func (u *User) Modify() (*User, error) {
	return nil, errors.New("na")
}

func (u *User) Remove() (*User, error) {
	return nil, errors.New("na")
}
