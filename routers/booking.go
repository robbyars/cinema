package routers

import (
	"cinema/controllers"
	"cinema/databases/connection"
	"cinema/helpers/common"
	"cinema/middlewares"
	"cinema/repositories"

	"github.com/gin-gonic/gin"
)

func BookingInitiator(router *gin.Engine) {
	api := router.Group("/api/bookings")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateBookingRouter)
		api.GET("", GetBookingByCustomerRouter)
		api.PUT("", UpdateStatusBookingRouter)
	}
}

func CreateBookingRouter(ctx *gin.Context) {
	var (
		bookingRepo = repositories.BookingNewRepository(connection.DBConnections)
		bookingSrv  = controllers.BookingNewService(bookingRepo)
	)

	err := bookingSrv.CreateBookingService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added booking data")
}

func GetBookingByCustomerRouter(ctx *gin.Context) {
	var (
		bookingRepo = repositories.BookingNewRepository(connection.DBConnections)
		bookingSrv  = controllers.BookingNewService(bookingRepo)
	)

	bookings, err := bookingSrv.GetBookingByCustomerService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	totalData := len(bookings)
	common.GenerateSuccessResponseWithListData(ctx, "successfully get booking data", int64(totalData), bookings)
}

func UpdateStatusBookingRouter(ctx *gin.Context) {
	var (
		bookingRepo = repositories.BookingNewRepository(connection.DBConnections)
		bookingSrv  = controllers.BookingNewService(bookingRepo)
	)

	err := bookingSrv.UpdateStatusBookingService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update booking data")
}
