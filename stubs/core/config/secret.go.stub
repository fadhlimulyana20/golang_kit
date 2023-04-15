package config

import "os"

type Secret struct {
	Key    string
	AesKey string
}

type SecretCfg interface {
	Load() *Secret
}

func NewSecretCfg() SecretCfg {
	return &Secret{}
}

func (s *Secret) Load() *Secret {
	s.Key = os.Getenv("SECRET")
	s.AesKey = os.Getenv("AES_SECRET")
	return s
}
