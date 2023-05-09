package values

import (
	"time"
)

type SessionConfiguration struct {
	UserId   int
	Name     string
	Duration time.Duration
}
