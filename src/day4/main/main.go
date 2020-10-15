package main

import (
	"day4/controller"
	"day4/mapstore"
	"day4/router"
	"fmt"
	"net/http"
)

func main() {
	h := controller.CustomController{
		Store: mapstore.NewMapStore(),
	}
	r := router.InitializeRoutes(h) // configure routes

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	fmt.Println("Listening...")
	server.ListenAndServe() // Run the http server
}
