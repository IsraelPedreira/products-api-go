package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, db_name, port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("falha ao conectar no banco")
	}

	DB = database
}
