package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/stvnreis/mail-sender/src/queue"
	"github.com/stvnreis/mail-sender/src/queue/listener"
	"github.com/stvnreis/mail-sender/src/services"
)

func main() {
	
	godotenv.Load()

	queueName := os.Getenv("QUEUE_NAME")
	from := os.Getenv("EMAIL_SENDER")
	apiKey := os.Getenv("EMAIL_API_KEY")
	smtpHost := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	amqpClient, err := queue.Create(os.Getenv("AMQP_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	defer amqpClient.Close()
	
	mailService := services.CreateMailService(smtpHost, port, from, apiKey)
	listener.Create(amqpClient, queueName, mailService).Listen()
}
