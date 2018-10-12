package main

import (
	"github.com/amstee/easy-cut/services/barber/src"
	"net/http"
	"fmt"
)

func main() {
	src.PrintHello()
	router := src.CreateStatusRouter()
	fmt.Println("Starting Barber Service on port 8080")
	http.ListenAndServe(":8080", router)
}
