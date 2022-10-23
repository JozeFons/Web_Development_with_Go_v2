package main

import (
	"log"
	"net/http"
	"github.com/JozeFons/Web_Development_with_Go_v2/controllers"
	"github.com/JozeFons/Web_Development_with_Go_v2/templates"
	"github.com/JozeFons/Web_Development_with_Go_v2/views"
	"github.com/go-chi/chi/v5"
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

// func SelectHandler(w http.ResponseWriter, r *http.Request) {
// 	path := r.URL.Path
// 	switch path {
// 	case "/":
// 		path = "home.gohtml"
// 	case "/contact":
// 		path = "contact.gohtml"
// 	case "/faq":
// 		path = "faq.gohtml"
// 	default:
// 		//wrong page index
// 	}

// 	t, err := views.Parse(path)
// 	if err != nil {
// 		log.Printf("Parsing template: %v", err)
// 		http.Error(w, "There was an error parsing the template!", http.StatusInternalServerError)
// 		return
// 	}
// 	t.Execute(w, nil)
// }

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(server.FS, "home.gohtml", "tailwind-css-styling.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(server.FS,"contact.gohtml", "tailwind-css-styling.gohtml"))))

	r.Get("/faq", controllers.StaticHandler(views.Must(views.ParseFS(server.FS, "faq.gohtml", "tailwind-css-styling.gohtml"))))
	
	users := controllers.Users{}
	users.Templates.New = views.Must(views.ParseFS(server.FS, "signup.gohtml", "tailwind-css-styling.gohtml"))
	r.Get("/signup", users.New)
	r.Post("/users", users.Create)
	
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Page not found", http.StatusNotFound)
	})
	// css := http.FileServer(http.Dir(""))
	// r.Handle("/*", http.StripPrefix("", css))
	log.Printf("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}