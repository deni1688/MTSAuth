package main

import "errors"

// User - schema for users
type User struct {
	ID        string   `json:"_id,omitempty"`
	CompanyID string   `json:"companyId,omitempty"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Email     string   `json:"email,omitempty"`
	Password  string   `json:"password,omitempty"`
	CreatedAt string   `json:"createdAt,omitempty"`
	IsRoot    bool     `json:"isRoot,omitempty"`
	Roles     []string `json:"roles,omitempty"`
}

// RegisterRoot - validates and creates a root user
func (user *User) RegisterRoot() (*User, error) {
	return nil, errors.New("Not implemented")
}

// CreateManager - creates a new manager user for a company
func (user *User) CreateManager() (*User, error) {
	return nil, errors.New("Not implemented")
}

// UpdateUser - validates and updates user
func (user *User) UpdateUser() (*User, error) {
	return nil, errors.New("Not implemented")
}

// AuthenticateUser - finds user by id
func (user *User) AuthenticateUser() (*User, error) {
	user.Roles = append(user.Roles, "ProductManager")

	return user, nil
}

// RemoveManagerByID - removes a manager by id
func (user *User) RemoveManagerByID() (*User, error) {
	return nil, errors.New("Not implemented")
}
