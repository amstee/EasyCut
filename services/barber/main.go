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
	"github.com/amstee/easy-cut/src/es"
	"github.com/amstee/easy-cut/services/barber/src/vars"
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
	if err := config.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config.Display()
	initES()
	router := mux.NewRouter()
	secureRoutes := middlewares.GetSecurityMiddleware()
	service := initalize(router, secureRoutes)
	common.Run(service)
}
