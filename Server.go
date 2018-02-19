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

	log.Println("ES server initialized")
	server.On(gosocketio.OnConnection , func (so gosocketio.Channel){
		log.Println("Client connected" + so.Ip())
		so.Join(Room)
	})

	server.On("send", func(c *gosocketio.Channel, msg Message) string{

	})
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	http.Handle("/data/", server)
	http.ListenAndServe(":5000", nil)
	return nil
}