package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/perms/src/vars"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/services/perms/src/core"
	"github.com/amstee/easy-cut/src/types"
)

func GetGroups(w http.ResponseWriter, r *http.Request) {
	var resp []vars.Group
	v := mux.Vars(r)

	if userId, ok := v["user"]; ok {
		err := core.GetUserGroups(&resp, userId); if err != nil {
			common.ResponseError("Unable to retrieve groups", err, w, http.StatusInternalServerError); return
		}
		common.ResponseJSON(resp, w, http.StatusOK); return
	}
	common.ResponseError("no user", nil, w, http.StatusBadRequest); return
}

func UpdateGroups(w http.ResponseWriter, r *http.Request) {
	var resp []vars.GroupResponse
	var update types.UpdateGroups
	v := mux.Vars(r)

	if userId, ok := v["user"]; ok {
		err := common.DecodeJSON(&update, r); if err == nil {
			groups, err := core.RetrieveGroups(); if err == nil {
				for _, g := range update.Groups {
					b := false
					for _, group := range groups.Groups {
						if g.Name == group.Name {
							err = core.UpdateGroup(userId, group, g.Delete); if err != nil {
								resp = append(resp, vars.GroupResponse{Name: group.Name, Success: false, Message: err.Error()})
							} else {
								resp = append(resp, vars.GroupResponse{Name: g.Name, Success: true})
							}
							b = true
						}
					}
					if !b {
						resp = append(resp, vars.GroupResponse{Name: g.Name, Success: false, Message: "group not found"})
					}
				}
				common.ResponseJSON(resp, w, http.StatusOK); return
			}
			common.ResponseError("unable to retrieve groups", err, w, http.StatusInternalServerError); return
		}
		common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
	}
	common.ResponseError("no user", nil, w, http.StatusBadRequest); return
}

func SetGroupRoutes(router *mux.Router) {
	router.HandleFunc("/update/{user}", UpdateGroups).Methods("PUT")
	router.HandleFunc("/get/{user}", GetGroups).Methods("GET")
}