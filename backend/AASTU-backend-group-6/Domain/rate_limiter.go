package domain

import "time"

type RateLimiter struct {
	Email          string
	Count          int
	LastRequest    time.Time
	SuspendedUntil time.Time
}
