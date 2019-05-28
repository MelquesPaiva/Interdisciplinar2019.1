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
	Collection = "sensores"
)

func (sensor *SensorDao) Connect() {
	session, err := mgo.Dial(sensor.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(sensor.Database)
}

func (sensor *SensorDao) GetAll() ([]models.Sensor, error) {
	var movies []models.Sensor
	err := db.C(Collection).Find(bson.M{}).All(&movies)
	return movies, err
}
