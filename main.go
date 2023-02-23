package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"encoding/json"
	"github.com/bobyard/indexer/models"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	APTOS = 0
	SUI   = 1
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
				list.ChainId = SUI
				list.TokenId = listEvent.MoveEvent.Fields.ListID
				list.SellerAddress = listEvent.MoveEvent.Fields.Owner
				s, err := strconv.Atoi(listEvent.MoveEvent.Fields.Ask)
				if err != nil {
					log.Panicf("%v", err)
				}

				list.SallerValue = int64(s)
				list.SellerCoinId = 1
				list.SellerEndTime = time.Now()

				_, err = engine.Insert(list)
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
			} else if strings.Contains(data, "BuyEvent") {
				var buy BuyEvent
				if err := json.Unmarshal(d.Body, &buy); err != nil {
					log.Panicf("%s", err)
				}

				list := new(models.Lists)
				_, err := engine.Where("token_id = ?", buy.MoveEvent.Fields.ListID).Delete(list)
				if err != nil {
					log.Panicf("%s", err)
				}

				// add orders table
				order := new(models.Orders)
				order.TokenId = buy.MoveEvent.Fields.ListID
				order.SellerAddress = buy.MoveEvent.Fields.Owner
				order.BuyerAddress = buy.MoveEvent.Fields.Buyer
				order.Amount = buy.MoveEvent.Fields.Ask
				order.CoinId = SUI
				order.ChainId = 1
				order.Time = time.Now()
				_, err = engine.Insert(list)
				if err != nil {
					log.Printf("%v", err)
				}
				log.Printf("recver Buy event and sueccess inserted")

			} else if strings.Contains(data, "OfferEvent") {
				var offer OfferToNftEvent
				if err := json.Unmarshal(d.Body, &offer); err != nil {
					log.Panicf("%s", err)
				}

				// add orders table
				offerDB := new(models.Offers)
				offerDB.TokenId = offer.MoveEvent.Fields.ListID
				offerDB.OfferId = offer.MoveEvent.Fields.OfferID
				offerDB.ChainId = SUI
				offerDB.BuyerAddress = offer.MoveEvent.Fields.Owner
				offerDB.Item = ""   //TODO
				offerDB.Amount = "" //TODO

				_, err = engine.Insert(offerDB)

				if err != nil {
					log.Printf("%v", err)
				}

				log.Printf("recevr offer")

			} else if strings.Contains(data, "CancelOfferEvent") {
				var cancel CancelOfferEvent
				if err := json.Unmarshal(d.Body, &cancel); err != nil {
					log.Panicf("%s", err)
				}

				offer := new(models.Offers)
				_, err := engine.Where("offer_id = ?", cancel.MoveEvent.Fields.OfferID).Delete(offer)
				if err != nil {
					log.Panicf("%s", err)
				}

				log.Printf("cancel offer")
			} else if strings.Contains(data, "AcceptOfferEvent") {
				var accpet AcceptOfferEvent
				if err := json.Unmarshal(d.Body, &accpet); err != nil {
					log.Panicf("%s", err)
				}

				list := new(models.Lists)
				_, err := engine.Where("token_id = ?", accpet.MoveEvent.Fields.ListID).Delete(list)
				if err != nil {
					log.Panicf("%s", err)
				}
				// let all offer cancel or maybe owner take by self
				offer := new(models.Offers)
				_, err = engine.Where("offer_id = ?", accpet.MoveEvent.Fields.OfferID).Delete(offer)
				if err != nil {
					log.Panicf("%s", err)
				}

				// add orders table
				order := new(models.Orders)
				order.TokenId = accpet.MoveEvent.Fields.ListID
				order.SellerAddress = accpet.MoveEvent.Fields.Owner
				order.BuyerAddress = accpet.MoveEvent.Fields.Buyer
				order.Amount = "1" //TODO make this real
				order.CoinId = SUI
				order.ChainId = 1
				order.Time = time.Now()
				_, err = engine.Insert(list)
				if err != nil {
					log.Printf("%v", err)
				}

			} else {
				log.Printf(" TODO ----------------------")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
