package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CreateBarber(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func SetBarberRoutes(router *mux.Router) {
	router.HandleFunc("/create", CreateBarber).Methods("POST")
}