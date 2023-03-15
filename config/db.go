package config

import "os"

type DbConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type DbConfigInf interface {
	Load() *DbConfig
	Get() *DbConfig
}

func NewDbConfig() DbConfigInf {
	return &DbConfig{}
}

func (d *DbConfig) Load() *DbConfig {
	d.Driver = os.Getenv("DB_DRIVER")
	d.Host = os.Getenv("DB_HOST")
	d.Port = os.Getenv("DB_PORT")
	d.User = os.Getenv("DB_USER")
	d.Password = os.Getenv("DB_PASSWORD")
	d.Database = os.Getenv("DB_NAME")
	return d
}

func (d *DbConfig) Get() *DbConfig {
	return d
}
