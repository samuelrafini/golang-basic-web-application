package controllers

import (
	"golang-basic-web-application/views"
	"net/http"
)

type Users struct {
	NewView *views.View
}
// create new users controller
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("default", "views/users/new.gohtml"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request){
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}