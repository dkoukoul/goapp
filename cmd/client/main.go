package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

func connectAndListen(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/goapp/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// Send open command
	err = c.WriteMessage(websocket.TextMessage, []byte("OPEN"))
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Printf("[conn #%d] %s\n", id, message)
	}
}

func main() {
	var numConnections int
	flag.IntVar(&numConnections, "n", 1, "number of parallel connections")
	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i < numConnections; i++ {
		wg.Add(1)
		go connectAndListen(i, &wg)
	}
	wg.Wait()
}
