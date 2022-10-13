package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/JozeFons/Web_Development_with_Go_v2/views"
)

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
		//wrong page index
	}
	
	t, err := views.Parse(path)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing the template!", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", SelectHandler)
	r.Get("/contact", SelectHandler)
	r.Get("/faq", SelectHandler)
	css := http.FileServer(http.Dir(""))
	r.Handle("/*", http.StripPrefix("", css))
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}