package main

import (
	"errors"
	"time"
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
	CompanyID string   `json:"companyId,omitempty"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Email     string   `json:"email,omitempty"`
	Password  string   `json:"password,omitempty"`
	IsRoot    bool     `json:"isRoot,omitempty"`
	Roles     []string `json:"roles,omitempty"`
}

func (user *User) createUser() (*User, error) {
	return nil, errors.New("Not implemented")
}

func (user *User) getUser() (*User, error) {
	return nil, errors.New("Not implemented")
}

func (user *User) getUsers() (*User, error) {
	return nil, errors.New("Not implemented")
}

func (user *User) updateUser() (*User, error) {
	return nil, errors.New("Not implemented")
}

func (user *User) deleteUser() (*User, error) {
	return nil, errors.New("Not implemented")
}

func (user *User) authenticateUser() (*User, error) {
	user.Roles = append(user.Roles, "ProductManager")

	return user, nil
}
