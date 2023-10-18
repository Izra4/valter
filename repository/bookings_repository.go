package repository

import (
	"Valter/db/sqlc"
	"context"
)

type BookingRepository interface {
	CreateBooking(arg sqlc.CreateBookingParams) error
	ShowAllBookings() ([]sqlc.Booking, error)
}

type bookingRepository struct {
	db *sqlc.Queries
}

func NewBookingRepository(db *sqlc.Queries) BookingRepository {
	return &bookingRepository{db}
}

func (br *bookingRepository) CreateBooking(arg sqlc.CreateBookingParams) error {
	_, err := br.db.CreateBooking(context.Background(), arg)
	if err != nil {
		return err
	}
	return nil
}

func (br *bookingRepository) ShowAllBookings() ([]sqlc.Booking, error) {
	data, err := br.db.ShowAllBookings(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}
