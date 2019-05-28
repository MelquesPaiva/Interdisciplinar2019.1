package models

import(
	"gopkg.in/mgo.v2/bson"
)

type Sensor struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
}