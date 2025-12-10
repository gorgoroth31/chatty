package main

import (
	"context"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	kitlog "github.com/go-kit/log"
	"github.com/gorgoroth31/chatty/models"
	"github.com/philippseith/signalr"
)

type receiver struct {
	signalr.Receiver
	View *tea.Program
}

func (r *receiver) Receive(message models.Message) {
	msg := models.MessageReceivedMsg{Message: message}
	r.View.Send(msg)
}

func handleClient(ipaddr string, p *tea.Program) {
	chatroom := models.GetChatroomInstance()
	chatroom.IsHost = false

	joinSignalR(ipaddr, p)
}

func joinSignalR(ipaddress string, p *tea.Program) {
	go func() {
		address := "http://" + ipaddress + ":8080/chat"

		chatRoom := models.GetChatroomInstance()

		c, err := signalr.NewClient(context.Background(), nil,
			signalr.WithReceiver(&receiver{
				View: p,
			}),
			signalr.WithConnector(func() (signalr.Connection, error) {
				creationCtx, _ := context.WithTimeout(context.Background(), 2*time.Second)
				return signalr.NewHTTPConnection(creationCtx, address)
			}),
			signalr.Logger(kitlog.NewLogfmtLogger(os.Stdout), false))
		if err != nil {
			return
		}
		if err != nil {
			fmt.Println(err)
		}
		c.Start()

		chatRoom.SignalrClient = c

		fmt.Println("Client started")
	}()
}
