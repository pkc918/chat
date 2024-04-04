package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkc918/chat/response"
	"github.com/pkc918/chat/service"
)

func GetFriendList(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		response.FailWithResponse(c, response.ErrUserNotFound)
	}
	claims := user.(jwt.MapClaims)
	friendList, err := service.GetFriendList(claims["id"].(string))
	if err != nil {
		response.FailWithResponse(c, err)
	}
	response.OkWithResponse(c, friendList)
}
