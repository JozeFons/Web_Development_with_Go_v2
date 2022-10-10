package main

import (
	"fmt"
	"net/http"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./home.html")
// }

// func contactHandler(w http.ResponseWriter, r *http.Request) {	
// 	http.ServeFile(w, r, "./contact.html")
// }
type Router struct {}

func (router Router) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		path = "home.html"
	case "/contact":
		path = "contact.html"
	case "/faq":
		path = "faq.html"
	default:
		
	}	
	http.ServeFile(w, r, "./"+path)
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}