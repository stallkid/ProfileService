package models

import "gopkg.in/mgo.v2/bson"

type ProfileComponent struct {
	ID    bson.ObjectId `bson: "_id" json: "id"`
	Title string        `bson: "title" json:"title"`
}
