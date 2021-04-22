package dao

import (
	"log"

	. "github.com/stallkid/ProfileService/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ComponentDAO struct {
	Server   string
	Database string
}

var component_db *mgo.Database

const (
	COMPONENT_COLLECTION = "components"
)

func (m *ComponentDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	component_db = session.DB(m.Database)
}

func (m *ComponentDAO) GetAllComponents() ([]ProfileComponent, error) {
	var components []ProfileComponent
	err := component_db.C(COMPONENT_COLLECTION).Find(bson.M{}).All(&components)
	return components, err
}

func (m *ComponentDAO) GetComponentByID(id string) (ProfileComponent, error) {
	var component ProfileComponent
	err := component_db.C(COMPONENT_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&component)
	return component, err
}

func (m *ComponentDAO) CreateComponent(component ProfileComponent) error {
	err := component_db.C(COMPONENT_COLLECTION).Insert(&component)
	return err
}

func (m *ComponentDAO) DeleteComponent(id string) error {
	err := component_db.C(COMPONENT_COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *ComponentDAO) UpdateComponent(id string, component ProfileComponent) error {
	err := component_db.C(COMPONENT_COLLECTION).UpdateId(bson.ObjectIdHex(id), &component)
	return err
}
