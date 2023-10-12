package config

import (
	mail "template/utils/mailer"
	"template/utils/minio"

	"gorm.io/gorm"
)

type Config struct {
	ENV    string
	DB     *gorm.DB
	SMTP   *mail.Mailer
	Secret string
	Minio  minio.MinioStorageContract
}
