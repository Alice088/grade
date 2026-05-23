package booking

import (
	"context"
)

type BookingRepository interface {
	Book(ctx context.Context, booking Booking) (int, error)
}
