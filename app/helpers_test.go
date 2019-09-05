package app

import (
	"errors"
	"testing"

	"github.com/deni1688/motusauth/models"
	"github.com/stretchr/testify/assert"
)

func Test_hashAndSalt(t *testing.T) {
	type args struct {
		str string
	}

	type test struct {
		desc   string
		expect int
		args   args
	}

	tests := []test{
		{"should return a hash of length 60", 60, args{"password432"}},
		{"should return a hash of length 0", 0, args{""}},
	}

	for _, tt := range tests {
		got := hashAndSalt(tt.args.str)
		assert.Equal(t, tt.expect, len(got), tt.args)
	}
}

func Test_comparePasswords(t *testing.T) {
	type args struct {
		hashedPwd string
		plainPwd  []byte
	}

	type test struct {
		desc   string
		expect error
		args   args
	}

	tests := []test{
		{"should not return error when password match", nil, args{hashAndSalt("password123"), []byte("password123")}},
		{"should return error when passwords do not match", errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password"), args{hashAndSalt("password123"), []byte("password321")}},
	}

	for _, tt := range tests {
		err := comparePasswords(tt.args.hashedPwd, tt.args.plainPwd)
		assert.Equal(t, tt.expect, err, tt.desc)
	}
}

func Test_validateUser(t *testing.T) {
	type args struct {
		u *models.User
	}

	type test struct {
		desc   string
		expect string
		args   args
	}

	validUser := &models.User{Email: "test@mail.com", Password: "12345678", FirstName: "Jon", LastName: "Snow"}
	missingEmail := &models.User{Email: "", Password: "minpasslength", FirstName: "Jon", LastName: "Snow"}
	missingPassword := &models.User{Email: "test@mail.com", Password: "", FirstName: "Jon", LastName: "Snow"}
	shortPassword := &models.User{Email: "test@mail.com", Password: "1234567", FirstName: "Jon", LastName: "Snow"}
	missingFirstName := &models.User{Email: "test@mail.com", Password: "12345678", FirstName: "", LastName: "Snow"}
	missingLastName := &models.User{Email: "test@mail.com", Password: "12345678", FirstName: "Jon", LastName: ""}

	tests := []test{
		{"should return error if Email is missing", "Email is required", args{missingEmail}},
		{"should return error if Password is missing", "Password is required", args{missingPassword}},
		{"should return error if Password to short", "Password min 8 letters", args{shortPassword}},
		{"should return error if FirstName missing", "Firstname is required", args{missingFirstName}},
		{"should return error if LastName missing", "Lastname is required", args{missingLastName}},
	}

	for _, tt := range tests {
		err := validateUser(tt.args.u)
		assert.Equal(t, tt.expect, err.Error(), tt.desc)
	}

	assert.Equal(t, nil, validateUser(validUser), "should return no error if all required fileds are valid")
}
