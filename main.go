package main

import (
	"github.com/gorilla/mux"
	"golang-basic-web-application/controllers"
	"golang-basic-web-application/views"
	"net/http"
)

var (
 homeView *views.View
 contactView *views.View
 //signupView *views.View
)

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	//err := homeView.Render()
	//	//err := homeView.Template.Execute(w, nil)
	//if err != nil {
	//	panic(err)
	//}
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	//err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	////err := homeView.Template.Execute(w, nil)
	//if err != nil {
	//	panic(err)
	//}
	must(contactView.Render(w, nil))
}

//func signup(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//	must(signupView.Render(w, nil))
//
//}


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
	//var err error
	//homeTemplate, err = template.ParseFiles(
	//	"views/home.gohtml",
	//	"views/layouts/footer.gohtml",
	//	)
	//if err != nil {
	//	panic(err)
	//}
	//contactTemplate, err = template.ParseFiles(
	//	"views/contact.gohtml",
	//	"views/layouts/footer.gohtml",
	//	)
	//if err != nil {
	//	panic(err)
	//}
	homeView = views.NewView("default", "views/home.gohtml")
	contactView = views.NewView("default", "views/contact.gohtml")
	//signupView = views.NewView("default", "views/signup.gohtml")
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", usersC.New)
	http.ListenAndServe(":8080", r)
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}