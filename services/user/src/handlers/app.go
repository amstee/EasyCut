package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/user/src/vars"
	"github.com/amstee/easy-cut/src/types"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
	"fmt"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/src/auth0"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user vars.UserCreation
	var result vars.UserResponse
	token, err := auth0.GetToken(); if err != nil {
		fmt.Println(err)
		common.ResponseJSON(types.HttpMessage{Message: "unable to retrieve api token", Success: false}, w, http.StatusInternalServerError)
		return
	}

	if err := common.DecodeJSON(&user, r); err != nil {
		fmt.Println(err)
		common.ResponseJSON(types.HttpMessage{Message: "unable to decode body", Success: false}, w, http.StatusInternalServerError)
		return
	}
	user.Connection = "Username-Password-Authentication"
	resp, err := request.ExpectJson(config.Content.GetApi() + "users", "POST", token.Format(), user,  &result); if err != nil {
		fmt.Println(err)
		common.ResponseJSON(types.HttpMessage{Message: "unable to create user", Success: false}, w, http.StatusInternalServerError)
		return
	}
	common.ResponseJSON(result, w, resp.StatusCode)
}

func SetUserRoutes(router *mux.Router) {
	router.HandleFunc("/create", CreateUser).Methods("POST")
	router.HandleFunc("/update", UpdateUser).Methods("PUT")
	router.HandleFunc("/get", GetUser).Methods("GET")
}
