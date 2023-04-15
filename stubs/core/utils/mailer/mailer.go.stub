package mailer

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	dialer *gomail.Dialer
}

type mailer interface {
	SendMail(from string, to string, subject string, template string) error
	GetMailer() *Mailer
}

func NewMailer(host string, port int, authEmail string, password string) mailer {
	return &Mailer{
		dialer: gomail.NewDialer(host, port, authEmail, password),
	}
}

func (m *Mailer) SendMail(from string, to string, subject string, template string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", template)

	if err := m.dialer.DialAndSend(mail); err != nil {
		logrus.Error(err.Error())
		return err
	}

	logrus.Info("mail sent")
	return nil
}

func (m *Mailer) GetMailer() *Mailer {
	return m
}
