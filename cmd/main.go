package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/hussein-samander/Booking-App/config"
	"github.com/hussein-samander/Booking-App/handlers"
	"github.com/hussein-samander/Booking-App/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	//Not directly using app.TemplateCache as the first variable because it has already been declared while
	//err has not, meaning one will have to be turned into an underscore
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("Cannot create template cache %v", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewConfig(&app)
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	//The "Repo" in handler.Repo.Home is the public variable declared in handlers, not the smallcase "repo"
	//declared within main. The reason we did not declare the capitalcase Repo in main is because
	//we do not want packages importing from each other to avoid an import cycle error
	fmt.Printf("Starting app on port 8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error servin' %v", err)
	}
}
