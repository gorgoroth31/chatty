package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/gorgoroth31/chatty/models"
	"github.com/philippseith/signalr"
)

type receiver struct {
	signalr.Receiver
}

func (r *receiver) Receive(message models.Message) {
	if message.User.Id == models.GetUserInstance().Id {
		return
	}
	fmt.Println(message.User.Name + ": " + message.Text)
}

func (r *receiver) InfoText(message string) {
	fmt.Println(message)
}

func handleClient(ipaddr string) {
	chatroom := models.GetChatroomInstance()
	chatroom.IsHost = false

	joinSignalR(ipaddr)

	inputLoop()
}

func inputLoop() {
	for {
		if isChatroomOpened {
			break
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Print("Enter your desired username:\n> ")

	in := bufio.NewReader(os.Stdin)

	username, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	user := models.GetUserInstance()

	sanitizedUsername := strings.TrimSpace(username)

	user.Name = sanitizedUsername

	for {
		textIn := bufio.NewReader(os.Stdin)
		text, err := textIn.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		text = strings.TrimSpace(text)

		message := models.Message{
			User: *models.GetUserInstance(),
			Text: text,
			Time: time.Now(),
			Id:   uuid.New(),
		}

		sendMessage(message)
	}
}

func sendMessage(message models.Message) {
	chatRoom := models.GetChatroomInstance()
	if chatRoom.SignalrClient == nil {
		fmt.Println("Not connected")
		return
	}

	chatRoom.SignalrClient.Send("SendChatMessage", message)
}

func joinSignalR(ipaddress string) {
	go func() {
		address := "http://" + ipaddress + ":8080/chat"

		chatRoom := models.GetChatroomInstance()

		c, err := signalr.NewClient(context.Background(), nil,
			signalr.WithReceiver(&receiver{}),
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
		isChatroomOpened = true
	}()
}
