package main

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/panjf2000/ants/v2"
)

func main() {
	defer ants.Release()
	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		fmt.Println(1)
		wg.Done()
	}

	runTimes := 100
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()

	websocketUri := url.URL{Scheme: "wss", Host: "fullnode.devnet.sui.io:443", Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(websocketUri.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	sendMsg := "{\"jsonrpc\":\"2.0\", \"id\": 1, \"method\": \"sui_subscribeEvent\", \"params\": [{\"All\":[{\"EventType\":\"MoveEvent\"}, {\"Package\":\"0x2\"}, {\"Module\":\"devnet_nft\"}]}]}"
	err = c.WriteMessage(websocket.TextMessage, []byte(sendMsg))
	if err != nil {
		fmt.Println("write:", err)
		return
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(message))
	}
}
