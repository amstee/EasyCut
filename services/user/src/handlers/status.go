package handlers

import (
	"net/http"
	"github.com/amstee/easy-cut/utils/types"
	"github.com/amstee/easy-cut/utils"
	"github.com/gorilla/mux"
)

func SystemStatus(w http.ResponseWriter, _ *http.Request)  {
	resp := types.StatusResponse{Status: "ok", Service: "user", Version: "0.0.1"}
	utils.ResponseJSON(resp, w, 200)
}

func SetStatusRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("", SystemStatus).Methods("GET")
	return router
}