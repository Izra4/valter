package handler

import (
	"Valter/service"
	"Valter/utility"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type BookingHandler struct {
	bookingService service.BookingService
	productService service.ProductService
}

func NewBookingHandler(bookingService service.BookingService, productService service.ProductService) *BookingHandler {
	return &BookingHandler{bookingService, productService}
}

func (bh *BookingHandler) CreateBooking(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, _ := strconv.Atoi(productIdStr)
	fname := c.PostForm("fname")
	lname := c.PostForm("lname")
	job := c.PostForm("job")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	country := c.PostForm("country")
	address := c.PostForm("address")
	message := c.PostForm("msg")
	id := utility.GenerateInv()

	if err := bh.bookingService.CreateBooking(c, id, fname, lname, job, email, phone, country, address, message, uint32(productId)); err != nil {
		return
	}

	product, err := bh.productService.GetProductsById(c, uint32(productId))
	if err != nil {
		return
	}
	formattedDate := utility.FormatDate(time.Now())
	formattedDueDate := utility.FormatDueDate(time.Now())
	utility.SendMailsBook(email, address, formattedDate, formattedDueDate, id, product.Name, fname)
	utility.HttpSuccessResponse(c, "Success to create a new booking", map[string]any{
		"time":    time.Now(),
		"fname":   fname,
		"lname":   lname,
		"email":   email,
		"country": country,
		"address": address,
		"message": message,
		"product": product.Name,
	})
}
