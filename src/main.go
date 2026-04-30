package main

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stvnreis/mail-sender/src/dto"
)

func main() {
	godotenv.Load()

	from := os.Getenv("EMAIL_SENDER")
	apiKey := os.Getenv("EMAIL_API_KEY")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpServer := smtpHost + ":" + os.Getenv("SMTP_PORT")
	queueName := os.Getenv("QUEUE_NAME")

	auth := smtp.PlainAuth("", from, apiKey, smtpHost)

	connection, err := amqp.Dial(os.Getenv("AMQP_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	q, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	msgs, err := channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Println("Message Received: ", string(d.Body))
			var email dto.EnviarEmailMessageBody

			err := json.Unmarshal(d.Body, &email)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Sending email to %s", email.To)
			err = smtp.SendMail(smtpServer, auth, from, []string{email.To}, []byte("Subject: " + email.Subject+ "\r\n\r\n"+email.Body))
			if err != nil {
				panic(err)
			}
		}
	}()

	fmt.Println("Waiting for messages.")
	<-forever
}
