package things

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/chrisjstevenson/go-gin-mgo/models"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	things := []models.Thing{}

	err := db.C(models.CollectionStuff).Find(nil).Sort("Name").All(&things)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"stuff/list": things})
}

func Create(c *gin.Context) {
	var json models.Thing
	c.BindJSON(&json)

	db := c.MustGet("db").(*mgo.Database)
	err := db.C(models.CollectionStuff).Insert(json)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, json)
}