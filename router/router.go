package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pkc918/chat/controller"
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

	// admin
	admin := api.Group("/admin")
	admin.GET("/ping")

	return r
}
