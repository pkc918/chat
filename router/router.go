package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pkc918/chat/controller"
	"github.com/pkc918/chat/middleware"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use()

	// api
	api := r.Group("/api")
	// client
	client := api.Group("/client")
	client.GET("/ping", controller.Pong)
	client.POST("/signUp", controller.SignUp)
	client.POST("/signIn", controller.SignIn)
	client.GET("/captcha", controller.SendCode)

	v1 := client.Group("/v1")
	v1.GET("/ws", controller.ConnectWs)
	v1.Use(middleware.JWTTokenValid())
	v1.GET("/friendList", controller.GetFriendList)

	// admin
	admin := api.Group("/admin")
	admin.GET("/ping")

	return r
}
