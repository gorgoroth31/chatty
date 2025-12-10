package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id     uuid.UUID
	User   User
	UserIp string
	Text   string
	Time   time.Time
}
