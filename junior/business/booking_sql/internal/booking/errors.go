package booking

import (
	"errors"
	"net/http"
)

var (
	ErrRoomNameTooLong = errors.New("room name too long")
)

type ErrTimeAlreadyTaken struct {
	msg    string
	fMsg   string
	status int
}

func NewErrTimeAlreadyTaken() error {
	return &ErrTimeAlreadyTaken{
		msg:    "time have already taken",
		fMsg:   "time have already taken",
		status: http.StatusConflict,
	}
}

func (e *ErrTimeAlreadyTaken) Error() string {
	return e.msg
}

func (e *ErrTimeAlreadyTaken) FMsg() string {
	return e.fMsg
}

func (e *ErrTimeAlreadyTaken) Status() int {
	return e.status
}
