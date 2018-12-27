package main

import (
	"encoding/json"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"net/http"
	"strconv"

	"github.com/thelmholtz/astroweather/model"
	"github.com/thelmholtz/except"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/predict", GenerateForecast)
	http.HandleFunc("/clima", GetForecast)
	appengine.Main()
}

//Index is a quick description of the service
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the home screen. Try issuing a GET to /predict to generate predictions, or /clima?dia={int} to get the prediction for a given day"))
}

//GetForecast is the handler function for `clima` endpoint.
func GetForecast(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	ctx := appengine.NewContext(r)

	day, err := strconv.Atoi(r.FormValue("dia"))
	if err != nil {
		log.Errorf(ctx, "Day was not castable to int: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		enc.Encode(except.New("TYPE", "Day must be a number"))
	}

	q := datastore.NewQuery("Forecast").Filter("Day =", normalizeDay(day)).Order("Day").Limit(1)

	var forecasts []model.Forecast

	log.Infof(ctx, "Getting Forecast")

	if _, err := q.GetAll(ctx, &forecasts); err != nil {

		//If the query fails, either because there's no records or no infrastructure, we fallback to calculating each day per request.
		f := new(model.Forecast)
		f.Day = day

		log.Errorf(ctx, "Query failed for day: %v, Error: %v", f.Day, err)

		if err := forecasts[0].Predict(); err != nil {
			log.Errorf(ctx, "Error making fallback predictions: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			enc.Encode(err)
			return
		}

		forecasts := make([]model.Forecast, 1)
		forecasts[0] = *f
	}
	log.Infof(ctx, "Forecasts: %v", forecasts)

	forecasts[0].Day = day

	enc.Encode(forecasts[0])
}

//GenerateForecast returns a list of all the days and their prediction
func GenerateForecast(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	enc := json.NewEncoder(w)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		enc.Encode(except.New("VERB", "Method not allowed"))
		return
	}

	forecasts := make([]model.Forecast, 360) //Simula el clima para los proximos 10 años Ferengienses (duran 360 dias)

	log.Infof(ctx, "Generating predictions")
	for i := range forecasts {
		f := model.Forecast{Day: i, Weather: "normal"}
		if err := f.Predict(); err != nil {
			log.Errorf(ctx, "Error generating predictions: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			enc.Encode(err)
			return
		}

		forecasts[i] = f
	}

	var currentEntries []model.Forecast

	log.Infof(ctx, "Querying current entries")
	keys, err := datastore.NewQuery("Forecast").GetAll(ctx, &currentEntries)
	if err != nil {
		log.Errorf(ctx, "Error querying the current dataset: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Infof(ctx, "Deleting the current entries")
	for _, k := range keys {
		if err := datastore.Delete(ctx, k); err != nil {
			log.Errorf(ctx, "Error deleting the entry: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	log.Infof(ctx, "Inserting records")
	for _, f := range forecasts {
		k := datastore.NewIncompleteKey(ctx, "Forecast", nil)
		if k, err = datastore.Put(ctx, k, &f); err != nil {
			log.Errorf(ctx, "Inserting forecast: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Infof(ctx, "%v records inserted correctly: %v", len(forecasts), forecasts)
	enc.Encode(forecasts)
}

//normalizeDays takes a day from minint to maxint and returns it's absolute modulo 360 (0 to 359); which is the amount of days in a Ferengian year.
//As the weather is yearly-periodical for the whole planetary system, normalization allows us to store a small sample of data and yet predict the weather for a theoretically infinite range.
//For instance, negative values can be queried with consistent results.
func normalizeDay(i int) int {
	for i < 0 {
		return (360 + (i % 360)) % 360 //Go's implementation of modulo yields negative numbers if the operand is negative; we need to normalize this again after shifting them from the range (-359, 0) to (0, 359)
	}
	return i % 360
}
