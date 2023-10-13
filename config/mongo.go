package config

import (
	"os"
	"strconv"
)

type MongoConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type MongoConfigInf interface {
	Load() *MongoConfig
	Get() *MongoConfig
}

func NewMongoConfig() MongoConfigInf {
	return &MongoConfig{}
}

func (d *MongoConfig) Load() *MongoConfig {
	d.Host = os.Getenv("MONGO_HOST")
	d.Port, _ = strconv.Atoi(os.Getenv("MONGO_PORT"))
	d.User = os.Getenv("MONGO_USER")
	d.Password = os.Getenv("MONGO_PASSWORD")
	d.Database = os.Getenv("MONGO_DATABASE")
	return d
}

func (d *MongoConfig) Get() *MongoConfig {
	return d
}
