package user

import "time"

type User struct {
	CreatedAt         time.Time
	Email             string
	EncryptedPassword string
	Id                uint32
	Token             string
	UpdatedAt         time.Time
}
