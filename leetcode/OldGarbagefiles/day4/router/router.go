package router

import (
	"day4/controller"

	"github.com/gorilla/mux"
)

func InitializeRoutes(c controller.CustomController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customers", c.GetAll).Methods("GET")
	r.HandleFunc("/api/customer/{id}", c.Get).Methods("GET")
	r.HandleFunc("/api/customer", c.Post).Methods("POST")
	r.HandleFunc("/api/customer/{id}", c.Put).Methods("PUT")
	r.HandleFunc("/api/customer/{id}", c.Delete).Methods("DELETE")
	return r
}
