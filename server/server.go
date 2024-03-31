package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkc918/chat/config"
	"github.com/pkc918/chat/db"
	"github.com/pkc918/chat/router"
	"github.com/spf13/viper"
)

func initConfig() {
	config.InitConfig()
}

func initGin() {
	r := router.InitRouter(gin.Default())
	port := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	fmt.Println(port)
	_ = r.Run(port)
}

func initDb() {
	db.InitDb()
}

func InitServer() {
	initConfig()
	initDb()
	initGin()
}
