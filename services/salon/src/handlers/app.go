package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/salon/src/vars"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
	"github.com/amstee/easy-cut/src/types"
	"github.com/amstee/easy-cut/services/salon/src/core"
	"github.com/amstee/easy-cut/src/logger"
	"github.com/amstee/easy-cut/src/auth0"
)

func CreateSalon(w http.ResponseWriter, r *http.Request) {
	var extract types.ExtractResponse
	role := types.CreateGroup("Salon", false)
	data := vars.Salon{}

	token, err := common.GetBearer(r); if err == nil {
		err = common.DecodeJSON(&data, r); if err == nil {
			resp, err := request.ExpectJson(config.GetServiceURL("security") + "/secure/extract?user=true",
											http.MethodGet, "Bearer " + token, nil, &extract)
			if err == nil && request.IsValid(resp.StatusCode) {
				token, err := auth0.GetToken(); if err != nil {
					common.ResponseError("failed to retrieve api token", err, w, http.StatusInternalServerError); return
				}
				resp, err := request.ExpectJson(config.GetServiceURL("perms") + "/update/" + extract.UserId,
												http.MethodPut, token.Format(), role, nil)
				if err == nil && request.IsValid(resp.StatusCode) {
					data.UserId = extract.UserId
					err = core.CreateSalon(&data); if err == nil {
						logger.Info.Println("Created salon: ", data.Id, " User:", extract.UserId)
						common.ResponseJSON(data, w, http.StatusCreated); return
					}
					common.ResponseError("failed to save salon", err, w, http.StatusInternalServerError); return
				}
				common.ResponseError("failed to update groups", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("invalid token", err, w, http.StatusBadRequest); return
		}
		common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
}

func GetSalon(w http.ResponseWriter, r *http.Request) {
	data := vars.Salon{}
	v := mux.Vars(r)

	salonId, ok := v["salon"]; if ok {
		err := core.FindSalon(&data, salonId); if err == nil {
			common.ResponseJSON(data, w, http.StatusOK); return
		}
		common.ResponseError("salon not found", err, w, http.StatusNotFound); return
	}
	common.ResponseError("salon not found in url", nil, w, http.StatusBadRequest)
}

func ListSalon(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	var extract vars.ExtractQuery

	if err := extract.Load(v); err != nil {
		common.ResponseError("invalid url parameters", err, w, http.StatusBadRequest); return
	}
	data, err := core.FindSalons(extract); if err == nil {
		common.ResponseJSON(data, w, http.StatusOK); return
	}
	common.ResponseError("Cannot find salons", err, w, http.StatusInternalServerError)
}

func UpdateSalon(w http.ResponseWriter, r *http.Request) {
	data := vars.UpdateSalon{}
	result := vars.Salon{}
	v := mux.Vars(r)

	salonId, ok := v["salon"]; if ok {
		err := common.DecodeJSON(&data, r); if err == nil {
			err = core.UpdateSalon(data, salonId); if err == nil {
				err = core.FindSalon(&result, salonId); if err == nil {
					logger.Info.Println("Updated salon: ", salonId)
					common.ResponseJSON(result, w, http.StatusOK); return
				}
				common.ResponseError("Unable to get salon", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("Unable to update salon", err, w, http.StatusInternalServerError); return
		}
		common.ResponseError("invalid body", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("salon not found in url", nil, w, http.StatusBadRequest)
}

func DeleteSalon(w http.ResponseWriter, r *http.Request) {
	var extract types.ExtractResponse
	data := vars.Salon{}
	v := mux.Vars(r)

	salonId, ok := v["salon"]; if ok {
		token, err := common.GetBearer(r); if err == nil {
			err := core.FindSalon(&data, salonId)
			if err == nil {
				resp, err := request.ExpectJson(config.GetServiceURL("security") + "/secure/extract?user=true",
					http.MethodGet, "Bearer " + token, nil, &extract)
				if err == nil && request.IsValid(resp.StatusCode) {
					if extract.UserId == data.UserId {
						err := core.DeleteSalon(data.Id, false); if err == nil {
							logger.Info.Println("Deleted salon: ", salonId)
							common.ResponseJSON(types.HttpMessage{Message: "Salon deleted"}, w, http.StatusOK); return
						}
						common.ResponseError("unable to delete salon", err, w, http.StatusInternalServerError); return
					}
					common.ResponseError("you can't delete this salon", nil, w, http.StatusUnauthorized); return
				}
				common.ResponseError("invalid token", err, w, http.StatusBadRequest); return
			}
			common.ResponseError("salon not found", err, w, http.StatusNotFound); return
		}
		common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("salon not found in url", nil, w, http.StatusBadRequest)
}


func SetSalonRoutes(router *mux.Router) {
	router.HandleFunc("/create", CreateSalon).Methods("POST")
	router.HandleFunc("/get/{salon}", GetSalon).Methods("GET")
	router.HandleFunc("/list", ListSalon).Methods("GET")
	router.HandleFunc("/update/{salon}", UpdateSalon).Methods("PUT")
	router.HandleFunc("/delete/{salon}", DeleteSalon).Methods("DELETE")
}