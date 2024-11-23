package repositories

import (
	"cinema/structs"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ShowtimeRepository interface {
	CreateShowtimeRepository(showtime *structs.Showtime) (err error)
	GetAllShowtimeRepository() (showtime []structs.Showtime, err error)
	GetShowtimeByCinemaRepository(showtime structs.Showtime) (result []structs.Showtime, err error)
	DeleteShowtimeRepository(showtime *structs.Showtime) (err error)
	UpdateShowtimeRepository(showtime structs.Showtime) (err error)
}

type showtimeRepository struct {
	db *gorm.DB
}

func ShowtimeNewRepository(database *gorm.DB) ShowtimeRepository {
	return &showtimeRepository{
		db: database,
	}
}

func (r *showtimeRepository) CreateShowtimeRepository(showtime *structs.Showtime) (err error) {
	result := r.db.Where("movie_id = ? AND cinema_hall_id = ? AND showtime_date = ?", showtime.MovieID, showtime.CinemaHallID, showtime.ShowtimeDate).Find(&showtime)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 0 {
		return fmt.Errorf("showtime is exists, please check list showtime first!")
	}

	return r.db.Create(showtime).Error
}

func (r *showtimeRepository) GetAllShowtimeRepository() (result []structs.Showtime, err error) {
	err = r.db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *showtimeRepository) GetShowtimeByCinemaRepository(showtime structs.Showtime) (result []structs.Showtime, err error) {
	err = r.db.Where("cinema_hall_id = ?", showtime.CinemaHallID).Find(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, fmt.Errorf("showtime with cinema hall id %d not found", showtime.CinemaHallID)
		}
		return result, err
	}

	return result, nil
}

func (r *showtimeRepository) DeleteShowtimeRepository(showtime *structs.Showtime) (err error) {
	result := r.db.Where("id = ?", showtime.ID).Delete(&structs.Showtime{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("showtime with id %d not found", showtime.ID)
	}

	return nil
}

func (r *showtimeRepository) UpdateShowtimeRepository(showtime structs.Showtime) (err error) {
	result1 := r.db.Where("movie_id = ? AND cinema_hall_id = ? AND showtime_date = ?", showtime.MovieID, showtime.CinemaHallID, showtime.ShowtimeDate).Find(&showtime)
	if result1.Error != nil {
		return result1.Error
	}

	if result1.RowsAffected != 0 {
		return fmt.Errorf("showtime is exists, please check list showtime first")
	}

	result := r.db.Model(&structs.Showtime{}).Where("id = ?", showtime.ID).
		Updates(map[string]interface{}{
			"movie_id":       showtime.MovieID,
			"cinema_hall_id": showtime.CinemaHallID,
			"showtime_date":  showtime.ShowtimeDate,
			"price":          showtime.Price,
		})
	if result.RowsAffected == 0 {
		return errors.New("showtime with the specified id not found/somethineg went wrong" + err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
