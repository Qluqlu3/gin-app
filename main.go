package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	name string
}

func main() {
	app := gin.Default()
	app.GET("/index", func(content *gin.Context) {
		content.JSON(http.StatusOK, gin.H{
			// "message": dbGetAll(),
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

func dbGetAll() []User {
	db := gormConnect()
	defer db.Close()
	var users []User
	db.Find(&users)
	return users
}
