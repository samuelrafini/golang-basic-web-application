package controllers

import (
	"fmt"
	"github.com/gorilla/schema"
	"golang-basic-web-application/views"
	"net/http"
)

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
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

func (u *Users) Create(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
	//fmt.Fprintln(w, r.PostFormValue("email"))
	//fmt.Fprintln(w, r.PostForm["email"])
	//fmt.Fprintln(w, r.PostFormValue("password"))
	//fmt.Fprintln(w, r.PostForm["password"])
}
