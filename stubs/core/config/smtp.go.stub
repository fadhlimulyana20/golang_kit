package config

import (
	"os"
	"strconv"
)

type Smtp struct {
	Host      string
	Port      int
	AuthEmail string
	Password  string
}

type SMTPConfig interface {
	Load() *Smtp
}

func NewSMTPConfig() SMTPConfig {
	return &Smtp{}
}

func (s *Smtp) Load() *Smtp {
	s.Host = os.Getenv("SMTP_HOST")
	s.Port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	s.AuthEmail = os.Getenv("SMTP_USER")
	s.Password = os.Getenv("SMTP_PASSWORD")
	return s
}

func (s *Smtp) Get() *Smtp {
	return s
}
