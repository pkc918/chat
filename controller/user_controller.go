package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkc918/chat/dto"
	"github.com/pkc918/chat/response"
	"github.com/pkc918/chat/service"
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

	actInfo, err := service.SignUp(account)
	if err != nil {
		response.FailWithResponse(c, err)
		return
	}
	response.OkWithResponse(c, actInfo)
}

func SignIn(c *gin.Context) {
	fmt.Println("登录")
}
