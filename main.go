package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JozeFons/Web_Development_with_Go_v2/controllers"
	"github.com/JozeFons/Web_Development_with_Go_v2/models"
	server "github.com/JozeFons/Web_Development_with_Go_v2/templates"
	"github.com/JozeFons/Web_Development_with_Go_v2/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
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

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(server.FS, "contact.gohtml", "tailwind-css-styling.gohtml"))))

	r.Get("/faq", controllers.StaticHandler(views.Must(views.ParseFS(server.FS, "faq.gohtml", "tailwind-css-styling.gohtml"))))

	// Setup database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	controllers.CheckError(err)
	defer db.Close()

	err = db.Ping()
	controllers.CheckError(err)
	fmt.Println("Connected!")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL
		);
	`)
	controllers.CheckError(err)
	fmt.Println("Table created!")

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}

	// Setup our controllers
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(server.FS, "signup.gohtml", "tailwind-css-styling.gohtml"))

	usersC.Templates.SignIn = views.Must(views.ParseFS(server.FS, "signin.gohtml", "tailwind-css-styling.gohtml"))
	
	usersC.Templates.PwReset = views.Must(views.ParseFS(server.FS, "password_reset.gohtml", "tailwind-css-styling.gohtml"))
	
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/pw_reset", usersC.PwReset)
	r.Get("/users/id", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Page not found", http.StatusNotFound)
	})
	// css := http.FileServer(http.Dir(""))
	// r.Handle("/*", http.StripPrefix("", css))
	log.Printf("Starting the server on :3000...")
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false),
	)
	err = http.ListenAndServe(":3000", csrfMw(r))
	controllers.CheckError(err)
}
