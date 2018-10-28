package main

import (
	"github.com/amstee/easy-cut/src/config"
	"fmt"
	"os"
	"github.com/gorilla/mux"
	"github.com/amstee/easy-cut/src/middlewares"
	"github.com/amstee/easy-cut/src/common"
	"github.com/urfave/negroni"
	"github.com/rs/cors"
	"github.com/amstee/easy-cut/services/barber/src/handlers"
	"github.com/amstee/easy-cut/src/auth0"
)

func initalize(router *mux.Router, secureRoutes mux.MiddlewareFunc) *negroni.Negroni {
	status := router.PathPrefix("/status").Subrouter()
	barber := router.PathPrefix("/").Subrouter()
	router.Use(secureRoutes)

	n := negroni.New()
	c := cors.New(cors.Options{AllowedOrigins: config.Content.Origins})
	n.Use(negroni.NewLogger())
	n.Use(c)
	handlers.SetStatusRoutes(status)
	handlers.SetBarberRoutes(barber)
	n.UseHandler(router)
	return n
}

func main() {
	if err := config.Content.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := auth0.LoadToken(); err != nil {
		fmt.Println("Auth0", err)
		os.Exit(1)
	}
	config.Content.Display()
	router := mux.NewRouter()
	secureRoutes := middlewares.GetSecurityMiddleware()
	service := initalize(router, secureRoutes)
	common.Run(service)
}
