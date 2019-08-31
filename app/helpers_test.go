package app

import (
	"testing"
)

func Test_hashAndSalt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"should return a hash of length 60", args{"password432"}, 60},
		{"should return a hash of length 0", args{""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashAndSalt(tt.args.str); len(got) != tt.want {
				t.Errorf("hashAndSalt() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_comparePasswords(t *testing.T) {
	type args struct {
		hashedPwd string
		plainPwd  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"should not return error when password match", args{hashAndSalt("password123"), []byte("password123")}, false},
		{"should return error when passwords do not match", args{hashAndSalt("password123"), []byte("password321")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := comparePasswords(tt.args.hashedPwd, tt.args.plainPwd); (err != nil) != tt.wantErr {
				t.Errorf("comparePasswords() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
