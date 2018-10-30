package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/barber/src/core"
	"github.com/amstee/easy-cut/services/barber/src/vars"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
)

func CreateBarber(w http.ResponseWriter, r *http.Request) {
	role := vars.GetBarberCreation()
	result := vars.BarberResponse{}
	v := mux.Vars(r)

	userId, ok := v["user"]; if ok {
		token, err := common.GetBearer(r); if err == nil {
			err = common.DecodeJSON(&result.Barber, r); if err == nil {
				resp, err := request.ExpectJson(config.GetServiceURL("user") +  "/update/" + userId,
					http.MethodPut, "Bearer " + token, role, &result.User)
				if err == nil && resp.StatusCode == http.StatusOK {
					err = core.CreateBarber(&result.Barber, userId); if err == nil {
						common.ResponseJSON(result, w, http.StatusOK); return
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
			resp, err := request.ExpectJson(config.GetServiceURL("user") + userId, http.MethodGet,
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
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func UpdateBarber(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func SetBarberRoutes(router *mux.Router) {
	router.HandleFunc("/create/{user}", CreateBarber).Methods("POST")
	router.HandleFunc("/get/{user}", GetBarber).Methods("GET")
	router.HandleFunc("/list", ListBarbers).Methods("GET")
	router.HandleFunc("/update/{user}", UpdateBarber).Methods("PUT")
}