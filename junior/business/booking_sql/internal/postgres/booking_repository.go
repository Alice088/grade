package postgres

import (
	"alice088/booking_sql/internal/booking"
	"context"

	"github.com/jackc/pgx/v5"
)

type PostgresRepository struct {
	Conn *pgx.Conn
}

func NewBookingRepository(conn *pgx.Conn) booking.BookingRepository {
	return &PostgresRepository{
		Conn: conn,
	}
}

func (r *PostgresRepository) Book(ctx context.Context, booking booking.Booking) (int, error) {
	return 0, nil
}
