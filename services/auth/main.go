package main

import (
	"github.com/amstee/easy-cut/services/auth/src"
	"net/http"
	"fmt"
)

func main() {
	src.PrintHello()
	router := src.CreateStatusRouter()
	fmt.Println("Starting Auth Service on port 8080")
	http.ListenAndServe(":8080", router)
}
