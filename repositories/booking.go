package repositories

import (
	"cinema/structs"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBookingRepository(booking *structs.Booking) (err error)
	GetBookingByCustomerRepository(booking structs.Booking) (result []structs.Booking, err error)
	UpdateStatusBookingRepository(booking structs.Booking) (err error)
}

type bookingRepository struct {
	db *gorm.DB
}

func BookingNewRepository(database *gorm.DB) BookingRepository {
	return &bookingRepository{
		db: database,
	}
}

func (r *bookingRepository) CreateBookingRepository(booking *structs.Booking) (err error) {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetBookingByCustomerRepository(booking structs.Booking) (result []structs.Booking, err error) {
	err = r.db.Where("customer_id = ?", booking.CustomerID).Preload("Customers").Preload("Showtimes").Find(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, fmt.Errorf("booking with customer id %d not found", booking.CustomerID)
		}
		return result, err
	}

	return result, nil
}

func (r *bookingRepository) UpdateStatusBookingRepository(booking structs.Booking) (err error) {
	result := r.db.Model(&structs.Booking{}).Where("customer_id=? AND showtime_id=? AND status=?", booking.CustomerID, booking.ShowtimeID, "Not Paid").
		Updates(map[string]interface{}{
			"status": booking.Status,
		})
	if result.RowsAffected == 0 {
		return errors.New("booking with the specified id not found/something went wrong" + err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
