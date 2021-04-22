package profilerouter

import (
	"encoding/json"
	"net/http"

	. "github.com/stallkid/ProfileService/config/dao"
	. "github.com/stallkid/ProfileService/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = ProfileDAO{}

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
	profiles, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, profiles)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profile, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Profile ID")
		return
	}
	respondWithJson(w, http.StatusOK, profile)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	profile.ID = bson.NewObjectId()
	if err := dao.Create(profile); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, profile)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], profile); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": profile.Name + " atualizado com sucesso!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
