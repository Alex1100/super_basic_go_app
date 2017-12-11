package controllers

import (
	"super_basic_go_app/views"
)

func NewStatic() *StaticController {
	return &StaticController{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

type StaticController struct {
	Home    *views.View
	Contact *views.View
}
