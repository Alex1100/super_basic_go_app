package controllers

import (
	"super_basic_go_app/views"
)

func NewStatic() *StaticController {
	return &StaticController{
		Home:    views.NewView("bootstrap", "views/static/home.html"),
		Contact: views.NewView("bootstrap", "views/static/contact.html"),
	}
}

type StaticController struct {
	Home    *views.View
	Contact *views.View
}
