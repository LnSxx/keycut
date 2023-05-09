package entities

import "time"

type User struct {
	Id        int
	Username  string
	CreatedAt time.Time
}
