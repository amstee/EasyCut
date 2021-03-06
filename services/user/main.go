package main

import (
	"fmt"
	"os"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/amstee/easy-cut/src/config"
	"github.com/rs/cors"
	"github.com/amstee/easy-cut/services/user/src/handlers"
	"github.com/amstee/easy-cut/src/middlewares"
	"github.com/amstee/easy-cut/src/common"
	"github.com/amstee/easy-cut/src/status"
	"github.com/amstee/easy-cut/src/auth0"
)

func initialize(router *mux.Router, secureRoutes mux.MiddlewareFunc) *negroni.Negroni {
	stat := router.PathPrefix("/status").Subrouter()
	user := router.PathPrefix("/").Subrouter()
	router.Use(secureRoutes)

	n := negroni.New()
	c := cors.New(cors.Options{AllowedOrigins: config.Content.Origins})
	n.Use(negroni.NewLogger())
	n.Use(c)
	status.SetStatusRoutes(stat)
	handlers.SetUserRoutes(user)
	n.UseHandler(router)
	return n
}

func main() {
	if err := config.Content.LoadConfig("API_CLIENT"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config.Content.Display()
	if err := auth0.LoadToken(); err != nil {
		fmt.Println("Auth0 error : ", err)
		os.Exit(1)
	}
	router := mux.NewRouter()
	secureRoutes := middlewares.GetSecurityMiddleware()
	service := initialize(router, secureRoutes)
	common.Run(service)
}
