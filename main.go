package main

import (
	"log"
	"os"
	"strings"
	"time"

	"encoding/json"
	"github.com/bobyard/indexer/models"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var engine *xorm.Engine

func Connect() {
	connStr := "user=obj password=!Woaini521 dbname=objdb host=127.0.0.1 port=5432 sslmode=disable"
	var err error
	engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Panicf("%v", err)
	}
}

func main() {
	Connect()

	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))

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
				var listEvent ListEvent
				if err := json.Unmarshal(d.Body, &listEvent); err != nil {
					log.Panicf("%s", err)
				}

				list := new(models.Lists)
				list.ChainId = 1
				list.TokenId = listEvent.MoveEvent.Fields.ListID
				list.SellerAddress = listEvent.MoveEvent.Fields.Owner
				list.SallerValue = 100000000
				list.SellerCoinId = 1
				list.SellerEndTime = time.Now()
				_, err := engine.Insert(list)
				if err != nil {
					log.Printf("%v", err)
				}
				log.Printf("recver list event and sueccess inserted")

			} else if strings.Contains(data, "MarketCreateEvent") {
				var create MarketCreate
				if err := json.Unmarshal(d.Body, &create); err != nil {
					log.Panicf("%s", err)
				}
				log.Printf("recver Market Create Event")
			}

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
