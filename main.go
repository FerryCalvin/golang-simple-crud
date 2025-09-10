package main

import (
	"golang/config"
	"golang/controllers/homecontroller"
	"log"
	"net/http"
	"golang/controllers/categorycontroller"
	"golang/controllers/productcontroller"
)

func main() {
	config.ConnectDB()

	// HomePage
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server starting on 8080")
	http.ListenAndServe(":8080", nil)
}