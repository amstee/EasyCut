package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/user/src/vars"
	"github.com/amstee/easy-cut/utils"
	"github.com/amstee/easy-cut/utils/types"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user vars.UserCreation

	if err := utils.DecodeJSON(&user, r); err != nil {
		utils.ResponseJSON(types.HttpMessage{Message: "unable to decode body", Success: false}, w, http.StatusInternalServerError)
		return
	}

}

func SetUserRoutes(router *mux.Router) {
	router.HandleFunc("", CreateUser).Methods("POST")
	router.HandleFunc("", UpdateUser).Methods("PUT")
	router.HandleFunc("", GetUser).Methods("GET")
}
