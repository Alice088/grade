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
	tx, err := r.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, nil
	}

	
	return 0, nil
}
