package main

import (
	"fmt"
	"github.com/bobyard/indexer/db"
	"github.com/bobyard/indexer/models"
	"github.com/bobyard/indexer/pkg/logger"
	"github.com/bobyard/indexer/suimodels"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"sync"
	"time"

	_ "github.com/bobyard/indexer/pkg/logger"
	"github.com/gorilla/websocket"
	"github.com/panjf2000/ants/v2"
)

func main() {
	defer ants.Release()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bobYardDB := os.Getenv("BOBYARD")
	db.Connect(bobYardDB)

	msgChanel := make(chan []byte, 1000)

	// Use the common pool.
	var wg sync.WaitGroup
	worker := func() {
		for {
			select {
			case msg := <-msgChanel:
				logger.Logger.Info().Msg(string(msg))

				if res := CatchToDB(msg); res {
					logger.Logger.Error().Err(fmt.Errorf("faild to catch to db"))
				}
			}
		}
		wg.Done()
	}

	runTimes := 100
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(worker)
	}

	websocketUri := url.URL{Scheme: "ws", Host: "localhost:9000", Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(websocketUri.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	sendMsg := "{\"jsonrpc\":\"2.0\", \"id\": 1, \"method\": \"sui_subscribeEvent\", \"params\": [{\"All\":[{\"EventType\":\"MoveEvent\"}, {\"Package\":\"0x1647fc0e5f28c100e2c60fac3ddfb15c1bcf1422001a4fbd5f746be0eb64a0c5\"}]}]}"
	err = c.WriteMessage(websocket.TextMessage, []byte(sendMsg))
	if err != nil {
		logger.Logger.Error().Err(err)
		return
	}

	//For the websocket
	wg.Add(1)
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				logger.Logger.Error().Err(err)
			}

			msgChanel <- message
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for {
			collections := make([]*models.Collections, 0)
			err := db.Engine.Where("chain_id = ?", db.SUI).Find(&collections)
			if err != nil {
				logger.Logger.Error().Err(err)
			}

			for _, collection := range collections {
				objects := make([]*suimodels.Objects, 0)
				err = db.Sui.Where("object_type = ?", collection.CollectionId).Find(&objects)
				if err != nil {
					logger.Logger.Error().Err(err)
				}

				collection.Supply = int64(len(objects))
				//TODO find offer and update foolr price
				_, err := db.Engine.Id(collection.Id).Update(collection)
				if err != nil {
					logger.Logger.Error().Err(err)
				}
			}

			m, _ := time.ParseDuration("5s")
			time.Sleep(m)
		}
		wg.Done()
	}()

	wg.Wait()
}
