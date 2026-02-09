package mailers

import (
	"context"
	"log"
	"time"
)

type SMTPMailer struct {
	host string
	port int
}

func NewSMTPMailer(host string, port int) *SMTPMailer {
	return &SMTPMailer{
		host: host,
		port: port,
	}
}

func (s *SMTPMailer) SendTasksFetchedEmail(ctx context.Context, receiver string) error {
	time.Sleep(time.Second * 2)
	log.Printf("Tasks being sent to: %s\n", receiver)
	return nil
}
