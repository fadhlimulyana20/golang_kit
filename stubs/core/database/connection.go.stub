package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sqlDBStruct struct {
	driver   string
	host     string
	port     string
	user     string
	password string
	database string
}

type SqlDB interface {
	DSN() string
	ORM() *gorm.DB
}

func NewSqlDB(driver string, host string, port string, user string, password string, database string) SqlDB {
	return &sqlDBStruct{
		driver:   driver,
		host:     host,
		port:     port,
		user:     user,
		password: password,
		database: database,
	}
}

func (s *sqlDBStruct) DSN() string {
	dsn := ""
	if s.driver == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.host, s.port, s.user, s.password, s.database)
	} else if s.driver == "mysql" {
		// Refer to https://github.com/go-sql-driver/mysql#dsn-data-source-name
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", s.user, s.password, s.host, s.port, s.database)
	} else {
		log.Panic("Driver is not supported")
	}

	return dsn
}

func (s *sqlDBStruct) ORM() *gorm.DB {
	newLogger := logger.New(
		logrus.New(), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // ambang Slow SQL
			LogLevel:                  logger.Silent, // tingkat Log
			IgnoreRecordNotFoundError: true,          // mengabaikan kesalahan ErrRecordNotFound  untuk logger
			Colorful:                  false,         // nonaktifkan warna
		},
	)

	dsnConn := s.DSN()
	var db *gorm.DB
	var err error

	if s.driver == "postgres" {
		db, err = gorm.Open(postgres.Open(dsnConn), &gorm.Config{
			Logger: newLogger,
		})
	} else if s.driver == "mysql" {
		db, err = gorm.Open(mysql.Open(dsnConn), &gorm.Config{
			Logger: newLogger,
		})
	} else {
		logrus.Fatal("Invalid DSN or driver")
	}

	if err != nil {
		logrus.Fatal("ORM failed to connect to DB")
	}

	return db
}

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
	} else if driver == "mysql" {
		dsn = fmt.Sprintf("mysql://%s:%s@%s:%s/%s", user, password, host, port, database)
	} else {
		log.Panic("Driver is not supported")
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
