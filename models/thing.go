package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionStuff = "stuff"
)

type Thing struct {
	Id 	  bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string
	Value string
}
