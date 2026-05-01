package listener

import (
	"fmt"

	"github.com/stvnreis/mail-sender/src/dto"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stvnreis/mail-sender/src/queue"
	"github.com/stvnreis/mail-sender/src/services"
)

type Listener interface {
	GetMessages() (<-chan amqp.Delivery, error)
}

type EnviarEmailQueueListener struct {
	client *queue.AmqpClient
	queueName string
	mailService *services.MailService
}

func Create(client *queue.AmqpClient, queueName string, mailService *services.MailService) *EnviarEmailQueueListener {
	return &EnviarEmailQueueListener{
		client: client,
		queueName: queueName,
		mailService: mailService,
	}
}

func (l *EnviarEmailQueueListener) GetMessages() (<-chan amqp.Delivery, error) {
	q, err := l.client.GetChannel().QueueDeclare(
		l.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if (err != nil) {
		return nil, err
	}

	msgs, err := l.client.GetChannel().Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (l *EnviarEmailQueueListener) Listen() {

	msgs, err := l.GetMessages()
	if err != nil {
		panic(err)
	}

	var forever chan struct{}

	go func() {
		fmt.Println("Aguardando mensagens...")
		for d := range msgs {
			fmt.Println("Mensagem Recebida -", string(d.Body))

			email, err := dto.FromJson(d.Body)
			if err != nil {
				panic(err)
			}

			l.mailService.SendEmail(email.To, email.Subject, email.Body)
		}
	}()
	<-forever
}