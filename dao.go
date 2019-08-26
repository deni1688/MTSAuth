package main

import "errors"

func (u *User) createUser() (*User, error) {
	return nil, errors.New("na")
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
