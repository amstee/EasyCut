package middlewares

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/amstee/easy-cut/src/types"
	"github.com/amstee/easy-cut/src/config"
	"bytes"
	"errors"
	"github.com/amstee/easy-cut/src/common"

	"regexp"
)

type GroupsQuery struct {
	Groups []string				`json:"groups"`
}

type GroupsResponse struct {
	Groups map[string]bool		`json:"groups"`
}

func ResponseError(data interface{}, w http.ResponseWriter, statusCode int) {
	jsonResponse, err := json.Marshal(data); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

func RequestGroups(token string, permissions GroupsQuery, complement string) (*GroupsResponse, error) {
	var result GroupsResponse
	client := &http.Client{}
	route := "/secure/groups"

	jsonData, err := json.Marshal(permissions); if err != nil {
		return nil, err
	}
	if complement != "" {
		route = "/secure/match"
	}
	req, err := http.NewRequest("POST", config.GetServiceURL("security") + route + complement,
								bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + token)
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req); if err != nil || resp.StatusCode != 200 {
		if err == nil {
			err = errors.New("invalid status code received from auth")
		}
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result); if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetSecurityMiddleware() (mux.MiddlewareFunc) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := ""
			route := mux.CurrentRoute(r)
			cr, err := route.GetPathTemplate(); if err != nil {
				common.ResponseError("Invalid request", err, w, http.StatusBadRequest); return
			}
			for _, perm := range config.Content.Routes {
				match, err := regexp.MatchString(perm.Route, cr); if err != nil {
					common.ResponseError("Error in config file", err, w, http.StatusInternalServerError); return
				}
				if match {
					if len(perm.Permissions) != 0 {
						token, err := common.GetBearer(r); if err != nil {
							common.ResponseError("token missing", err, w, http.StatusBadRequest); return
						}
						if perm.Match {
							v := mux.Vars(r)
							tmp, ok := v["user"]; if !ok {
								common.ResponseError("user missing", nil, w, http.StatusBadRequest); return
							}
							user =  "/" + tmp
						}
						groups, err := RequestGroups(token, GroupsQuery{Groups: perm.Permissions}, user)
						if err != nil {
							common.ResponseError("cannot retrieve permissions", err, w, http.StatusInternalServerError); return
						}
						for _, v := range groups.Groups {
							if !v {
								common.ResponseError("insufficient permissions", err, w, http.StatusInternalServerError); return
							}
						}
					}
					next.ServeHTTP(w, r)
					return
				}
			}
			ResponseError(types.HttpMessage{Message: "route " + cr + " not found", Success: false},
							w, http.StatusBadRequest)
		})
	}
}