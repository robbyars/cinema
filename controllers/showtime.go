package controllers

import (
	"cinema/repositories"
	"cinema/structs"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShowtimeService interface {
	CreateShowtimeService(ctx *gin.Context) (err error)
	GetAllShowtimeService(ctx *gin.Context) (result []structs.Showtime, err error)
	GetShowtimeByCinemaService(ctx *gin.Context) (result []structs.Showtime, err error)
	UpdateShowtimeService(ctx *gin.Context) (err error)
	DeleteShowtimeService(ctx *gin.Context) (err error)
}

type showtimeService struct {
	repository repositories.ShowtimeRepository
}

func ShowtimeNewService(repository repositories.ShowtimeRepository) ShowtimeService {
	return &showtimeService{
		repository,
	}
}

func (service *showtimeService) CreateShowtimeService(ctx *gin.Context) (err error) {
	var newShowtime structs.Showtime

	err = ctx.ShouldBind(&newShowtime)
	if err != nil {
		return err
	}

	err = newShowtime.ValidateCreate()
	if err != nil {
		return err
	}

	err = service.repository.CreateShowtimeRepository(&newShowtime)
	if err != nil {
		err = errors.New("failed to add new showtime/something went wrong: " + err.Error())
		return err
	}

	return
}

func (service *showtimeService) GetAllShowtimeService(ctx *gin.Context) (result []structs.Showtime, err error) {
	return service.repository.GetAllShowtimeRepository()
}

func (service *showtimeService) GetShowtimeByCinemaService(ctx *gin.Context) (result []structs.Showtime, err error) {
	var show structs.Showtime
	cinema_id := ctx.Param("cinema_id")

	// Convert id (string) to int
	show.CinemaHallID, err = strconv.Atoi(cinema_id)
	if err != nil {
		err = errors.New("invalid cinema id format")
		return
	}

	// Ambil mobil dari repository
	return service.repository.GetShowtimeByCinemaRepository(show)
}

func (service *showtimeService) DeleteShowtimeService(ctx *gin.Context) (err error) {
	var show structs.Showtime
	id := ctx.Param("id")

	// Convert id (string) to int
	show.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("invalid showtime id format")
		return
	}

	// Panggil repository untuk menghapus mobil
	return service.repository.DeleteShowtimeRepository(&show)
}

func (service *showtimeService) UpdateShowtimeService(ctx *gin.Context) (err error) {
	var show structs.Showtime
	id := ctx.Param("id")

	// Bind request body to Showtime struct
	err = ctx.ShouldBind(&show)
	if err != nil {
		return err
	}

	err = show.ValidateUpdate()
	if err != nil {
		return err
	}

	// Convert id (string) to int
	show.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("invalid showtime id format")
		return
	}

	// Panggil repository untuk memperbarui mobil
	return service.repository.UpdateShowtimeRepository(show)
}
