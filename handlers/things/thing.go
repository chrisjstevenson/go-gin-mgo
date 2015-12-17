package things

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/chrisjstevenson/go-gin-mgo/models"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	things := []models.Thing{}

	err := db.C(models.CollectionStuff).Find(nil).All(&things)
	if err != nil {
		c.Error(err)
	}

	// gin.H is a shortcut for map[string]interface{}
	// c.JSON(http.StatusOK, gin.H{"stuff/list": things})
	c.JSON(http.StatusOK, things)
}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var json models.Thing
	err := c.BindJSON(&json)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionStuff).Insert(json)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, json)
}

func GetOne(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	thing := models.Thing{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionStuff).FindId(oID).One(&thing)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, thing)
}

func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionStuff).RemoveId(oID)
	if err != nil {
		c.Error(err)
	}

	// What to do here if this is close to REST
	//c.Redirect(http.StatusMovedPermanently, "/stuff")
	c.Data(204, "application/json", make([]byte, 0))
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	thing := models.Thing{}
	err := c.Bind(&thing)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{ "_id": bson.ObjectIdHex(c.Param("_id")) }
	doc := bson.M{
		"name":		thing.Name,
		"value":	thing.Value,
	}
	err = db.C(models.CollectionStuff).Update(query, doc)
	if err != nil {
		c.Error(err)
	}

	c.Data(http.StatusOK, "application/json", make([]byte, 0))
	//c.Rediret(http.StatusMovedPermanently, "/stuff"
}
