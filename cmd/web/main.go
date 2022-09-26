package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/hamza-starcevic/bookings/driver"
	"github.com/hamza-starcevic/bookings/pkg/config"
	handler "github.com/hamza-starcevic/bookings/pkg/handlers"
	"github.com/hamza-starcevic/bookings/pkg/models"
	render "github.com/hamza-starcevic/bookings/pkg/renderers"
)

const portNumber = ":8080"

// * Getting app config
var app config.AppConfig
var session *scs.SessionManager

// * Main is the main application function
func main() {

	//*What I am going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	//* Setting up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	//! Connect	to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=hstarcevic password=Kenansin@2002")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	defer db.SQL.Close()

	//* Storing session in app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	//* Storing template cache in app config
	app.TemplateCache = tc
	app.UseCache = true

	//* Storing session in app config
	repo := handler.NewRepo(&app, db)
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
