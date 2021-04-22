package models

import "gopkg.in/mgo.v2/bson"

type Profile struct {
	ID         bson.ObjectId `bson: "_id" json:"id"`
	PersonData []struct {
		ID          bson.ObjectId `bson: "_id" json:"id"`
		Title       string        `bson: "title" json:"title"`
		Data        string        `bson: "Data" json:"data"`
		ComponentID string        `bson: "componentId" json:"componentId"`
	}
}
