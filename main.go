package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/indexgit push -u origin ma", func(content *gin.Context) {
		content.JSON(http.StatusOK, gin.H{
			"message": "app",
		})
	})
	app.Run(":5000")
}
