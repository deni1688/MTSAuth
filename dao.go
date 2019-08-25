package main

import "errors"

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
