package main

import (
	"fmt"
	"net/http"

	"github.com/YuvrajSingh3110/go-chat/pkg/webSocket"
)

func serveWS(pool *webSocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket end point reached")

	conn, err := webSocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &webSocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := webSocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Chat app")
	setupRoutes()
	http.ListenAndServe(":5000", nil)
}
