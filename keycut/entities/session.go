package entities

import "time"

type Session struct {
	Id        int
	UserId    int
	Name      string
	CreatedAt time.Time
	ExpiresAt time.Time
	Token     string
}
