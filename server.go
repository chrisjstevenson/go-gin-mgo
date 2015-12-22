package main

import (
	"github.com/chrisjstevenson/go-gin-mgo/db"
	"github.com/chrisjstevenson/go-gin-mgo/handlers/things"
	"github.com/chrisjstevenson/go-gin-mgo/middlewares"
	"github.com/gin-gonic/gin"
)

const (
	Port = "9002"
	Prefix = "/api/v1"
)

func init() {
	db.Connect()
}

func main() {
	router := gin.Default()

	router.Use(middlewares.Connect)
	router.GET(Prefix + "/things/:_id", things.GetOne)
	router.GET(Prefix +"/things", things.List)
	router.POST(Prefix + "/things", things.Create)
	router.DELETE(Prefix + "/things/:_id", things.Delete)
	router.PUT(Prefix + "/things/:_id", things.Update)

	router.Run(":" + Port)
}

// https://github.com/WhiteHouse/api-standards
// Version number in the URL
// Allow users to request formats
// URLs include plural nouns.
// Formats should be in the form of api/v1/resource/{id}.json