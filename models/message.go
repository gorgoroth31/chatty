package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id   uuid.UUID
	User User
	Text string
	Time time.Time
}
