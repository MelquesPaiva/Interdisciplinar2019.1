package models

import(
	"gopkg.in/mgo.v2/bson"
)

type Sensor struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	A0 float64 `bson:"a0" json:"a0"`
	A1 float64 `bson:"a1" json:"a1"`
	A2 float64 `bson:"a2" json:"a2"`
	X float64 `bson:"x" json:"x"`
	Y float64 `bson:"y" json:"y"`
	Z float64 `bson:"z" json:"z"`
}