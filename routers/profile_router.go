package routers

import (
	"encoding/json"
	"net/http"

	. "github.com/stallkid/ProfileService/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := profileDao.GetAllProfiles()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, profiles)
}

func GetProfileByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profile, err := profileDao.GetProfileByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Profile ID")
		return
	}
	respondWithJson(w, http.StatusOK, profile)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	profile.ID = bson.NewObjectId()
	if err := profileDao.CreateProfile(profile); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, profile)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := profileDao.UpdateProfile(params["id"], profile); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": profile.ID.String() + " atualizado com sucesso!"})
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := profileDao.DeleteProfile(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
