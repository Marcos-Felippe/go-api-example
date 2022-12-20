package main

import (
	"net/http"

	"github.com/projetosgo/exemploapi/application/handlers"
)

func main() {

	http.HandleFunc("/hello", handlers.GetHello)

	http.HandleFunc("/create", handlers.CreateUser)
	http.HandleFunc("/get", handlers.GetUser)
	http.HandleFunc("/getall", handlers.GetAllUsers)
	http.HandleFunc("/update", handlers.UpdateUser)
	http.HandleFunc("/delete", handlers.DeleteUser)

	http.ListenAndServe(":8080", nil)

}
