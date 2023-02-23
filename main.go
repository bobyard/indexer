package main

import (
	"log"
	"strings"

	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/guest")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	log.Printf("CONNECT TO RABBITMQ")

	msgs, err := ch.Consume(
		"hello", // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)

	failOnError(err, "Failed to register a consumer")
	var forever chan struct{}

	go func() {

		for d := range msgs {
			log.Printf("message: %s", d.Body)

			data := string(d.Body)
			if strings.Contains(data, "ListEvent") {
				var list ListEvent
				if err := json.Unmarshal(d.Body, &list); err != nil {
					log.Panicf("%s", err)
				}
				log.Printf("%v", list)
				//TODO 入库

			} else if strings.Contains(data, "MarketCreateEvent") {
				var create MarketCreate
				if err := json.Unmarshal(d.Body, &create); err != nil {
					log.Panicf("%s", err)
				}
				log.Printf("%v", create)
			}

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
