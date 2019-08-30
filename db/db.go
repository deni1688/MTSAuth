package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect ...
func Connect() *gorm.DB {
	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	conn, err := gorm.Open("postgres", dbInfo)

	if err != nil {
		log.Fatal(err)
	}

	return conn
}

// Init ...
func Init(tables []interface{}) {
	conn := Connect()

	defer conn.Close()

	for _, t := range tables {
		conn.AutoMigrate(t)
	}
}
