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

func (user *User) registerRoot() (*User, error) {
	return nil, errors.New("Not implemented")
}

func (user *User) createManager() (*User, error) {
	return nil, errors.New("Not implemented")
}

func (user *User) updateUser() (*User, error) {
	return nil, errors.New("Not implemented")
}

func (user *User) authenticateUser() (*User, error) {
	user.Roles = append(user.Roles, "ProductManager")

	return user, nil
}

func (user *User) removeManagerByID() (*User, error) {
	return nil, errors.New("Not implemented")
}
