package routers

import (
	"cinema/controllers"
	"cinema/databases/connection"
	"cinema/helpers/common"
	"cinema/middlewares"
	"cinema/repositories"

	"github.com/gin-gonic/gin"
)

func Cinema_hallInitiator(router *gin.Engine) {
	api := router.Group("/api/cinema_halls")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateCinema_hallRouter)
		api.GET("", GetAllCinema_hallRouter)
		api.GET("/:id", GetCinema_hallRouter)
		api.PUT("/:id", UpdateCinema_hallRouter)
		api.DELETE("/:id", DeleteCinema_hallRouter)
	}
}

func CreateCinema_hallRouter(ctx *gin.Context) {
	var (
		cinema_hallRepo = repositories.Cinema_hallNewRepository(connection.DBConnections)
		cinema_hallSrv  = controllers.Cinema_hallNewService(cinema_hallRepo)
	)

	err := cinema_hallSrv.CreateCinema_hallService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added cinema_hall data")
}

func GetAllCinema_hallRouter(ctx *gin.Context) {
	var (
		cinema_hallRepo = repositories.Cinema_hallNewRepository(connection.DBConnections)
		cinema_hallSrv  = controllers.Cinema_hallNewService(cinema_hallRepo)
	)

	cinema_halls, err := cinema_hallSrv.GetAllCinema_hallService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	totalData := len(cinema_halls)
	common.GenerateSuccessResponseWithListData(ctx, "successfully get all cinema_hall data", int64(totalData), cinema_halls)
}

func GetCinema_hallRouter(ctx *gin.Context) {
	var (
		cinema_hallRepo = repositories.Cinema_hallNewRepository(connection.DBConnections)
		cinema_hallSrv  = controllers.Cinema_hallNewService(cinema_hallRepo)
	)

	cinema_halls, err := cinema_hallSrv.GetCinema_hallService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get cinema_hall data", cinema_halls)
}

func DeleteCinema_hallRouter(ctx *gin.Context) {
	var (
		cinema_hallRepo = repositories.Cinema_hallNewRepository(connection.DBConnections)
		cinema_hallSrv  = controllers.Cinema_hallNewService(cinema_hallRepo)
	)

	err := cinema_hallSrv.DeleteCinema_hallService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete cinema_hall data")
}

func UpdateCinema_hallRouter(ctx *gin.Context) {
	var (
		cinema_hallRepo = repositories.Cinema_hallNewRepository(connection.DBConnections)
		cinema_hallSrv  = controllers.Cinema_hallNewService(cinema_hallRepo)
	)

	err := cinema_hallSrv.UpdateCinema_hallService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update cinema_hall data")
}
