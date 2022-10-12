package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./home.html")
// }

// func contactHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./contact.html")
// }

// func SelectHandler(w http.ResponseWriter, r *http.Request) {
// 	path := r.URL.Path
// 	switch path {
// 	case "/":
// 		path = "home.html"
// 	case "/contact":
// 		path = "contact.html"
// 	case "/faq":
// 		path = "faq.html"
// 	default:
// 		path = "page_not_found.html"
// 	}

// 	http.ServeFile(w, r, "./"+path)
// }

func SelectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		path = "home.gohtml"
	case "/contact":
		path = "contact.gohtml"
	case "/faq":
		path = "faq.gohtml"
	default:
		path = "page_not_found.gohtml"
	}
	
	path = filepath.Join("templates", path)
	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing the template!", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w,nil)
	if err != nil{
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error on executing the template!", http.StatusInternalServerError)
		return
	}
}



func main() {
	r := chi.NewRouter()
	r.Get("/", SelectHandler)
	r.Get("/contact", SelectHandler)
	r.Get("/faq", SelectHandler)
	r.Get("/page_not_found", SelectHandler)
	// s := http.FileServer(http.Dir(""))
	// r.Handle("/*", http.StripPrefix("", s))
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}