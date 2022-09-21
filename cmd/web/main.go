package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/hamza-starcevic/bookings/pkg/config"
	handler "github.com/hamza-starcevic/bookings/pkg/handlers"
	render "github.com/hamza-starcevic/bookings/pkg/renderers"
)

const portNumber = ":8080"

// * Getting app config
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	//* Setting up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session
	//* Storing session in app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	//* Storing template cache in app config
	app.TemplateCache = tc
	app.UseCache = true

	//* Storing session in app config
	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
