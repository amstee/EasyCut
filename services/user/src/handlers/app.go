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
	var result vars.UserResponse
	v := mux.Vars(r)

	if userId, ok := v["user"]; ok {
		if token, err := auth0.GetToken(); err == nil {
			resp, err := request.ExpectJson(config.Content.GetApi() + "users/" + config.Content.TPrefix + userId,
											http.MethodGet, token.Format(), nil, &result)
			if err == nil && resp.StatusCode == 200 {
				common.ResponseJSON(result, w, http.StatusOK)
				return
			}
			common.ResponseJSON(types.HttpMessage{Message: "unable to get user " + userId, Success: false},
								w, http.StatusExpectationFailed)
			return
		}
		common.ResponseJSON(types.HttpMessage{Message: "unable to retrieve api token", Success: false},
							w, http.StatusInternalServerError)
		return
	}
	common.ResponseJSON(types.HttpMessage{Message: "user not found in url", Success: false},
						w, http.StatusBadRequest)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user vars.UserUpdate
	var result vars.UserResponse
	v := mux.Vars(r)

	if userId, ok := v["user"]; ok {
		token, err := auth0.GetToken()
		if err == nil {
			if err := common.DecodeJSON(&user, r); err == nil {
				resp, err := request.ExpectJson(config.Content.GetApi() + "users/" + config.Content.TPrefix + userId,
												http.MethodPatch, token.Format(), user, &result)
				if err == nil && resp.StatusCode == http.StatusOK {
					common.ResponseJSON(result, w, http.StatusOK)
					return
				}
				fmt.Println(err)
				common.ResponseJSON(types.HttpMessage{Message: "unable to update user", Success: false}, w,
					http.StatusInternalServerError)
				return
			}
			fmt.Println(err)
			common.ResponseJSON(types.HttpMessage{Message: "unable to retrieve user data", Success: false},
				w, http.StatusBadRequest)
			return
		}
		fmt.Println(err)
		common.ResponseJSON(types.HttpMessage{Message: "unable to retrieve api token", Success: false},
			w, http.StatusInternalServerError)
		return
	}
	common.ResponseJSON(types.HttpMessage{Message: "user not found in url", Success: false},
						w, http.StatusBadRequest)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user vars.UserCreation
	var result vars.UserResponse
	token, err := auth0.GetToken(); if err != nil {
		fmt.Println(err)
		common.ResponseJSON(types.HttpMessage{Message: "unable to retrieve api token", Success: false},
							w, http.StatusInternalServerError)
		return
	}

	if err := common.DecodeJSON(&user, r); err != nil {
		fmt.Println(err)
		common.ResponseJSON(types.HttpMessage{Message: "unable to decode body", Success: false},
							w, http.StatusInternalServerError)
		return
	}
	user.Connection = "Username-Password-Authentication"
	resp, err := request.ExpectJson(config.Content.GetApi() + "users", http.MethodPost, token.Format(), user,  &result)
	if err != nil {
		fmt.Println(err)
		common.ResponseJSON(types.HttpMessage{Message: "unable to create user", Success: false},
							w, http.StatusInternalServerError)
		return
	}
	common.ResponseJSON(result, w, resp.StatusCode)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var search vars.Search
	var result []vars.UserResponse
	vals := r.URL.Query()

	search.Build(vals)
	token, err := auth0.GetToken(); if err != nil {
		fmt.Println(err)
		common.ResponseJSON(types.HttpMessage{Message: "unable to retrieve api token", Success: false},
							w, http.StatusInternalServerError)
		return
	}
	fmt.Println(search.GetSearch())
	resp, err := request.ExpectJson(config.Content.GetApi() + "users" + search.GetSearch(),
		http.MethodGet, token.Format(), nil, &result)
	if err != nil {
		fmt.Println(err)
		common.ResponseJSON(types.HttpMessage{Message: "unable to get users list", Success: false},
			w, http.StatusInternalServerError)
		return
	}
	common.ResponseJSON(result, w, resp.StatusCode)
}

func SetUserRoutes(router *mux.Router) {
	router.HandleFunc("/create", CreateUser).Methods("POST")
	router.HandleFunc("/list", ListUsers).Methods("GET")
	router.HandleFunc("/update/{user}", UpdateUser).Methods("PUT")
	router.HandleFunc("/{user}", GetUser).Methods("GET")
}
