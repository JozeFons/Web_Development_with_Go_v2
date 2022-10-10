package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./home.html")
// }

// func contactHandler(w http.ResponseWriter, r *http.Request) {	
// 	http.ServeFile(w, r, "./contact.html")
// }

func SelectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		path = "home.html"
	case "/contact":
		path = "contact.html"
	case "/faq":
		path = "faq.html"
	default:
		path = "page_not_found.html"
	}
		
	http.ServeFile(w, r, "./"+path)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", SelectHandler)
	r.Get("/contact", SelectHandler)
	r.Get("/faq", SelectHandler)
	r.Get("/page_not_found", SelectHandler)
	s := http.FileServer(http.Dir(""))
	r.Handle("/*", http.StripPrefix("", s))
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}