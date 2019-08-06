package dao

import (
	"log"
	"go-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SensorDao struct {
	Server string 
	Database string
}

var db *mgo.Database

const (
	Collection = "Sensores"
)

func (sensor *SensorDao) Connect() {
	session, err := mgo.Dial(sensor.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(sensor.Database)
}

func (sensor *SensorDao) GetAll() ([]models.Sensor, error) {
	var sensores []models.Sensor
	err := db.C(Collection).Find(bson.M{}).All(&movies)
	return sensores, err
}

func (sensor * SensorDao) Create(modelSensor models.Sensor) error {
	err := db.C(Collection).Insert(modelSensor)
	return err
}
