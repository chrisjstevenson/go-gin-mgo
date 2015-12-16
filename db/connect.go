package db


import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session

	Mongo *mgo.DialInfo
)

const(
	MongoDBUrl = "mongodb://localhost:27017/things-and-stuff"
)


func Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}