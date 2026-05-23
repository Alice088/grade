package booking

import "time"

type Booking struct {
	UserName string
	RoomName string
	StartAt  time.Time
	EndAt    time.Time
}
