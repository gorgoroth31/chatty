package models

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type ViewModel struct {
	Messages  []Message
	InputText string
}

type MessageReceivedMsg struct {
	Message Message
}

func (m ViewModel) View() string {
	s := "Chatty - Your chat client"

	for _, v := range m.Messages {
		s += "\n" + v.UserIp + ": " + v.Text
	}

	s += "\nInput: " + m.InputText

	return s
}

func (m ViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case MessageReceivedMsg:
		m.Messages = append(m.Messages, msg.Message)
		return m, nil

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			if len(m.InputText) == 0 {
				return m, nil
			}

			chatRoom := GetChatroomInstance()

			message := Message{
				UserIp: GetUserInstance().IpAddr,
				Text:   m.InputText,
				Time:   time.Now(),
				Id:     uuid.New(),
			}

			chatRoom.SendMessage(message)
			m.InputText = ""
			return m, nil

		case "backspace":
			if len(m.InputText) == 0 {
				return m, nil
			}
			m.InputText = m.InputText[:len(m.InputText)-1]
			return m, nil

		case "tab":
			m.InputText += "\t"
			return m, nil

		default:
			m.InputText += msg.String()
			return m, nil
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m ViewModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
