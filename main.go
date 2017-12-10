package main

import (
	"fmt"
	gmux "github.com/gorilla/mux"
	"net/http"
	"super_basic_go_app/views"
)

var (
	homeView    *views.View
	contactView *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.html")
	contactView = views.NewView("bootstrap", "views/contact.html")
	signupView = views.NewView("bootstrap", "views/signup.html")

	mux := gmux.NewRouter()
	mux.HandleFunc("/", home).Methods("GET")
	mux.HandleFunc("/contact", contact).Methods("GET")
	mux.HandleFunc("/signup", signup).Methods("GET")

	fmt.Println(http.ListenAndServe(":3000", mux))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
