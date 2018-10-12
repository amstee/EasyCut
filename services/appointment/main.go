package main

import (
	"github.com/amstee/easy-cut/services/appointment/src"
	"net/http"
	"fmt"
)

func main() {
	src.PrintHello()
	router := src.CreateStatusRouter()
	fmt.Println("Starting Appointment Service on port 8080")
	http.ListenAndServe(":8080", router)
}
