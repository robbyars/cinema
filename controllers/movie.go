package controllers

import (
	"cinema/repositories"
	"cinema/structs"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieService interface {
	CreateMovieService(ctx *gin.Context) (err error)
	GetAllMovieService(ctx *gin.Context) (result []structs.Movie, err error)
	GetMovieByNameService(ctx *gin.Context) (result []structs.Movie, err error)
	UpdateMovieService(ctx *gin.Context) (err error)
	DeleteMovieService(ctx *gin.Context) (err error)
}

type movieService struct {
	repository repositories.MovieRepository
}

func MovieNewService(repository repositories.MovieRepository) MovieService {
	return &movieService{
		repository,
	}
}

func (service *movieService) CreateMovieService(ctx *gin.Context) (err error) {
	var newMovie structs.Movie

	err = ctx.ShouldBind(&newMovie)
	if err != nil {
		return err
	}

	err = newMovie.ValidateCreate()
	if err != nil {
		return err
	}

	err = service.repository.CreateMovieRepository(&newMovie)
	if err != nil {
		err = errors.New("failed to add new movie")
		return err
	}

	return
}

func (service *movieService) GetAllMovieService(ctx *gin.Context) (result []structs.Movie, err error) {
	return service.repository.GetAllMovieRepository()
}

func (service *movieService) GetMovieByNameService(ctx *gin.Context) (result []structs.Movie, err error) {
	var names structs.Movie
	name := ctx.Param("title")

	// Convert id (string) to int
	names.Title = name

	// Ambil mobil dari repository
	return service.repository.GetMovieRepository(names)
}

func (service *movieService) DeleteMovieService(ctx *gin.Context) (err error) {
	var car structs.Movie
	id := ctx.Param("id")

	// Convert id (string) to int
	car.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("invalid movie id format")
		return
	}

	// Panggil repository untuk menghapus mobil
	return service.repository.DeleteMovieRepository(&car)
}

func (service *movieService) UpdateMovieService(ctx *gin.Context) (err error) {
	var movie structs.Movie
	id := ctx.Param("id")

	// Bind request body to Movie struct
	err = ctx.ShouldBind(&movie)
	if err != nil {
		return err
	}

	err = movie.ValidateUpdate()
	if err != nil {
		return err
	}

	// Convert id (string) to int
	movie.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("invalid movie id format")
		return
	}

	// Panggil repository untuk memperbarui mobil
	return service.repository.UpdateMovieRepository(movie)
}
