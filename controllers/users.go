package controllers

import (
	"fmt"
	"net/http"
	"super_basic_go_app/views"
)

//NewUsers is used to create a new Users Controller
//This function will panic if the templates are not
//parsed correctly, and should only be used during
//initial setup
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.html"),
	}
}

type Users struct {
	NewView *views.View
}

//New is used to render the form where a user
//can create a new user account
//
//GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schmea:"password"`
}

//Create is used to process the signup
//form when a user submits it.
//This is used to create a new user
//account
//
//POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, form)
}
