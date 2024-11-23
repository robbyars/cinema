package routers

import (
	"cinema/controllers"
	"cinema/databases/connection"
	"cinema/helpers/common"
	"cinema/middlewares"
	"cinema/repositories"

	"github.com/gin-gonic/gin"
)

func ShowtimeInitiator(router *gin.Engine) {
	api := router.Group("/api/showtimes")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateShowtimeRouter)
		api.GET("", GetAllShowtimeRouter)
		api.GET("/:cinema_id", GetShowtimeByCinemaRouter)
		api.PUT("/:id", UpdateShowtimeRouter)
		api.DELETE("/:id", DeleteShowtimeRouter)
	}
}

func CreateShowtimeRouter(ctx *gin.Context) {
	var (
		showtimeRepo = repositories.ShowtimeNewRepository(connection.DBConnections)
		showtimeSrv  = controllers.ShowtimeNewService(showtimeRepo)
	)

	err := showtimeSrv.CreateShowtimeService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added showtime data")
}

func GetAllShowtimeRouter(ctx *gin.Context) {
	var (
		showtimeRepo = repositories.ShowtimeNewRepository(connection.DBConnections)
		showtimeSrv  = controllers.ShowtimeNewService(showtimeRepo)
	)

	showtimes, err := showtimeSrv.GetAllShowtimeService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	totalData := len(showtimes)
	common.GenerateSuccessResponseWithListData(ctx, "successfully get all showtime data", int64(totalData), showtimes)
}

func GetShowtimeByCinemaRouter(ctx *gin.Context) {
	var (
		showtimeRepo = repositories.ShowtimeNewRepository(connection.DBConnections)
		showtimeSrv  = controllers.ShowtimeNewService(showtimeRepo)
	)

	showtimes, err := showtimeSrv.GetShowtimeByCinemaService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	totalData := len(showtimes)
	common.GenerateSuccessResponseWithListData(ctx, "successfully get showtime data", int64(totalData), showtimes)
}

func DeleteShowtimeRouter(ctx *gin.Context) {
	var (
		showtimeRepo = repositories.ShowtimeNewRepository(connection.DBConnections)
		showtimeSrv  = controllers.ShowtimeNewService(showtimeRepo)
	)

	err := showtimeSrv.DeleteShowtimeService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete showtime data")
}

func UpdateShowtimeRouter(ctx *gin.Context) {
	var (
		showtimeRepo = repositories.ShowtimeNewRepository(connection.DBConnections)
		showtimeSrv  = controllers.ShowtimeNewService(showtimeRepo)
	)

	err := showtimeSrv.UpdateShowtimeService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update showtime data")
}
