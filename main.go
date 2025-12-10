package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorgoroth31/chatty/models"
)

func main() {
	user := models.GetUserInstance()
	user.IpAddr = GetLocalIP().String()

	p := tea.NewProgram(initialModel())

	if os.Args[1] == "start" {
		startServer()
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
