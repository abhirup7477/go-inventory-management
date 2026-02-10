package mailers

import (
	"context"
	"errors"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type SMTPMailer struct {
	username string
	password string
	host     string
	port     string
}

func NewSMTPMailer() (*SMTPMailer, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("Error loading .env file")
	}
	return &SMTPMailer{
		username: os.Getenv("SMTP_USERNAME"),
		password: os.Getenv("SMTP_PASSWORD"),
		host:     os.Getenv("SMTP_HOST"),
		port:     os.Getenv("SMTP_PORT"),
	}, nil
}

func (s *SMTPMailer) SendTasksFetchedEmail(ctx context.Context, recepient string) error {
	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	addr := s.host + ":" + s.port
	msg := `
		All Tasks have been retreived for you.
		You can check now.
	`

	err := smtp.SendMail(addr, auth, s.username, []string{recepient}, []byte(msg))
	if err != nil {
		return err
	}

	log.Println("Notification email sent to:", recepient)

	return nil
}
