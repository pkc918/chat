package db

import (
	"fmt"
	"github.com/pkc918/chat/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/url"
)

var DB *gorm.DB

func InitDb() {
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	dbname := viper.GetString("datasource.dbname")
	charset := viper.GetString("datasource.charset")
	parseTime := viper.GetString("datasource.parseTime")
	loc := viper.GetString("datasource.loc")
	driverName := viper.GetString("datasource.driverName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		username,
		password,
		host,
		port,
		dbname,
		charset,
		parseTime,
		url.QueryEscape(loc),
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: driverName,
		DSN:        dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})

	if err != nil {
		panic(fmt.Sprintf("fail to connect database, err: %s", err.Error()))
	}
	initAutoMigrate(db)
	DB = db
}

func initAutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.Account{})

	if err != nil {
		panic(fmt.Sprintf("fail to automigrate table, err: %s", err.Error()))
	}
}
