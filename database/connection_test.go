package database

import (
	"log"
	"testing"
)

var (
	driver   = "postgres"
	host     = "localhost"
	port     = "5432"
	user     = "root"
	password = "password"
	database = "sunatlem"
)

func TestSqlDsn(t *testing.T) {
	sqlDB := NewSqlDB(driver, host, port, user, password, database)
	dsn := sqlDB.DSN()
	if dsn == "" {
		t.Fatal("Cannot get DSN")
	}
}

func TestSqlORM(t *testing.T) {
	sqlDB := NewSqlDB(driver, host, port, user, password, database)
	sqlDB.ORM()
}

func TestSqlORMPing(t *testing.T) {
	sqlDB := NewSqlDB(driver, host, port, user, password, database)
	orm := sqlDB.ORM()
	if d, err := orm.DB(); err == nil {
		if err := d.Ping(); err != nil {
			log.Fatal(err)
		}
	}
}
