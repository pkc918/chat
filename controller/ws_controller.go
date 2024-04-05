package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkc918/chat/ws"
)

var hub *ws.Hub

func init() {
	hub = ws.NewHub()
	go hub.Run()
}

func ConnectWs(c *gin.Context) {
	ws.ServeWs(hub, c)
}
