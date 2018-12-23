package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/gorilla/mux"

	"github.com/thelmholtz/astroweather/model"
)

//App is an application container
type App struct {
	Router   *mux.Router
	DialInfo *mgo.DialInfo
}

// Run starts serving all endpoints
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

//Initialize builds the app runtime (router and routes)
func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeDB()
	app.initializeRoutes()
}

func (app *App) initializeRoutes() {
	app.Router.Path("/predict").Methods("GET", "POST").HandlerFunc(app.GenerateForecast) //TODO Should only be POST but GET allows for easy testing with the browser
	app.Router.Path("/clima").Queries("dia", "{dia}").HandlerFunc(app.GetForecast)
}

func (app *App) initializeDB() {
	connectionString, present := os.LookupEnv("MONGODB_URL")
	if !present {
		log.Fatal("No connection string specified for db")
	}

	dialInfo, err := mgo.ParseURL(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.DialInfo = dialInfo
}

//GetForecast is the handler function for `clima` endpoint.
func (app *App) GetForecast(w http.ResponseWriter, r *http.Request) {

	day := r.FormValue("dia")

	enc := json.NewEncoder(w)

	session, err := mgo.DialWithInfo(app.DialInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	c := session.DB("test").C("inventory")

	var forecast = new(model.Forecast)
	log.Print("Querying day: ", day)
	if err := c.Find(bson.M{"day": day}).One(forecast); err != nil {
		log.Fatal(err)
	}
	/*
		forecast := model.Forecast{D: r.FormValue("dia"), Weather: "normal"}

		log.Print("Queried weather for day: " + forecast.D)

		if err := forecast.Predict(); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			enc.Encode(err) //err type should be json encodable to allow this; see github.com/thelmholtz/except
			return
		}
	*/
	enc.Encode(forecast)
}

//GenerateForecast returns a list of all the days and their prediction
func (app *App) GenerateForecast(w http.ResponseWriter, r *http.Request) {

	forecasts := make([]model.Forecast, 360*10) //Simula el clima para los proximos 10 años Ferengienses (duran 360 dias)

	enc := json.NewEncoder(w)

	log.Print("Generating predictions")
	for i := range forecasts {
		f := model.Forecast{D: strconv.Itoa(i), Weather: "normal"}
		if err := f.Predict(); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			enc.Encode(err)
			return
		}

		forecasts[i] = f
	}

	session, err := mgo.DialWithInfo(app.DialInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	c := session.DB("test").C("inventory")

	log.Print("Removing all records")
	c.RemoveAll("{}")
	log.Print("Inserting records")
	for _, f := range forecasts {
		if err := c.Insert(f); err != nil {
			log.Fatal(err)
		}
	}
	log.Print("Records inserted correctly")
	enc.Encode(forecasts)
}
