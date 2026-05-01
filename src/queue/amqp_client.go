package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type AmqpClient struct {
	conn *amqp.Connection
	channel *amqp.Channel	
}

func Create(connectionString string) (*AmqpClient, error) {

	connection, err := amqp.Dial(connectionString)
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		conn.Close()
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return &AmqpClient{
		conn: conn,
		channel: channel,
	}, nil
}

func (c *AmqpClient) GetChannel() *amqp.Channel {
	return c.channel
}

func (c *AmqpClient) Close() {
	c.channel.Close()
	c.conn.Close()
}