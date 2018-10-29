package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/barber/src/core"
	"github.com/amstee/easy-cut/services/barber/src/vars"
	"github.com/amstee/easy-cut/src/auth0"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
	"github.com/pkg/errors"
)

func CreateBarber(w http.ResponseWriter, r *http.Request) {
	role := vars.GetBarberCreation()
	result := vars.BarberResponse{}
	v := mux.Vars(r)

	userId, ok := v["user"]; if ok {
		token, err := auth0.GetToken(); if err == nil {
			err = common.DecodeJSON(&result.Barber, r); if err == nil {
				resp, err := request.ExpectJson(config.GetApi() + "users/" + config.Content.TPrefix + userId,
					http.MethodPatch, token.Format(), role, &result.User)
				if err == nil && resp.StatusCode == http.StatusOK {
					err = core.CreateBarber(&result.Barber); if err == nil {
						common.ResponseJSON(result, w, http.StatusOK)
						return
					}
					common.ResponseError("unable to save barber data", err, w, http.StatusInternalServerError); return
				}
				common.ResponseError("unable to create barber role", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
		}
		common.ResponseError("unable to retrieve api token", err, w, http.StatusInternalServerError); return
	}
	common.ResponseError("user not found in url", errors.New(""), w, http.StatusInternalServerError); return
}

func GetBarber(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func ListBarbers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func UpdateBarber(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func SetBarberRoutes(router *mux.Router) {
	router.HandleFunc("/create/{user}", CreateBarber).Methods("POST")
	router.HandleFunc("/get/{barber}", GetBarber).Methods("GET")
	router.HandleFunc("/list", ListBarbers).Methods("GET")
	router.HandleFunc("/update/{barber}", UpdateBarber).Methods("PUT")
}