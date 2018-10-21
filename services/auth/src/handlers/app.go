package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/src/types"
	"github.com/amstee/easy-cut/services/auth/src/vars"
	"github.com/amstee/easy-cut/src/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/amstee/easy-cut/services/auth/src/core"
	"github.com/amstee/easy-cut/services/auth/src/config"
)

func CheckToken(w http.ResponseWriter, r *http.Request) {
	tokenString, err := common.GetBearer(r); if err != nil {
		common.ResponseJSON(types.HttpMessage{Message: err.Error(), Success: false}, w,  http.StatusInternalServerError)
		return
	}

	token, err := jwt.Parse(tokenString, core.CheckTokenValidity); if err != nil {
		common.ResponseJSON(types.HttpMessage{Message: "unable to parse jwt " + err.Error(), Success: false}, w, http.StatusInternalServerError)
		return
	}
	_, err = core.CheckTokenValidity(token); if err != nil {
		common.ResponseJSON(types.HttpMessage{Message: "token is invalid", Success: false}, w, http.StatusInternalServerError)
		return
	}
	common.ResponseJSON(types.HttpMessage{Message: "token is valid", Success: true}, w, 200)
}

func Groups(w http.ResponseWriter, r *http.Request) {
	var groups vars.GroupsParam
	var resp *vars.GroupsResponse

	token, err := common.GetBearer(r); if err != nil {
		common.ResponseJSON(types.HttpMessage{Message: err.Error(), Success: false}, w, http.StatusInternalServerError)
		return
	}
	err = common.DecodeJSON(&groups, r); if err != nil {
		common.ResponseJSON(types.HttpMessage{Message: "unable to decode json body", Success: false}, w, http.StatusInternalServerError)
		return
	}
	resp, err = core.CheckGroups(groups.Groups, token); if err != nil {
		common.ResponseJSON(types.HttpMessage{Message: "unable to check user groups", Success: false}, w, http.StatusInternalServerError)
		return
	}
	common.ResponseJSON(resp, w, 200)
}

func GroupsAndMatch(w http.ResponseWriter, r *http.Request) {
	var groups vars.GroupsParam
	var resp *vars.GroupsResponse
	v := mux.Vars(r)

	if userId, ok := v["user"]; ok {
		tokenString, err := common.GetBearer(r); if err != nil {
			common.ResponseJSON(types.HttpMessage{Message: err.Error(), Success: false}, w,  http.StatusInternalServerError)
			return
		}
		claims, err := core.GetPermissionClaims(tokenString); if err != nil {
			common.ResponseJSON(types.HttpMessage{Message: "invalid token : " + err.Error(), Success: false}, w,  http.StatusBadRequest)
			return
		}
		err = common.DecodeJSON(&groups, r); if err != nil {
			common.ResponseJSON(types.HttpMessage{Message: "unable to decode json body", Success: false}, w, http.StatusInternalServerError)
			return
		}
		if claims.Subject == (config.Content.TPrefix + userId) {
			resp, err = core.CheckGroups(groups.Groups, tokenString); if err != nil {
				common.ResponseJSON(types.HttpMessage{Message: "unable to check user groups", Success: false}, w, http.StatusInternalServerError)
				return
			}
			common.ResponseJSON(resp, w, 200)
			return
		}
		common.ResponseJSON(types.HttpMessage{Message: "Forbidden access", Success: false}, w,  http.StatusForbidden)
		return
	}
	common.ResponseJSON(types.HttpMessage{Message: "user not found in request", Success: false}, w,  http.StatusBadRequest)
}

func Permissions(w http.ResponseWriter, r *http.Request) {
	var perms vars.PermissionsParam
	resp := vars.PermissionsResponse{Scopes: make(map[string]bool)}
	var isAllowed bool

	token, err := common.GetBearer(r); if err != nil {
		common.ResponseJSON(types.HttpMessage{Message: err.Error(), Success: false}, w,  http.StatusInternalServerError)
		return
	}

	err = common.DecodeJSON(&perms, r); if err != nil {
		common.ResponseJSON(types.HttpMessage{Message: "unable to decode json body", Success: false}, w, http.StatusInternalServerError)
		return
	}
	for _, scope := range perms.Scopes {
		isAllowed = core.CheckScope(scope, token)
		resp.Scopes[scope] = isAllowed
	}
	common.ResponseJSON(resp, w, 200)
}

func SetAuthenticationRoutes(router *mux.Router) {
	router.HandleFunc("", CheckToken).Methods("GET")
}

func SetAuthenticatedRoutes(router *mux.Router) {
	router.HandleFunc("/permissions", Permissions).Methods("POST")
	router.HandleFunc("/groups/{user}", GroupsAndMatch).Methods("POST")
	router.HandleFunc("/groups", Groups).Methods("POST")
}