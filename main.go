package main

import (
	"golang/config"
	"golang/database"
	"golang/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DbConn *gorm.DB

func main() {

	// koneksi db
	conf := config.Config
	DbConn = database.InitDB(conf.Db)
	//export var public ke subfolder services
	services.DbConn = DbConn

	router := gin.Default()

	// gin.SetMode(gin.DebugMode)
	// router := gin.New()

	router.GET("/SearchBarang/:kategori", services.SearchBarangHandler)
	router.GET("/SearchBarangQuery", services.SearchBarangQueryHandler)
	router.GET("/SearchBarangJson", services.SearchBarangJsonHandler)
	router.GET("/SearchBarangHeader", services.SearchBarangHeaderHandler)

	router.Run(":8888")
}
