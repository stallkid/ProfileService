package routers

import (
	"encoding/json"
	"net/http"

	. "github.com/stallkid/ProfileService/config/dao"
)

var componentDao = ComponentDAO{}
var profileDao = ProfileDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
