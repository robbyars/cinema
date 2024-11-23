package repositories

import (
	"cinema/structs"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovieRepository(movie *structs.Movie) (err error)
	GetAllMovieRepository() (movie []structs.Movie, err error)
	GetMovieRepository(movie structs.Movie) (result []structs.Movie, err error)
	DeleteMovieRepository(movie *structs.Movie) (err error)
	UpdateMovieRepository(movie structs.Movie) (err error)
}

type movieRepository struct {
	db *gorm.DB
}

func MovieNewRepository(database *gorm.DB) MovieRepository {
	return &movieRepository{
		db: database,
	}
}

func (r *movieRepository) CreateMovieRepository(movie *structs.Movie) (err error) {
	return r.db.Create(movie).Error
}

func (r *movieRepository) GetAllMovieRepository() (result []structs.Movie, err error) {
	err = r.db.Preload("Showtimes").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *movieRepository) GetMovieRepository(movie structs.Movie) (result []structs.Movie, err error) {
	err = r.db.Where("title LIKE  ?", "%"+movie.Title+"%").Preload("Showtimes").First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, fmt.Errorf("movie with title like %s not found", movie.Title)
		}
		return result, err
	}

	return result, nil
}

func (r *movieRepository) DeleteMovieRepository(movie *structs.Movie) (err error) {
	result := r.db.Where("id = ?", movie.ID).Delete(&structs.Movie{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("movie with id %d not found", movie.ID)
	}

	return nil
}

func (r *movieRepository) UpdateMovieRepository(movie structs.Movie) (err error) {
	result := r.db.Model(&structs.Movie{}).Where("id = ?", movie.ID).
		Updates(map[string]interface{}{
			"title":        movie.Title,
			"genre":        movie.Genre,
			"duration":     movie.Duration,
			"rating":       movie.Rating,
			"release_date": movie.ReleaseDate,
			"description":  movie.Description,
		})
	if result.RowsAffected == 0 {
		return errors.New("movie with the specified id not found")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
