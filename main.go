package main

import (
	"fmt"
	gmux "github.com/gorilla/mux"
	"net/http"
	"super_basic_go_app/controllers"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	mux := gmux.NewRouter()
	mux.Handle("/", staticC.Home).Methods("GET")
	mux.Handle("/contact", staticC.Contact).Methods("GET")
	mux.HandleFunc("/signup", usersC.New).Methods("GET")
	mux.HandleFunc("/signup", usersC.Create).Methods("POST")

	fmt.Println(http.ListenAndServe(":3000", mux))
}
