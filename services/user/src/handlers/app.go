package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/user/src/vars"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/src/auth0"
	"github.com/amstee/easy-cut/src/logger"
	"github.com/amstee/easy-cut/src/types"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var result common.User
	v := mux.Vars(r)

	if userId, ok := v["user"]; ok {
		token, err := auth0.GetToken(); if err == nil {
			resp, err := request.ExpectJson(config.Content.GetApi() + "users/" + config.Content.Api.Tprefix + userId,
											http.MethodGet, token.Format(), nil, &result)
			if err == nil && resp.StatusCode == 200 {
				common.ResponseJSON(result, w, http.StatusOK); return
			}
			common.ResponseError("unable to update user " + userId, err, w, http.StatusExpectationFailed); return
		}
		common.ResponseError("unable to retrieve api token", err, w, http.StatusInternalServerError); return
	}
	common.ResponseError("user not found in url", nil, w, http.StatusInternalServerError)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user vars.UserUpdate
	var result common.User
	v := mux.Vars(r)

	if userId, ok := v["user"]; ok {
		token, err := auth0.GetToken()
		if err == nil {
			if err := common.DecodeJSON(&user, r); err == nil {
				resp, err := request.ExpectJson(config.Content.GetApi() + "users/" + config.Content.Api.Tprefix + userId,
												http.MethodPatch, token.Format(), user, &result)
				if err == nil && resp.StatusCode == http.StatusOK {
					logger.Info.Println("Updated user: ", userId, "with", user)
					common.ResponseJSON(result, w, http.StatusOK); return
				}
				common.ResponseError("unable to update user", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("unable to retrieve user data", err, w, http.StatusBadRequest); return
		}
		common.ResponseError("unable to retrieve api token", err, w, http.StatusInternalServerError); return
	}
	common.ResponseError("unable to update user", nil, w, http.StatusBadRequest); return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user vars.UserCreation
	var result common.User

	token, err := auth0.GetToken(); if err != nil {
		common.ResponseError("unable to retrieve api token", err, w, http.StatusInternalServerError); return
	}
	if err := common.DecodeJSON(&user, r); err != nil {
		common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
	}
	user.Connection = "Username-Password-Authentication"
	resp, err := request.ExpectJson(config.GetApi() + "users", http.MethodPost, token.Format(), user,  &result)
	if err != nil {
		common.ResponseError("unable to create user", err, w, http.StatusInternalServerError); return
	}
	resp, err = request.ExpectJson(config.GetServiceURL("perms") + "/update/" + result.UserId[6:],
									"PUT", token.Format(), types.CreateGroup("User", false), nil)
	if err == nil && request.IsValid(resp.StatusCode) {
		logger.Info.Println("Created user: ", result)
		common.ResponseJSON(result, w, resp.StatusCode); return
	}
	common.ResponseError("failed to add group to user", err, w, http.StatusInternalServerError); return
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var search vars.Search
	var result []common.User
	vals := r.URL.Query()

	search.Build(vals)
	token, err := auth0.GetToken(); if err != nil {
		common.ResponseError("unable to retrieve api token", err, w, http.StatusInternalServerError); return
	}
	resp, err := request.ExpectJson(config.Content.GetApi() + "users" + search.GetSearch(),
		http.MethodGet, token.Format(), nil, &result)
	if err != nil {
		common.ResponseError("unable to get user list", err, w, http.StatusInternalServerError); return
	}
	common.ResponseJSON(result, w, resp.StatusCode)
}

func SetUserRoutes(router *mux.Router) {
	router.HandleFunc("/create", CreateUser).Methods("POST")
	router.HandleFunc("/list", ListUsers).Methods("GET")
	router.HandleFunc("/update/{user}", UpdateUser).Methods("PUT")
	router.HandleFunc("/get/{user}", GetUser).Methods("GET")
}
