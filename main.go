package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorgoroth31/chatty/models"
)

// commands
// join 192.168.0.75
// start => startet chatroom und gibt IP zurück unter der der Raum erreichbar ist

// wenn nutzer nachricht sendet, wird die Nachricht an den Host geschickt, der schickt dann die Nachricht an alle weiter
// soll ich jeden gleich behandeln, also auch dass der Host die Nachricht an sich sendet und dann den send an den Rest triggert?
// soll ich für den host einen webserver aufmachen, an den dann die Nachrichten geschickt werden? => Glaube ja

func main() {
	user := models.GetUserInstance()
	user.IpAddr = GetLocalIP().String()

	p := tea.NewProgram(initialModel())

	if os.Args[1] == "start" {
		startServer()
		// joinMySelf
		joinSignalR(GetLocalIP().String(), p)

	} else if os.Args[1] == "join" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: chatty join <ip address>")
		}
		fmt.Println("Joining room with IP: " + os.Args[2])
		handleClient(os.Args[2], p)
	}

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func initialModel() models.ViewModel {
	return models.ViewModel{
		Messages:  []models.Message{},
		InputText: "",
	}
}
