package service

import (
	"Valter/db/sqlc"
	"Valter/repository"
	"Valter/utility"
	"errors"
	"github.com/gin-gonic/gin"
)

type BookingService interface {
	CreateBooking(c *gin.Context, id, fname, lname, job, emial, phone, country, address, message string, productId uint32) error
	ShowAllBookings(c *gin.Context) ([]sqlc.Booking, error)
}

type bookingService struct {
	bookingRepostory repository.BookingRepository
}

func NewBookingService(bookingRepository repository.BookingRepository) BookingService {
	return &bookingService{bookingRepository}
}

func (bs *bookingService) CreateBooking(c *gin.Context, id, fname, lname, job, emial, phone, country, address, message string, productId uint32) error {
	if fname == "" || lname == "" || job == "" || emial == "" || phone == "" || country == "" || address == "" {
		utility.HttpBadRequest(c, "Fill the empty field")
		return errors.New("values can't be empty")
	}
	data := sqlc.CreateBookingParams{
		ID:        id,
		Fname:     fname,
		Lname:     lname,
		Job:       job,
		Email:     emial,
		Phone:     phone,
		Country:   country,
		Address:   address,
		Message:   message,
		Productid: productId,
	}
	if err := bs.bookingRepostory.CreateBooking(data); err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to create booking", err)
		return err
	}
	return nil
}

func (bs *bookingService) ShowAllBookings(c *gin.Context) ([]sqlc.Booking, error) {
	return bs.bookingRepostory.ShowAllBookings()
}
