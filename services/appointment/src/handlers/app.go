package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/appointment/src/vars"
	"github.com/amstee/easy-cut/src/types"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
	"github.com/amstee/easy-cut/services/appointment/src/core"
	"github.com/amstee/easy-cut/src/logger"
)

func Schedule(w http.ResponseWriter, r *http.Request) {
	var appointment vars.Appointment
	var extract types.ExtractResponse

	token, err := common.GetBearer(r); if err == nil {
		err = common.DecodeJSON(&appointment, r); if err == nil {
			resp, err := request.ExpectJson(config.GetServiceURL("security") + "/secure/extract?user=true",
											http.MethodGet, "Bearer " + token, nil, &extract)
			if err == nil && request.IsValid(resp.StatusCode) {
				appointment.Id = ""
				err = appointment.Verify(token, extract.UserId); if err != nil {
					common.ResponseError("invalid request", err, w, http.StatusBadRequest); return
				}
				err = core.CreateAppointment(&appointment); if err == nil {
					logger.Info.Println("Created appointment: ", appointment.Id, " UserId ", appointment.UserId, " BarberId ", appointment.BarberId)
					common.ResponseJSON(appointment, w, http.StatusCreated); return
				}
				common.ResponseError("failed to schedule", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("invalid token", err, w, http.StatusBadRequest); return
		}
		common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
}

func GetAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment vars.Appointment
	v := mux.Vars(r)

	appointmentId, ok := v["appointment"]; if ok {
		err := core.FindAppointment(&appointment, appointmentId); if err == nil {
			common.ResponseJSON(appointment, w, http.StatusOK); return
		}
		common.ResponseError("appointment not found", err, w, http.StatusNotFound); return
	}
	common.ResponseError("appointment not found in url", nil, w, http.StatusBadRequest)
}

func FindAppointments(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	var extUser types.ExtractResponse
	var extract vars.ExtractQuery

	token, err := common.GetBearer(r); if err == nil {
		resp, err := request.ExpectJson(config.GetServiceURL("security")+"/secure/extract?user=true",
			http.MethodGet, "Bearer "+token, nil, &extUser)
			if err == nil && request.IsValid(resp.StatusCode) {
				if err := extract.Load(v, extUser.UserId); err != nil {
					common.ResponseError("invalid url parameters", err, w, http.StatusBadRequest); return
				}
				data, err := core.FindAppointments(extract); if err == nil {
					common.ResponseJSON(data, w, http.StatusOK); return
				}
				common.ResponseError("Cannot find any appointment", err, w, http.StatusInternalServerError); return
			}
		common.ResponseError("Unable to find user", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
}

func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var extract types.ExtractResponse
	var data vars.Appointment

	appointmentId, ok := v["appointment"]; if ok {
		token, err := common.GetBearer(r); if err == nil {
			err := core.FindAppointment(&data, appointmentId); if err == nil {
				resp, err := request.ExpectJson(config.GetServiceURL("security") + "/secure/extract?user=true",
												http.MethodGet, "Bearer " + token, nil, &extract)
				if err == nil && request.IsValid(resp.StatusCode) {
					if extract.UserId == data.UserId || extract.UserId == data.BarberId {
						err := core.DeleteAppointment(appointmentId, false); if err == nil {
							logger.Info.Println("Deleted appointment: ", appointmentId)
							common.ResponseJSON(types.HttpMessage{Message: "appointment deleted", Success: true}, w, http.StatusOK); return
						}
						common.ResponseError("failed to delete appointment", err, w, http.StatusInternalServerError); return
					}
					common.ResponseError("you can't delete this appointment", nil, w, http.StatusUnauthorized); return
				}
				common.ResponseError("invalid token", err, w, http.StatusBadRequest)
			}
			common.ResponseError("appointment not found", err, w, http.StatusBadRequest)
		}
		common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest)
	}
	common.ResponseError("appointment not found in url", nil, w, http.StatusBadRequest)
}

func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	var data vars.UpdateAppointment
	var result vars.Appointment
	v := mux.Vars(r)

	appointmentId, ok := v["appointment"]; if ok {
		err := common.DecodeJSON(&data, r); if err == nil {
			result.Id = ""
			err = core.UpdateAppointment(data, appointmentId); if err == nil {
				err = core.FindAppointment(&result, appointmentId); if err == nil {
					logger.Info.Println("Updated appointment: ", appointmentId)
					common.ResponseJSON(result, w, http.StatusOK); return
				}
				common.ResponseError("Unable to get appointment", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("Unable to update appointment", err, w, http.StatusInternalServerError); return
		}
		common.ResponseError("invalid body", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("appointment not found in url", nil, w, http.StatusBadRequest)
}

func SetAppointmentHandlers(router *mux.Router) {
	router.HandleFunc("/schedule", Schedule).Methods("POST")
	router.HandleFunc("/get/{appointment}", GetAppointment).Methods("GET")
	router.HandleFunc("/list", FindAppointments).Methods("GET")
	router.HandleFunc("/delete/{appointment}", DeleteAppointment).Methods("DELETE")
	router.HandleFunc("/update/{appointment}", UpdateAppointment).Methods("PUT")
}