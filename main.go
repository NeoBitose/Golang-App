package main

import (
	"crud_go_native/config"
	"crud_go_native/controllers/categorycontroller"
	"crud_go_native/controllers/homecontroller"
	"log"
	"net/http"
)

func main() { 
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running on port 8080"); 
	http.ListenAndServe(":8080", nil);
}