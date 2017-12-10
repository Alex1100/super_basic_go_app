package controllers

import (
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

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
