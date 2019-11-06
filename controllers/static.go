package controllers

import "golang-basic-web-application/views"

func NewStatic() *Static {
	return &Static{
		Home:views.NewView("default", "static/contact"),
		Contact:views.NewView("default", "static/home"),
	}
}

type Static struct {
	Home *views.View
	Contact *views.View
}