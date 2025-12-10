package models

import (
	"fmt"
	"sync"

	"github.com/philippseith/signalr"
)

var lock = &sync.Mutex{}

type chatroom struct {
	SignalrClient signalr.Client
	IsHost        bool
}

var singleInstance *chatroom

func GetChatroomInstance() *chatroom {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &chatroom{}
		}
	}

	return singleInstance
}

func (chatRoom *chatroom) SendMessage(message Message) {
	if chatRoom.SignalrClient == nil {
		fmt.Println("Not connected")
		return
	}

	chatRoom.SignalrClient.Send("SendChatMessage", message)
}
