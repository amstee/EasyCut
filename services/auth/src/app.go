package src

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/amstee/easy-cut/services/auth/src/types"
	"github.com/amstee/easy-cut/services/auth/src/utils"
	"github.com/dgrijalva/jwt-go"
)

func CheckToken(w http.ResponseWriter, r *http.Request) {
	tokenString, err := utils.GetBearer(r); if err != nil {
		utils.ResponseJSON(types.HttpMessage{Message: err.Error(), Success: false}, w,  http.StatusInternalServerError)
		return
	}

	token, err := jwt.Parse(tokenString, CheckTokenValidity); if err != nil {
		utils.ResponseJSON(types.HttpMessage{Message: "unable to parse jwt", Success: false}, w, http.StatusInternalServerError)
		return
	}
	_, err = CheckTokenValidity(token); if err != nil {
		utils.ResponseJSON(types.HttpMessage{Message: "token is invalid", Success: false}, w, http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(types.HttpMessage{Message: "token is valid", Success: true}, w, 200)
}

func Permissions(w http.ResponseWriter, r *http.Request) {
	var perms types.PermissionsParam
	resp := types.PermissionsResponse{Scopes: make(map[string]bool)}
	var isAllowed bool

	token, err := utils.GetBearer(r); if err != nil {
		utils.ResponseJSON(types.HttpMessage{Message: err.Error(), Success: false}, w,  http.StatusInternalServerError)
		return
	}

	err = utils.DecodeJSON(&perms, r); if err != nil {
		utils.ResponseJSON(types.HttpMessage{Message: "unable to decode json body", Success: false}, w, http.StatusInternalServerError)
		return
	}
	for _, scope := range perms.Scopes {
		isAllowed = CheckScope(scope, token)
		resp.Scopes[scope] = isAllowed
	}
	utils.ResponseJSON(resp, w, 200)
}

func SetAuthenticationRoutes(router *mux.Router) {
	router.HandleFunc("/token", CheckToken).Methods("GET")
}

func SetAuthenticatedRoutes(router *mux.Router) {
	router.HandleFunc("/permissions", Permissions).Methods("POST")
}