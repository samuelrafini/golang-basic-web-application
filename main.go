package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template
var contactTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

//func handlerFunc(w http.ResponseWriter, r *http.Request){
//	w.Header().Set("Content-Type", "text/html")
//	if r.URL.Path == "/" {
//		fmt.Fprint(w, "<h1>First server Home</h1>")
//	} else if r.URL.Path == "/contact"{
//		fmt.Fprint(w, "To get in touch please send a mail to <a href=\"mailto:check@check.com\">mail.mail.com</a>")
//	} else {
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprint(w,"<h1>page can not be found</h1>")
//	}
//}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", r)
}