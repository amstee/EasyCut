package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CreateSalon(w http.ResponseWriter, r *http.Request) {

}

func GetSalon(w http.ResponseWriter, r *http.Request) {

}

// This route should be able to take filters as url params enabling to find salons near etc ...
func ListSalon(w http.ResponseWriter, r *http.Request) {

}

// for update route you should use a specific data type to perform several actions
// Ex: add barber / remove barber / change info

// This add a tag to a list of tags --> use for adding / removing a barber
//POST test/_doc/1/_update
//{
//	"script" : {
//	"source": "ctx._source.tags.add(params.tag)",
//		"lang": "painless",
//		"params" : {
//		"tag" : "blue"
//		}
//	}
//}
func UpdateSalon(w http.ResponseWriter, r *http.Request) {

}

func DeleteSalon(w http.ResponseWriter, r *http.Request) {

}


func SetSalonRoutes(router *mux.Router) {
	router.HandleFunc("/create", CreateSalon).Methods("POST")
	router.HandleFunc("/get/{salon}", GetSalon).Methods("GET")
	router.HandleFunc("/list", ListSalon).Methods("GET")
	router.HandleFunc("/update/{salon}", UpdateSalon).Methods("PUT")
	router.HandleFunc("/delete/{salon}", DeleteSalon).Methods("DELETE")
}