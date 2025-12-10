package main

import (
	"fmt"
	"os"
)

// commands
// join 192.168.0.75
// start => startet chatroom und gibt IP zurück unter der der Raum erreichbar ist

// wenn nutzer nachricht sendet, wird die Nachricht an den Host geschickt, der schickt dann die Nachricht an alle weiter
// soll ich jeden gleich behandeln, also auch dass der Host die Nachricht an sich sendet und dann den send an den Rest triggert?
// soll ich für den host einen webserver aufmachen, an den dann die Nachrichten geschickt werden? => Glaube ja

var isChatroomOpened = false

func main() {
	if os.Args[1] == "start" {
		handleHost()
	} else if os.Args[1] == "join" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: chatty join <ip address>")
		}
		fmt.Println("Joining room with IP: " + os.Args[2])
		handleClient(os.Args[2])
	}
}
