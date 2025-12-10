package models

import "github.com/google/uuid"

type User struct {
	Id     uuid.UUID
	Name   string
	IpAddr string
	IsHost bool
}

var userInstance *User

func GetUserInstance() *User {
	if userInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if userInstance == nil {
			userInstance = &User{}
			userInstance.Id = uuid.New()
		}
	}

	return userInstance
}
