package config

import "os"

type minio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
}

type MinioCfg interface {
	Load() *minio
}

func NewMinioCfg() MinioCfg {
	return &minio{}
}

func (m *minio) Load() *minio {
	m.Endpoint = os.Getenv("MINIO_ENDPOINT")
	m.AccessKeyID = os.Getenv("MINIO_ACCESS_KEY_ID")
	m.SecretAccessKey = os.Getenv("MINIO_SECRET_ACCESS_KEY")
	m.BucketName = os.Getenv("MINIO_BUCKET_NAME")

	ssl := os.Getenv("MINIO_USE_SSL")
	if ssl == "True" {
		m.UseSSL = true
	} else {
		m.UseSSL = false
	}

	return m
}
