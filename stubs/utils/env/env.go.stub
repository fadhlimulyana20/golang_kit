package env

import (
	"path/filepath"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load(filepath.Join(".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// log.Info(fmt.Sprintf("Run in %s mode", os.Getenv("ENV")))
}
