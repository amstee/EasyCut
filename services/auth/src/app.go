package src

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"github.com/amstee/easy-cut/services/auth/src/types"
	"github.com/amstee/easy-cut/services/auth/src/utils"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func CheckToken(w http.ResponseWriter, r *http.Request) {
	authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	if len(authHeaderParts) < 2 {
		utils.ResponseJSON(types.HttpMessage{Message: "authorization header not found", Success: false}, w, http.StatusInternalServerError)
		return
	}
	tokenString := authHeaderParts[1]

	token, err := jwt.Parse(tokenString, nil); if err != nil {
		utils.ResponseJSON(types.HttpMessage{Message: "unable to parse jwt", Success: false}, w, http.StatusInternalServerError)
	}
	result, err := CheckTokenValidity(token); if err != nil {
		utils.ResponseJSON(types.HttpMessage{Message: "token is invalid", Success: false}, w, http.StatusInternalServerError)
	}
	fmt.Println(result)
	utils.ResponseJSON(types.HttpMessage{Message: "token is valid", Success: true}, w, 200)
}

func Permissions(w http.ResponseWriter, r *http.Request) {
	var perms types.PermissionsParam
	var resp types.PermissionsResponse
	var isAllowed bool

	authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	token := authHeaderParts[1]

	err := utils.DecodeJSON(&perms, r); if err != nil {
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
	router.HandleFunc("/token", CheckToken)
}

func SetAuthenticatedRoutes(router *mux.Router) {
	router.HandleFunc("/permissions", Permissions)
}
