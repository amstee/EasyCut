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
)

func CreateSalon(w http.ResponseWriter, r *http.Request) {
	var extract types.ExtractResponse
	role := vars.GetSalonRole()
	data := vars.Salon{}

	token, err := common.GetBearer(r); if err == nil {
		err = common.DecodeJSON(&data, r); if err == nil {
			resp, err := request.ExpectJson(config.GetServiceURL("security") + "/secure/extract?user=true",
											http.MethodGet, "Bearer " + token, nil, &extract)
			if err == nil && resp.StatusCode == http.StatusOK {
				resp, err := request.ExpectJson(config.GetServiceURL("user") + "/update/" + extract.UserId,
												http.MethodPut, "Bearer " + token, role, nil)
				if err == nil && resp.StatusCode == http.StatusOK {
					data.UserId = extract.UserId
					err = core.CreateSalon(&data); if err == nil {
						common.ResponseJSON(data, w, http.StatusOK); return
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

}

// This route should be able to take filters as url params enabling to find salons near etc ...
func ListSalon(w http.ResponseWriter, r *http.Request) {

}

// for update route you should use a specific data type to perform several actions
// Ex: add barber / remove barber / change info

// This add a tag to a list of tags --> use for adding / removing a barber
//POST test/_doc/1/_update
//{
//	"script" : {
//	"source": "ctx._source.tags.add(params.tag)",
//		"lang": "painless",
//		"params" : {
//		"tag" : "blue"
//		}
//	}
//}
func UpdateSalon(w http.ResponseWriter, r *http.Request) {

}

func DeleteSalon(w http.ResponseWriter, r *http.Request) {

}


func SetSalonRoutes(router *mux.Router) {
	router.HandleFunc("/create", CreateSalon).Methods("POST")
	router.HandleFunc("/get/{salon}", GetSalon).Methods("GET")
	router.HandleFunc("/list", ListSalon).Methods("GET")
	router.HandleFunc("/update/{salon}", UpdateSalon).Methods("PUT")
	router.HandleFunc("/delete/{salon}", DeleteSalon).Methods("DELETE")
}