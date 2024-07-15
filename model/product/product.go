package porduct

import "time"

type Product struct {
	CreatedAt   time.Time
	Id          uint32
	Name        string
	Price       float64
	Description string
}
