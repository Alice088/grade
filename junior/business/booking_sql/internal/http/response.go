package http

type DefaultResponse[T any] struct {
	OK    bool    `json:"ok"`
	Error *string `json:"error,omitzero"`
	Data  *T      `json:"data,omitzero"`
}

type Book struct {
	ID int `json:"id"`
}

type BookResponse DefaultResponse[Book]

func OK[T any](obj T) DefaultResponse[T] {
	return DefaultResponse[T]{
		OK:   true,
		Data: &obj,
	}
}

func Err(err string) DefaultResponse[string] {
	return DefaultResponse[string]{
		OK:   false,
		Data: &err,
	}
}
