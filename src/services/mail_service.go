package services

import (
	"net/smtp"
	"fmt"
)

type MailService struct {
	host string
	port string
	auth smtp.Auth
	from string
}

func CreateMailService(host string, port string, from string, apiKey string) *MailService {
	return &MailService{
		host: host,
		port: port,
		auth: smtp.PlainAuth("", from, apiKey, host),
		from: from,
	}
}

func (s *MailService) SendEmail(to string, subject string, body string) error {

	fmt.Printf("Enviando email para: %s \r\n", to)

	err := smtp.SendMail(s.host+":"+s.port, s.auth, s.from, []string{to}, []byte("Subject: " + subject+ "\r\n\r\n"+body))
	if err != nil {
		panic(err)
	}

	return nil
}