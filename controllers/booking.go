package controllers

import (
	"cinema/repositories"
	"cinema/structs"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type BookingService interface {
	CreateBookingService(ctx *gin.Context) (err error)
	GetBookingByCustomerService(ctx *gin.Context) (result []structs.Booking, err error)
	UpdateStatusBookingService(ctx *gin.Context) (err error)
}

type bookingService struct {
	repository repositories.BookingRepository
}

func BookingNewService(repository repositories.BookingRepository) BookingService {
	return &bookingService{
		repository,
	}
}

func (service *bookingService) CreateBookingService(ctx *gin.Context) (err error) {
	var newBooking structs.Booking

	err = ctx.ShouldBind(&newBooking)
	if err != nil {
		return err
	}

	user, exists := ctx.Get("userid")
	if !exists {
		return err
	}
	var number = user.(int64)

	booking := structs.Booking{
		CustomerID:  int(number),
		ShowtimeID:  newBooking.ShowtimeID,
		BookingDate: time.Now(),
		SeatNumber:  newBooking.SeatNumber,
		Status:      "Not Paid",
	}

	err = booking.ValidateCreate()
	if err != nil {
		return err
	}

	err = service.repository.CreateBookingRepository(&booking)
	if err != nil {
		err = errors.New("failed to add new booking/something went wrong: " + err.Error())
		return err
	}

	return
}

func (service *bookingService) GetBookingByCustomerService(ctx *gin.Context) (result []structs.Booking, err error) {
	var bok structs.Booking
	// id := ctx.Param("customer_id")

	// // Convert id (string) to int
	user, exists := ctx.Get("userid")
	if !exists {
		return result, err
	}
	var number = user.(int64)
	bok.CustomerID = int(number)

	// Ambil mobil dari repository
	return service.repository.GetBookingByCustomerRepository(bok)
}

func (service *bookingService) UpdateStatusBookingService(ctx *gin.Context) (err error) {
	var bok structs.Booking

	// Bind request body to Booking struct
	err = ctx.ShouldBind(&bok)
	if err != nil {
		return err
	}

	err = bok.ValidateUpdate()
	if err != nil {
		return err
	}

	user, exists := ctx.Get("userid")
	if !exists {
		return err
	}
	var number = user.(int64)
	bok.CustomerID = int(number)

	// Panggil repository untuk memperbarui mobil
	return service.repository.UpdateStatusBookingRepository(bok)
}
