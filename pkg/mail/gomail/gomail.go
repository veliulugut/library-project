package gomail

import (
	"context"
	"fmt"
	"gopkg.in/gomail.v2"
	"library/pkg/mail"
)

var _ mail.Mail = (*GoMail)(nil)

func New(host string, port int, from string, password string) *GoMail {
	return &GoMail{
		Host:     host,
		Port:     port,
		From:     from,
		Password: password,
	}
}

type GoMail struct {
	Host     string
	Port     int
	From     string
	Password string
}

func (g *GoMail) SendResetPasswordEmail(_ context.Context, email, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", g.From)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Reset Password Validation Code")
	m.SetBody("text/plain", fmt.Sprintf("Reset code : %s", otp))

	d := gomail.NewDialer(g.Host, g.Port, g.From, g.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("send reset password email :%w", err)
	}

	return nil
}
