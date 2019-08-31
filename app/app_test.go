// Package app is the main business logic layer

package app

import (
	"testing"

	"github.com/deni1688/motusauth/models"
)

func TestValidateUser(t *testing.T) {
	type args struct {
		u *models.User
	}

	validUser := &models.User{Email: "test@mail.com", Password: "12345678", FirstName: "Jon", LastName: "Snow"}
	missingEmail := &models.User{Email: "", Password: "minpasslength", FirstName: "Jon", LastName: "Snow"}
	missingPassword := &models.User{Email: "test@mail.com", Password: "", FirstName: "Jon", LastName: "Snow"}
	shortPassword := &models.User{Email: "test@mail.com", Password: "1234567", FirstName: "Jon", LastName: "Snow"}
	missingFirstName := &models.User{Email: "test@mail.com", Password: "12345678", FirstName: "", LastName: "Snow"}
	missingLastName := &models.User{Email: "test@mail.com", Password: "12345678", FirstName: "Jon", LastName: ""}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"should return no error if all required fileds are valid", args{validUser}, false},
		{"should return error if Email is missing", args{missingEmail}, true},
		{"should return error if Password is missing", args{missingPassword}, true},
		{"should return error if Password to short", args{shortPassword}, true},
		{"should return error if FirstName missing", args{missingFirstName}, true},
		{"should return error if LastName missing", args{missingLastName}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
