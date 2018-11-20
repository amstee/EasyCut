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
					common.ResponseJSON(rate, w, http.StatusCreated); return
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
	var rate vars.Rating
	v := mux.Vars(r)

	ratingId, ok := v["rating"]; if ok {
		err := core.FindRating(&rate, ratingId); if err == nil {
			common.ResponseJSON(rate, w, http.StatusOK); return
		}
		common.ResponseError("rating not found", err, w, http.StatusNotFound); return
	}
	common.ResponseError("rating not found in url", nil, w, http.StatusBadRequest)
}

func FindRatings(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	var extract vars.ExtractQuery

	if err := extract.Load(v); err != nil {
		common.ResponseError("invalid url parameters", err, w, http.StatusBadRequest); return
	}
	data, err := core.FindRatings(extract); if err == nil {
		common.ResponseJSON(data, w, http.StatusOK); return
	}
	common.ResponseError("Cannot find ratings", err, w, http.StatusInternalServerError)
}

func DeleteRating(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var extract types.ExtractResponse
	var data vars.Rating

	ratingId, ok := v["rating"]; if ok {
		token, err := common.GetBearer(r); if err == nil {
			err := core.FindRating(&data, ratingId); if err == nil {
				resp, err := request.ExpectJson(config.GetServiceURL("security") + "/secure/extract?user=true",
												http.MethodGet, "Bearer " + token, nil, &extract)
				if err == nil && request.IsValid(resp.StatusCode) {
					if extract.UserId == data.UserId {
						err := core.DeleteRating(ratingId, false); if err == nil {
							logger.Info.Println("Deleted rating: ", ratingId)
							common.ResponseJSON(types.HttpMessage{Message: "Rating deleted", Success: true}, w, http.StatusOK); return
						}
						common.ResponseError("failed to delete rating", err, w, http.StatusInternalServerError); return
					}
					common.ResponseError("you can't delete this rating", nil, w, http.StatusUnauthorized); return
				}
				common.ResponseError("invalid token", err, w, http.StatusBadRequest)
			}
			common.ResponseError("rating not found", err, w, http.StatusBadRequest)
		}
		common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest)
	}
	common.ResponseError("rating not found in url", nil, w, http.StatusBadRequest)
}

func UpdateRating(w http.ResponseWriter, r *http.Request) {
	var data vars.UpdateRating
	var result vars.Rating
	v := mux.Vars(r)

	ratingId, ok := v["rating"]; if ok {
		err := common.DecodeJSON(&data, r); if err == nil {
			err = core.UpdateRating(data, ratingId); if err == nil {
				err = core.FindRating(&result, ratingId); if err == nil {
					logger.Info.Println("Updated rating: ", ratingId)
					common.ResponseJSON(result, w, http.StatusOK); return
				}
				common.ResponseError("Unable to get rating", err, w, http.StatusInternalServerError); return
			}
			common.ResponseError("Unable to update rating", err, w, http.StatusInternalServerError); return
		}
		common.ResponseError("invalid body", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("rating not found in url", nil, w, http.StatusBadRequest)
}

func SetRatingHandlers(router *mux.Router) {
	router.HandleFunc("/rate", Rate).Methods("POST")
	router.HandleFunc("/get/{rating}", GetRating).Methods("GET")
	router.HandleFunc("/list", FindRatings).Methods("GET")
	router.HandleFunc("/delete/{rating}", DeleteRating).Methods("DELETE")
	router.HandleFunc("/update/{rating}", UpdateRating).Methods("PUT")
}