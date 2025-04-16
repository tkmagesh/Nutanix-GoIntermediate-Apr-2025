package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Model
type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float64 `json:"cost"`
	Category string  `json:"category"`
}

// Inmemory data
var products []Product = []Product{
	{Id: 100, Name: "Pen", Cost: 10, Category: "stationary"},
	{Id: 101, Name: "Pencil", Cost: 5, Category: "stationary"},
	{Id: 102, Name: "Marker", Cost: 50, Category: "stationary"},
}

type AppServer struct {
}

// http.Handler interface implmentation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Welcome to the world of APIs")
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "error encoding data", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "request payload error", http.StatusBadRequest)
				return
			}
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
		}

	case "/customer":
		fmt.Fprintln(w, "List of customers")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}
}

func main() {
	appServer := &AppServer{}
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Println("Error starting server :", err)
	}
}
