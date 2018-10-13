package main

import (
	"github.com/amstee/easy-cut/services/auth/src"
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"github.com/urfave/negroni"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	wrapper := http.NewServeMux()
	wrapper.Handle("/", router)
	status := router.PathPrefix("/status").Subrouter()
	authenticated := router.PathPrefix("/secure").Subrouter()
	authentication := router.PathPrefix("/auth").Subrouter()
	n := negroni.New()
	c := cors.New(cors.Options{AllowedOrigins: []string{
		"http://easy-cut.ovh",
		"https://easy-cut.ovh",
		"localhost",
	}})
	checkAuth, err := src.GetJwtMiddleware()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wrapper.Handle("/secure", negroni.New(negroni.HandlerFunc(checkAuth.HandlerWithNext)))
	n.Use(negroni.NewLogger())
	n.Use(c)

	src.SetStatusRoutes(status)
	src.SetAuthenticationRoutes(authentication)
	src.SetAuthenticatedRoutes(authenticated)

	fmt.Println("Starting Auth Service on port 8080")
	n.UseHandler(wrapper)
	n.Run(":8080")
}
