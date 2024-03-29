package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pkc918/chat/controller/client"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use()

	// api
	api := r.Group("/api")
	// client
	client := api.Group("/client")
	client.GET("/ping", controller.Pong)

	// admin
	admin := api.Group("/admin")
	admin.GET("/ping")

	return r
}
