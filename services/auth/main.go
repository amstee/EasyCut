package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"github.com/urfave/negroni"
	"github.com/rs/cors"
	"github.com/auth0/go-jwt-middleware"
	"github.com/amstee/easy-cut/services/auth/src/core"
	"github.com/amstee/easy-cut/services/auth/src/handlers"
)

func initialize(router *mux.Router, checkAuth *jwtmiddleware.JWTMiddleware) (*negroni.Negroni) {
	auth := mux.NewRouter()
	status := router.PathPrefix("/status").Subrouter()
	authentication := router.PathPrefix("/token").Subrouter()
	authenticated := auth.PathPrefix("/secure").Subrouter()

	router.PathPrefix("/secure").Handler(
		negroni.New(negroni.HandlerFunc(checkAuth.HandlerWithNext),
		negroni.Wrap(auth), ))
	n := negroni.New()
	c := cors.New(cors.Options{AllowedOrigins: []string{
		"http://easy-cut.ovh",
		"https://easy-cut.ovh",
		"localhost",
	}})
	authenticated.Use()
	n.Use(negroni.NewLogger())
	n.Use(c)
	handlers.SetStatusRoutes(status)
	handlers.SetAuthenticationRoutes(authentication)
	handlers.SetAuthenticatedRoutes(authenticated)
	n.UseHandler(router)
	err := core.SetJwks(); if err != nil {
		fmt.Println(err)
		fmt.Println("unable to retrieve jwks")
		os.Exit(1)
	}
	return n
}

func run(service *negroni.Negroni) {
	fmt.Println("Starting Auth Service on port 8080")
	service.Run(":8080")
}

func main() {
	router := mux.NewRouter()
	checkAuth, err := core.GetJwtMiddleware(); if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	service := initialize(router, checkAuth)
	run(service)
}
