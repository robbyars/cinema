package repositories

import (
	"cinema/structs"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Cinema_hallRepository interface {
	CreateCinema_hallRepository(cinema_hall *structs.Cinema_hall) (err error)
	GetAllCinema_hallRepository() (cinema_hall []structs.Cinema_hall, err error)
	GetCinema_hallRepository(cinema_hall structs.Cinema_hall) (result structs.Cinema_hall, err error)
	DeleteCinema_hallRepository(cinema_hall *structs.Cinema_hall) (err error)
	UpdateCinema_hallRepository(cinema_hall structs.Cinema_hall) (err error)
}

type cinema_hallRepository struct {
	db *gorm.DB
}

func Cinema_hallNewRepository(database *gorm.DB) Cinema_hallRepository {
	return &cinema_hallRepository{
		db: database,
	}
}

func (r *cinema_hallRepository) CreateCinema_hallRepository(cinema_hall *structs.Cinema_hall) (err error) {
	return r.db.Create(cinema_hall).Error
}

func (r *cinema_hallRepository) GetAllCinema_hallRepository() (result []structs.Cinema_hall, err error) {
	err = r.db.Preload("Showtimes").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *cinema_hallRepository) GetCinema_hallRepository(cinema_hall structs.Cinema_hall) (result structs.Cinema_hall, err error) {
	err = r.db.Where("id = ?", cinema_hall.ID).Preload("Showtimes").First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, fmt.Errorf("cinema hall with id %d not found", cinema_hall.ID)
		}
		return result, err
	}

	return result, nil
}

func (r *cinema_hallRepository) DeleteCinema_hallRepository(cinema_hall *structs.Cinema_hall) (err error) {
	result := r.db.Where("id = ?", cinema_hall.ID).Delete(&structs.Cinema_hall{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("cinema hall with id %d not found", cinema_hall.ID)
	}

	return nil
}

func (r *cinema_hallRepository) UpdateCinema_hallRepository(cinema_hall structs.Cinema_hall) (err error) {
	result := r.db.Model(&structs.Cinema_hall{}).Where("id = ?", cinema_hall.ID).
		Updates(map[string]interface{}{
			"name":     cinema_hall.Name,
			"capacity": cinema_hall.Capacity,
			"location": cinema_hall.Location,
		})
	if result.RowsAffected == 0 {
		return errors.New("cinema_hall with the specified id not found/something went wrong")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
