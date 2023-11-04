package main

import (
	"fmt"
	"log"
	"net/http"

	"pkg/websocket/"
)

// define our ws ep
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	//upgrade to ws
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	//listen indefinetely
	go websocket.Writer(ws)
	websocket.Reader(ws)
}
func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
