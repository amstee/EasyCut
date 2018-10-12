package main

import (
	"github.com/amstee/easy-cut/services/rating/src"
	"fmt"
	"net/http"
)

func main() {
	src.PrintHello()
	router := src.CreateStatusRouter()
	fmt.Println("Starting Rating Service on port 8080")
	http.ListenAndServe(":8080", router)
}
