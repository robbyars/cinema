package routers

import (
	"cinema/controllers"
	"cinema/databases/connection"
	"cinema/helpers/common"
	"cinema/middlewares"
	"cinema/repositories"

	"github.com/gin-gonic/gin"
)

func MovieInitiator(router *gin.Engine) {
	api := router.Group("/api/movies")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateMovieRouter)
		api.GET("", GetAllMovieRouter)
		api.GET("/:title", GetMovieByNameRouter)
		api.PUT("/:id", UpdateMovieRouter)
		api.DELETE("/:id", DeleteMovieRouter)
	}
}

func CreateMovieRouter(ctx *gin.Context) {
	var (
		movieRepo = repositories.MovieNewRepository(connection.DBConnections)
		movieSrv  = controllers.MovieNewService(movieRepo)
	)

	err := movieSrv.CreateMovieService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added movie data")
}

func GetAllMovieRouter(ctx *gin.Context) {
	var (
		movieRepo = repositories.MovieNewRepository(connection.DBConnections)
		movieSrv  = controllers.MovieNewService(movieRepo)
	)

	movies, err := movieSrv.GetAllMovieService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	totalData := len(movies)
	common.GenerateSuccessResponseWithListData(ctx, "successfully get all movie data", int64(totalData), movies)
}

func GetMovieByNameRouter(ctx *gin.Context) {
	var (
		movieRepo = repositories.MovieNewRepository(connection.DBConnections)
		movieSrv  = controllers.MovieNewService(movieRepo)
	)

	movies, err := movieSrv.GetMovieByNameService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	totalData := len(movies)
	common.GenerateSuccessResponseWithListData(ctx, "successfully get movie data", int64(totalData), movies)
}

func DeleteMovieRouter(ctx *gin.Context) {
	var (
		movieRepo = repositories.MovieNewRepository(connection.DBConnections)
		movieSrv  = controllers.MovieNewService(movieRepo)
	)

	err := movieSrv.DeleteMovieService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete movie data")
}

func UpdateMovieRouter(ctx *gin.Context) {
	var (
		movieRepo = repositories.MovieNewRepository(connection.DBConnections)
		movieSrv  = controllers.MovieNewService(movieRepo)
	)

	err := movieSrv.UpdateMovieService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update movie data")
}
