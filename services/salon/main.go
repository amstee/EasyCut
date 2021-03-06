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
	"github.com/amstee/easy-cut/src/es"
	"github.com/amstee/easy-cut/services/salon/src/handlers"
	"github.com/amstee/easy-cut/services/salon/src/vars"
	"github.com/amstee/easy-cut/src/status"
	"github.com/amstee/easy-cut/src/auth0"
)

func initalize(router *mux.Router, secureRoutes mux.MiddlewareFunc) *negroni.Negroni {
	stat := router.PathPrefix("/status").Subrouter()
	salon := router.PathPrefix("/").Subrouter()
	router.Use(secureRoutes)

	n := negroni.New()
	c := cors.New(cors.Options{AllowedOrigins: config.Content.Origins})
	n.Use(negroni.NewLogger())
	n.Use(c)
	status.SetStatusRoutes(stat)
	handlers.SetSalonRoutes(salon)
	n.UseHandler(router)
	return n
}

func initES() {
	if err := es.InitClient(vars.Register); err != nil {
		fmt.Println("Es : ", err)
		return
	} else {
		res, err := es.GetVersion()
		if err != nil {
			fmt.Println("Es version : ", err)
		}
		fmt.Println("Elastic search version --> ", res)
	}
}

func main() {
	if err := config.Load("API_CLIENT"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config.Display()
	if err := auth0.LoadToken(); err != nil {
		fmt.Println("Auth0 error : ", err)
		os.Exit(1)
	}
	initES()
	router := mux.NewRouter()
	secureRoutes := middlewares.GetSecurityMiddleware()
	service := initalize(router, secureRoutes)
	common.Run(service)
}
