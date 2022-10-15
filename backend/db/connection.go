package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	dbname := os.Getenv("POSTGRESQL_DATABASE")
	var DSN = "host" + "=" + host + " user" + "=" + user + " dbname" + "=" + dbname + " password" + "=" + password + " sslmode=disable" + " port" + "=" + port + " TimeZone=America/Santiago"
	//var DSN = "host=postgres user=postgres password=postgres dbname=tarea1 port=5432 sslmode=disable TimeZone=America/Santiago"
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}

}
