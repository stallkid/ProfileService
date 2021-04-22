package dao

import (
	"log"

	. "github.com/stallkid/ProfileService/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProfileDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "profiles"
)

func (m *ProfileDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ProfileDAO) GetAll() ([]Profile, error) {
	var profiles []Profile
	err := db.C(COLLECTION).Find(bson.M{}).All(&profiles)
	return profiles, err
}

func (m *ProfileDAO) GetByID(id string) (Profile, error) {
	var profile Profile
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&profile)
	return profile, err
}

func (m *ProfileDAO) Create(profile Profile) error {
	err := db.C(COLLECTION).Insert(&profile)
	return err
}

func (m *ProfileDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *ProfileDAO) Update(id string, profile Profile) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &profile)
	return err
}
