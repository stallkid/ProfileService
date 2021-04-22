package routers

import (
	"encoding/json"
	"net/http"

	. "github.com/stallkid/ProfileService/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func GetAllComponents(w http.ResponseWriter, r *http.Request) {
	components, err := componentDao.GetAllComponents()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, components)
}

func GetComponentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	components, err := componentDao.GetComponentByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Profile ID")
		return
	}
	respondWithJson(w, http.StatusOK, components)
}

func CreateComponents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var component ProfileComponent
	if err := json.NewDecoder(r.Body).Decode(&component); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	component.ID = bson.NewObjectId()
	if err := componentDao.CreateComponent(component); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, component)
}

func UpdateComponents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var component ProfileComponent
	if err := json.NewDecoder(r.Body).Decode(&component); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := componentDao.UpdateComponent(params["id"], component); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": component.Title + " atualizado com sucesso!"})
}

func DeleteComponents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := componentDao.DeleteComponent(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
