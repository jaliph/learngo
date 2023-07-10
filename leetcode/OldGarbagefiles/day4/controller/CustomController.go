package controller

import (
	"day4/domain"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomController struct {
	Store domain.CustomerStore
}

func (c CustomController) Post(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		fmt.Println("Error Ocurred while parsing customer json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create note
	if err := c.Store.Create(customer); err != nil {
		fmt.Println("Error Ocurred while storing customer json", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("New Customer has been created")
	w.WriteHeader(http.StatusCreated)
}

func (c CustomController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// delete
	if err := c.Store.Delete(id); err != nil {
		fmt.Println("Error while deleting", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Customer Deleted!")
	w.WriteHeader(http.StatusNoContent)
}

func (c CustomController) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var customer domain.Customer
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		fmt.Println("Error ocurred while parsing input object", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update
	if err := c.Store.Update(id, customer); err != nil {
		fmt.Println("Error in update", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Customer has been updated")
	w.WriteHeader(http.StatusNoContent)
}

func (c CustomController) GetAll(w http.ResponseWriter, r *http.Request) {
	if customers, err := c.Store.GetAll(); err != nil {
		fmt.Println("Error ocurred at Get All", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("All Customers are ", customers)
		j, err := json.Marshal(customers)
		if err != nil {
			fmt.Println("Error ocurred while parsing customers")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func (c CustomController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Get by id
	if customer, err := c.Store.GetByID(id); err != nil {
		fmt.Println("Error ocurred at get by id", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Customers found by Id", customer)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(customer)
		if err != nil {
			fmt.Println("Error ocurred while parsing customers")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
