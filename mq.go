package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func connect(connStr string) *amqp.Connection {
	conn, err := amqp.Dial(connStr)
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}
func createChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed on openning channel")
	return ch
}

func DeclareQ(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare("q-logs", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return q

}

func sendData(ch *amqp.Channel, qname string, body []byte, contentType string) {
	err := ch.Publish("",
		qname,
		false, false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        body})
	failOnError(err, "Fail to send JSON")

}
