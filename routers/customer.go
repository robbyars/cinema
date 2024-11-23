package routers

import (
	"cinema/controllers"
	"cinema/databases/connection"
	"cinema/helpers/common"
	"cinema/middlewares"
	"cinema/repositories"

	"github.com/gin-gonic/gin"
)

func CustInitiator(router *gin.Engine) {
	router.POST("/login", Login)
	router.POST("/signup", SignUp)
	api := router.Group("/api/customers")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.GET("", GetAllCustomer)
		api.PUT("/:username", UpdateCustomer)
		api.DELETE("/:username", DeleteCustomer)
	}
}

func Login(ctx *gin.Context) {
	var (
		userRepo = repositories.CustNewRepository(connection.DBConnections)
		userSrv  = controllers.CustNewService(userRepo)
	)

	token, err := userSrv.LoginService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully login", token)
}

func SignUp(ctx *gin.Context) {
	var (
		userRepo = repositories.CustNewRepository(connection.DBConnections)
		userSrv  = controllers.CustNewService(userRepo)
	)

	err := userSrv.SignUpService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully create user")
}

func GetAllCustomer(ctx *gin.Context) {
	var (
		userRepo = repositories.CustNewRepository(connection.DBConnections)
		userSrv  = controllers.CustNewService(userRepo)
	)

	users, err := userSrv.GetAllCustomerService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	totalData := len(users)
	common.GenerateSuccessResponseWithListData(ctx, "successfully get all customer data", int64(totalData), users)
}

func DeleteCustomer(ctx *gin.Context) {
	var (
		userRepo = repositories.CustNewRepository(connection.DBConnections)
		userSrv  = controllers.CustNewService(userRepo)
	)

	err := userSrv.DeleteCustomerService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete customer data")
}

func UpdateCustomer(ctx *gin.Context) {
	var (
		userRepo = repositories.CustNewRepository(connection.DBConnections)
		userSrv  = controllers.CustNewService(userRepo)
	)

	err := userSrv.UpdateCustomerService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update customer data")
}
