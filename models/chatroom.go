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

func resetInstance() {
	if singleInstance != nil {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Resetting room instance now.")
		singleInstance = &chatroom{}
	}
}
