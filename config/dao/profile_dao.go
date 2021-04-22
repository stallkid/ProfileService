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

var profile_db *mgo.Database

const (
	PROFILE_COLLECTION = "profiles"
)

func (m *ProfileDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	profile_db = session.DB(m.Database)
}

func (m *ProfileDAO) GetAllProfiles() ([]Profile, error) {
	var profiles []Profile
	err := profile_db.C(PROFILE_COLLECTION).Find(bson.M{}).All(&profiles)
	return profiles, err
}

func (m *ProfileDAO) GetProfileByID(id string) (Profile, error) {
	var profile Profile
	err := profile_db.C(PROFILE_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&profile)
	return profile, err
}

func (m *ProfileDAO) CreateProfile(profile Profile) error {
	err := profile_db.C(PROFILE_COLLECTION).Insert(&profile)
	return err
}

func (m *ProfileDAO) DeleteProfile(id string) error {
	err := profile_db.C(PROFILE_COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *ProfileDAO) UpdateProfile(id string, profile Profile) error {
	err := profile_db.C(PROFILE_COLLECTION).UpdateId(bson.ObjectIdHex(id), &profile)
	return err
}
