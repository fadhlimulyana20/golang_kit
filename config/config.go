package config

import (
	"os"

	_ "template/utils/env"
)

func init() {
	db := NewDbConfig()
	db.Load()
}

func Env() string {
	return os.Getenv("ENV")
}
