package errors

type HTTPFriendly interface {
	FMsg() string
	Status() int
	Error() string
}
