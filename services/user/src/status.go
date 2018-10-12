package src

import (
	"net/http"
	"github.com/gorilla/mux"
)

func SystemStatus(w http.ResponseWriter, _ *http.Request)  {
	w.Write([]byte("User Service Version : 0.0.1"))
}

func CreateStatusRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/status", SystemStatus)
	return router
}

