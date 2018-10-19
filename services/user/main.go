package main

import (
	"fmt"
	"os"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"strconv"
	"github.com/amstee/easy-cut/src/config"
	"github.com/rs/cors"
	"github.com/amstee/easy-cut/services/user/src/handlers"
	"github.com/amstee/easy-cut/src/middlewares"
	"github.com/amstee/easy-cut/src/auth0"
)

func initialize(router *mux.Router, secureRoutes mux.MiddlewareFunc) *negroni.Negroni {
	status := router.PathPrefix("/status").Subrouter()
	user := router.PathPrefix("/").Subrouter()
	router.Use(secureRoutes)

	n := negroni.New()
	c := cors.New(cors.Options{AllowedOrigins: config.Content.Origins})
	n.Use(negroni.NewLogger())
	n.Use(c)
	handlers.SetStatusRoutes(status)
	handlers.SetUserRoutes(user)
	n.UseHandler(router)
	return n
}

func run(service *negroni.Negroni) {
	fmt.Printf("Starting User Service on port %d\n", config.Content.Port)
	service.Run(":" + strconv.Itoa(config.Content.Port))
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
	service := initialize(router, secureRoutes)
	run(service)
}
