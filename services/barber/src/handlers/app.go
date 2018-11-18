package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/barber/src/core"
	"github.com/amstee/easy-cut/services/barber/src/vars"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
	"github.com/amstee/easy-cut/src/types"
	"github.com/amstee/easy-cut/src/logger"
	"github.com/amstee/easy-cut/src/auth0"
)

func CreateBarber(w http.ResponseWriter, r *http.Request) {
	role := vars.GetBarberCreation()
	result := vars.BarberResponse{}
	v := mux.Vars(r)

	authToken, err := auth0.GetToken(); if err != nil {
		common.ResponseError("unable to retrieve api token", err, w, http.StatusBadRequest); return
	}
	userId, ok := v["user"]; if ok {
		token, err := common.GetBearer(r); if err == nil {
			err = common.DecodeJSON(&result.Barber, r); if err == nil {
				resp, err := request.ExpectJson(config.GetServiceURL("user") +  "/update/" + userId,
					http.MethodPut, "Bearer " + token, role, &result.User)
				if err == nil && resp.StatusCode == http.StatusOK {
					err = core.CreateBarber(&result.Barber, userId); if err == nil {
						resp, err = request.ExpectJson(config.GetServiceURL("perms") + "/update/" + userId,
							"PUT", authToken.Format(), types.CreateGroup("Barber", false), nil)
						if err == nil && request.IsValid(resp.StatusCode) {
							logger.Info.Println("Created barber: ", userId)
							common.ResponseJSON(result, w, http.StatusCreated); return
						}
						common.ResponseError("failed to add group to barber", err, w, http.StatusInternalServerError); return
					}
					common.ResponseError("unable to save barber's data", err, w, http.StatusInternalServerError); return
				}
				common.ResponseError("unable to create barber role", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
		}
		common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("user not found in url", nil, w, http.StatusBadRequest); return
}

func GetBarber(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	result := vars.BarberResponse{}

	userId, ok := v["user"]; if ok {
		token, err := common.GetBearer(r); if err == nil {
			resp, err := request.ExpectJson(config.GetServiceURL("user") + "/get/" + userId, http.MethodGet,
											"Bearer " + token, nil, &result.User)
			if err == nil && resp.StatusCode == http.StatusOK {
				err = core.FindBarber(&result.Barber, userId); if err == nil {
					common.ResponseJSON(result, w, http.StatusOK); return
				}
				common.ResponseError("unable to find this barber's info", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("unable to retrieve user", err, w, http.StatusInternalServerError); return
		}
		common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("user not found in url", nil, w, http.StatusBadRequest); return
}

func ListBarbers(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	var users []common.User
	result := make(map[string]*vars.BarberResponse)

	v.Set("group", "Barber")
	token, err := common.GetBearer(r); if err == nil {
		fetch, err := request.CreateFetchJson(config.GetServiceURL("user") + "/list", http.MethodGet,
										"Bearer " + token, nil)
		if err == nil {
			fetch.URL.RawQuery = v.Encode()
			resp, err := request.FetchJson(fetch, &users)
			if err == nil && resp.StatusCode == http.StatusOK {
				if len(users) <= 0 {
					common.ResponseError("your query didn't match any barber", nil, w, http.StatusOK); return
				}
				for i := range users {
					result[users[i].UserId[6:]] = &vars.BarberResponse{User: users[i]}
				}
				err = core.FindBarbers(result); if err == nil {
					common.ResponseJSON(core.MtoL(result), w, http.StatusOK); return
				}
				common.ResponseError("barbers not found", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("unable to retrieve users", err, w, http.StatusInternalServerError); return
		}
		common.ResponseError("unable to construct request", err, w, http.StatusInternalServerError); return
	}
	common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
}

func UpdateBarber(w http.ResponseWriter, r *http.Request) {
	var barberUpdate vars.Barber
	v := mux.Vars(r)

	userId, ok := v["user"]; if ok {
		err := common.DecodeJSON(&barberUpdate, r); if err == nil {
			err = core.UpdateBarber(&barberUpdate, userId); if err == nil {
				logger.Info.Println("Updated barber: ", userId)
				common.ResponseJSON(barberUpdate, w, http.StatusOK); return
			}
			common.ResponseError("unable to update barber", err, w, http.StatusInternalServerError); return
		}
		common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("user not found in url", nil, w, http.StatusBadRequest); return
}

func DeleteBarber(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)

	userId, ok := v["user"]; if ok {
		err := core.DeleteBarber(userId); if err == nil {
			logger.Info.Println("Deleted barber: ", userId)
			common.ResponseJSON(types.HttpMessage{Message: "Barber deleted"}, w, http.StatusOK); return
		}
		common.ResponseError("unable to delete barber", err, w, http.StatusInternalServerError); return
	}
	common.ResponseError("user not found in url", nil, w, http.StatusBadRequest); return
}

func SetBarberRoutes(router *mux.Router) {
	router.HandleFunc("/create/{user}", CreateBarber).Methods("POST")
	router.HandleFunc("/get/{user}", GetBarber).Methods("GET")
	router.HandleFunc("/list", ListBarbers).Methods("GET")
	router.HandleFunc("/update/{user}", UpdateBarber).Methods("PUT")
	router.HandleFunc("/delete/{user}", DeleteBarber).Methods("DELETE")
}