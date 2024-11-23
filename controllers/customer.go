package controllers

import (
	"cinema/helpers/common"
	"cinema/middlewares"
	"cinema/repositories"
	"cinema/structs"

	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type CustService interface {
	LoginService(ctx *gin.Context) (result structs.LoginResponse, err error)
	SignUpService(ctx *gin.Context) (err error)
	UpdateCustomerService(ctx *gin.Context) (err error)
	DeleteCustomerService(ctx *gin.Context) (err error)
	GetAllCustomerService(ctx *gin.Context) (result []structs.Customer, err error)
}

type customerService struct {
	repository repositories.CustRepository
}

func CustNewService(repository repositories.CustRepository) CustService {
	return &customerService{
		repository,
	}
}

func (service *customerService) LoginService(ctx *gin.Context) (result structs.LoginResponse, err error) {
	var userReq structs.LoginRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	err = userReq.ValidateLogin()
	if err != nil {
		return
	}

	user, err := service.repository.Login(userReq)
	if err != nil {
		return
	}

	if common.IsEmptyField(user.Username) {
		err = errors.New("invalid account")
		return
	}

	matches := common.CheckPassword(user.Password, userReq.Password)
	if !matches {
		err = errors.New("wrong username or password")
		return
	}

	jwtToken, err := middlewares.GenerateJwtToken()
	if err != nil {
		return
	}

	middlewares.DummyRedis[jwtToken] = middlewares.UserLoginRedis{
		UserId:    int64(user.ID),
		Username:  user.Username,
		LoginAt:   time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 60),
	}

	result.Token = jwtToken

	return
}

func (service *customerService) SignUpService(ctx *gin.Context) (err error) {
	var userReq structs.SignUpRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return err
	}

	err = userReq.ValidateSignUp()
	if err != nil {
		return err
	}

	user, err := userReq.ConvertToModelForSignUp()
	if err != nil {
		return err
	}

	err = service.repository.SignUp(user)
	if err != nil {
		return err
	}

	return nil
}

func (service *customerService) DeleteCustomerService(ctx *gin.Context) (err error) {
	var cust structs.Customer
	username := ctx.Param("username")

	// Convert id (string) to int
	cust.Username = username

	return service.repository.Delete(cust)
}

func (service *customerService) UpdateCustomerService(ctx *gin.Context) (err error) {
	var cust structs.Customer

	err = ctx.ShouldBind(&cust)
	if err != nil {
		return err
	}

	err = cust.ValidateUpdate()
	if err != nil {
		return err
	}

	username := ctx.Param("username")

	cust.Username = username
	// Panggil repository untuk memperbarui mobil
	return service.repository.Update(cust)
}

func (service *customerService) GetAllCustomerService(ctx *gin.Context) (result []structs.Customer, err error) {
	return service.repository.GetList()
}
