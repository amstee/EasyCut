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
		common.ResponseError("unable to retrieve token", err, w, http.StatusInternalServerError); return
	}
	token, err := jwt.Parse(tokenString, core.CheckTokenValidity); if err != nil {
		common.ResponseError("unable to pars jwt", err, w, http.StatusInternalServerError); return
	}
	_, err = core.CheckTokenValidity(token); if err != nil {
		common.ResponseError("invalid token", err, w, http.StatusBadRequest); return
	}
	common.ResponseJSON(types.HttpMessage{Message: "token is valid", Success: true}, w, 200)
}

func Groups(w http.ResponseWriter, r *http.Request) {
	var groups vars.GroupsParam
	var resp *vars.GroupsResponse

	token, err := common.GetBearer(r); if err != nil {
		common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
	}
	err = common.DecodeJSON(&groups, r); if err != nil {
		common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
	}
	resp, err = core.CheckGroups(groups.Groups, token); if err != nil {
		common.ResponseError("unable to check user groups", err, w, http.StatusInternalServerError); return
	}
	common.ResponseJSON(resp, w, 200)
}

func ExtractToken(w http.ResponseWriter, r *http.Request) {
	var resp types.ExtractResponse
	var content vars.ExtractContent
	vals := r.URL.Query()

	tokenString, err := common.GetBearer(r); if err != nil {
		common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
	}
	content.Load(vals)
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := core.GetCertificate(token); if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}); if err != nil {
		common.ResponseError("unable to parse token", err, w, http.StatusBadRequest); return
	}
	claims, ok := token.Claims.(*jwt.StandardClaims); if ok && token.Valid {
		if content.User {
			resp.UserId = claims.Subject[6:]
		}
		common.ResponseJSON(resp, w, http.StatusOK); return
	}
	common.ResponseError("unable to parse token", err, w, http.StatusBadRequest); return
}

func GroupsAndMatch(w http.ResponseWriter, r *http.Request) {
	var groups vars.GroupsParam
	var resp *vars.GroupsResponse
	v := mux.Vars(r)

	if userId, ok := v["user"]; ok {
		tokenString, err := common.GetBearer(r); if err != nil {
			common.ResponseError("unable to retrieve token", err, w, http.StatusInternalServerError); return
		}
		claims, err := core.GetPermissionClaims(tokenString); if err != nil {
			common.ResponseError("invalid token", err, w, http.StatusBadRequest); return
		}
		err = common.DecodeJSON(&groups, r); if err != nil {
			common.ResponseError("unable to decode json body", err, w, http.StatusBadRequest); return
		}
		if claims.Subject == (config.Content.TPrefix + userId) {
			resp, err = core.CheckGroups(groups.Groups, tokenString); if err != nil {
				common.ResponseError("unable to check user groups", err, w, http.StatusInternalServerError); return
			}
			common.ResponseJSON(resp, w, 200); return
		}
		common.ResponseError("token doesnt match user", nil, w, http.StatusForbidden); return
	}
	common.ResponseError("user not found", nil, w, http.StatusBadRequest)
}

func Permissions(w http.ResponseWriter, r *http.Request) {
	var perms vars.PermissionsParam
	resp := vars.PermissionsResponse{Scopes: make(map[string]bool)}
	var isAllowed bool

	token, err := common.GetBearer(r); if err != nil {
		common.ResponseError("unable to retrieve token", err, w, http.StatusBadRequest); return
	}
	err = common.DecodeJSON(&perms, r); if err != nil {
		common.ResponseError("unable to decode body", err, w, http.StatusBadRequest); return
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
	router.HandleFunc("/match/{user}", GroupsAndMatch).Methods("POST")
	router.HandleFunc("/groups", Groups).Methods("POST")
	router.HandleFunc("/extract", ExtractToken).Methods("GET")
}