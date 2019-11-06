package controllers

import "golang-basic-web-application/views"

func NewStatic() *Static {
	return &Static{
		Home:views.NewView("default", "views/static/contact.gohtml"),
		Contact:views.NewView("default", "views/static/home.gohtml"),
	}
}

type Static struct {
	Home *views.View
	Contact *views.View
}