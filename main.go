package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := gin.Default()
	app.GET("/index", func(content *gin.Context) {
		content.JSON(http.StatusOK, gin.H{
			"message": "app",
		})
	})

	app.Run(":5000")
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	DB_NAME := "ginApp"
	CONNECT := USER + ":" + PASS + "@/" + DB_NAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
