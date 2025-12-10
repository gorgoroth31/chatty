package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorgoroth31/chatty/models"
	"github.com/philippseith/signalr"
)

func startServer() {
	chatroom := models.GetChatroomInstance()
	chatroom.IsHost = true

	go func() {
		hub := chat{}

		server, _ := signalr.NewServer(context.TODO(),
			signalr.SimpleHubFactory(&hub),
			signalr.KeepAliveInterval(2*time.Second))

		router := http.NewServeMux()

		server.MapHTTP(signalr.WithHTTPServeMux(router), "/chat")

		ip := GetLocalIP()
		port := ":8080"
		fmt.Println("Chatroom open on: " + ip.String() + port)
		if err := http.ListenAndServe(ip.String()+port, router); err != nil {
			log.Fatal(err)
		}
	}()
}

type chat struct {
	signalr.Hub
}

func (chat *chat) SendChatMessage(message models.Message) {
	chat.Clients().All().Send("receive", message)
}
