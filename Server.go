package main

import (
	"log"
	"net/http"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)


func InitServer() error{
	log.Println("Initializing websocket server...")
	//create
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	log.Println("WS server initialized")
	server.On(gosocketio.OnConnection , func (so gosocketio.Channel){
		log.Println("Client connected" + so.Ip())
		so.Join(Room)
	})

	server.On(DataEvent, func(c *gosocketio.Channel, msg DataGram) string {
		c.BroadcastTo(Room, DataEvent, msg)
		return "OK"
	})

	http.Handle("/", http.FileServer(http.Dir("./asset")))
	http.Handle("/socket.io/", server)
	http.ListenAndServe(":5000", nil)
	return nil
}