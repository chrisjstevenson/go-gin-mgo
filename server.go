package main

import (
	"net/http"

	"github.com/chrisjstevenson/go-gin-mgo/db"
	"github.com/chrisjstevenson/go-gin-mgo/handlers/things"
	"github.com/chrisjstevenson/go-gin-mgo/middlewares"
	"github.com/gin-gonic/gin"
)

const (
	Port = "9002"
)

func init() {
	db.Connect()
}


func main() {
	router := gin.Default()

	router.Use(middlewares.Connect)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "message": http.StatusOK })
	})

	router.GET("/stuff", things.List)
	router.POST("/addStuff", things.Create)

	router.Run(":" + Port)
}