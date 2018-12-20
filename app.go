package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thelmholtz/astroweather/handlers"
)

//App is an application container
type App struct {
	Router *mux.Router
}

// Run starts serving all endpoints
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

//Initialize builds the app runtime (router and routes)
func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) initializeRoutes() {
	app.Router.Path("/clima").Queries("dia", "{dia}").HandlerFunc(handlers.Weather)
}
