package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader requuires read and write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//need to find the origin of our con
	//will allow to make req from react
	//to here
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define reader to allow listen for new msgs to our ws ep
func reader(conn *websocket.Conn) {
	for {
		//read msgs
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// define our ws ep
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	//upgrade to ws
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	//listen indefinetely
	reader(ws)
}
func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "simple server")
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
