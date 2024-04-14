package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkc918/chat/dto"
	"github.com/pkc918/chat/response"
	"github.com/pkc918/chat/service"
	"github.com/pkc918/chat/utils"
	"github.com/pkc918/chat/validator"
	"net/http"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func SignUp(c *gin.Context) {
	account := &dto.SignUpDTO{}
	if err := c.BindJSON(account); err != nil {
		response.FailWithResponse(c, response.ErrParameterIsInvalid)
		return
	}

	if err := validator.Validator(account); err != nil {
		response.FailWithResponse(c, err)
		return
	}

	actInfo, err := service.SignUp(account)
	if err != nil {
		response.FailWithResponse(c, err)
		return
	}
	response.OkWithResponse(c, actInfo)
}

func SignIn(c *gin.Context) {
	account := &dto.SignInDTO{}
	if err := c.BindJSON(account); err != nil {
		response.FailWithResponse(c, response.ErrParameterIsInvalid)
	}

	_account, err := service.SignIn(account)
	if err != nil {
		response.FailWithResponse(c, err)
		return
	}

	signedAccessToken, err := service.CreateJWTToken(_account)
	if err != nil {
		response.FailWithResponse(c, err)
		return
	}

	response.OkWithResponse(c, signedAccessToken)
}

func SendCode(c *gin.Context) {
	toEmailAccount := c.Query("email")
	if err := utils.SendCodeByQQEmail(toEmailAccount); err != nil {
		fmt.Println(err)
		response.FailWithResponse(c, err)
		return
	}
}
