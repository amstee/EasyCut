package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/amstee/easy-cut/utils"
	"github.com/amstee/easy-cut/utils/types"
)

func SystemStatus(w http.ResponseWriter, _ *http.Request)  {
	resp := types.StatusResponse{Status: "ok", Service: "auth", Version: "0.0.1"}
	utils.ResponseJSON(resp, w, 200)
}

func SetStatusRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("", SystemStatus).Methods("GET")
	return router
}