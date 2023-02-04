package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		logrus.Fatal("Failed to connect to DB")
	}

	// defer db.Close()

	return db
}

func ORM() *gorm.DB {
	newLogger := logger.New(
		logrus.New(), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // ambang Slow SQL
			LogLevel:                  logger.Silent, // tingkat Log
			IgnoreRecordNotFoundError: true,          // mengabaikan kesalahan ErrRecordNotFound  untuk logger
			Colorful:                  false,         // nonaktifkan warna
		},
	)

	dsnConn := dsn()
	db, err := gorm.Open(postgres.Open(dsnConn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logrus.Fatal("ORM failed to connect to DB")
	}

	return db
}
