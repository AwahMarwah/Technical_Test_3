package cart

import "time"

type Cart struct {
	Id        uint32
	UserId    uint32
	Status    string
	CreatedAt time.Time
}
