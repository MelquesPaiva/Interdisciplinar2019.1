package sensorrouter

import (
	"encoding/json"
	"net/http"
	"go-restapi/config/dao"
	"go-restapi/models"
	"gopkg.in/mgo.v2/bson"
)

var dataaccess = dao.SensorDao{}


func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	sensor, err := dataaccess.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, sensor)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var sensor models.Sensor
	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	sensor.ID = bson.NewObjectId()
	if err := dataaccess.Create(sensor); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, sensor)
}