package booking

import "errors"

var (
	ErrRoomNameTooLong = errors.New("room name too long")
)

type ErrTimeAlreadyTaken struct {
	Msg string
}

func (e *ErrTimeAlreadyTaken) Error() string {
	return e.Msg
}

func (e *ErrTimeAlreadyTaken) FMsg() string {
	return e.Msg
}
