package booking

import "time"

type Booking struct {
	UserID   int
	RoomName string
	StartAt  time.Time
	EndAt    time.Time
}
