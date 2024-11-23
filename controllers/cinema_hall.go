package controllers

import (
	"cinema/repositories"
	"cinema/structs"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CinemaHallService interface {
	CreateCinema_hallService(ctx *gin.Context) (err error)
	GetAllCinema_hallService(ctx *gin.Context) (result []structs.Cinema_hall, err error)
	GetCinema_hallService(ctx *gin.Context) (result structs.Cinema_hall, err error)
	UpdateCinema_hallService(ctx *gin.Context) (err error)
	DeleteCinema_hallService(ctx *gin.Context) (err error)
}

type cinema_hallService struct {
	repository repositories.Cinema_hallRepository
}

func Cinema_hallNewService(repository repositories.Cinema_hallRepository) CinemaHallService {
	return &cinema_hallService{
		repository,
	}
}

func (service *cinema_hallService) CreateCinema_hallService(ctx *gin.Context) (err error) {
	var newCinema_hall structs.Cinema_hall

	err = ctx.ShouldBind(&newCinema_hall)
	if err != nil {
		return err
	}

	err = newCinema_hall.ValidateCreate()
	if err != nil {
		return err
	}

	err = service.repository.CreateCinema_hallRepository(&newCinema_hall)
	if err != nil {
		err = errors.New("failed to add new cinema_hall/something went wrong: " + err.Error())
		return err
	}

	return
}

func (service *cinema_hallService) GetAllCinema_hallService(ctx *gin.Context) (result []structs.Cinema_hall, err error) {
	return service.repository.GetAllCinema_hallRepository()
}

func (service *cinema_hallService) GetCinema_hallService(ctx *gin.Context) (result structs.Cinema_hall, err error) {
	var car structs.Cinema_hall
	id := ctx.Param("id")

	// Convert id (string) to int
	car.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("invalid cinema_hall id format")
		return
	}

	// Ambil mobil dari repository
	return service.repository.GetCinema_hallRepository(car)
}

func (service *cinema_hallService) DeleteCinema_hallService(ctx *gin.Context) (err error) {
	var car structs.Cinema_hall
	id := ctx.Param("id")

	// Convert id (string) to int
	car.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("invalid cinema_hall id format")
		return
	}

	// Panggil repository untuk menghapus mobil
	return service.repository.DeleteCinema_hallRepository(&car)
}

func (service *cinema_hallService) UpdateCinema_hallService(ctx *gin.Context) (err error) {
	var cinema structs.Cinema_hall
	id := ctx.Param("id")

	// Bind request body to Cinema_hall struct
	err = ctx.ShouldBind(&cinema)
	if err != nil {
		return err
	}

	err = cinema.ValidateUpdate()
	if err != nil {
		return err
	}

	// Convert id (string) to int
	cinema.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("invalid cinema_hall id format")
		return
	}

	// Panggil repository untuk memperbarui mobil
	return service.repository.UpdateCinema_hallRepository(cinema)
}
