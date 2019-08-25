package main

import "time"

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
	Email     string   `gorm:"unique;not null" json:"email,omitempty"`
	Password  string   `json:"password,omitempty"`
	IsRoot    bool     `json:"isRoot,omitempty"`
	Roles     []string `json:"roles,omitempty"`
}
