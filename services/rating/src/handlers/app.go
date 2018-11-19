package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/rating/src/vars"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/src/request"
	"github.com/amstee/easy-cut/src/config"
	"github.com/amstee/easy-cut/src/types"
	"github.com/amstee/easy-cut/src/logger"
	"github.com/amstee/easy-cut/services/rating/src/core"
)

func Rate(w http.ResponseWriter, r *http.Request) {
	var data vars.Rate
	var rate vars.Rating
	var extract types.ExtractResponse

	token, err := common.GetBearer(r); if err == nil {
		err = common.DecodeJSON(&data, r); if err == nil {
			resp, err := request.ExpectJson(config.GetServiceURL("security") + "/secure/extract?user=true",
											http.MethodGet, "Bearer " + token, nil, &extract)
			if err == nil && request.IsValid(resp.StatusCode) {
				rate = vars.Rating{UserId: extract.UserId,  Comment: data.Comment, Stars: data.Stars,
									Created: data.Created, Updated: data.Updated}
				rate.TargetId, err = data.GetTarget(); if err != nil {
					common.ResponseError("invalid body", err, w, http.StatusBadRequest); return
				}
				rate.TargetType, err = data.GetTargetType(); if err != nil {
					common.ResponseError("invalid body", err, w, http.StatusBadRequest); return
				}
				err = core.CreateRating(&rate); if err == nil {
					logger.Info.Println("Created rating: ", rate.Id, " TargetId ", rate.TargetId, " TargetType ", rate.TargetType)
					common.ResponseJSON(data, w, http.StatusCreated); return
				}
				common.ResponseError("failed to rate", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("invalid token", err, w, http.StatusBadRequest); return
		}
		common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
}

func GetRating(w http.ResponseWriter, r *http.Request) {

}

func FindRatings(w http.ResponseWriter, r *http.Request) {

}

func DeleteRating(w http.ResponseWriter, r *http.Request) {

}

func UpdateRating(w http.ResponseWriter, r *http.Request) {

}

func SetRatingHandlers(router *mux.Router) {
	router.HandleFunc("/rate", Rate).Methods("POST")
	router.HandleFunc("/get/{rating}", GetRating).Methods("GET")
	router.HandleFunc("/list", FindRatings).Methods("GET")
	router.HandleFunc("/delete/{rating}", DeleteRating).Methods("DELETE")
	router.HandleFunc("/update/{rating}", UpdateRating).Methods("PUT")
}