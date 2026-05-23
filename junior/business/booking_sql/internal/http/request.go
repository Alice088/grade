package http

import "time"

type BookRequest struct {
	UserName string    `json:"user_name" binding:"required,min=2,max=100"`
	RoomName string    `json:"room_name" binding:"required,min=5,max=20"`
	StartAt  time.Time `json:"start_at" binding:"required"`
	EndAt    time.Time `json:"end_at" binding:"required"`
}
