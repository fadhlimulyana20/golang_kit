package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func dsn() string {
	dsn := ""
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	if driver == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	}

	return dsn
}

func Connection() *sql.DB {
	dsnConn := dsn()
	db, err := sql.Open("postgres", dsnConn)
	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	// defer db.Close()

	return db
}
