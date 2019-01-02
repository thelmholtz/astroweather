package main

import (
	"encoding/json"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"net/http"
	"sort"
	"strconv"

	"github.com/thelmholtz/astroweather/model"
	"github.com/thelmholtz/except"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/forecast/predict", GenerateForecast)
	http.HandleFunc("/forecast", GetForecast)
	appengine.Main()
}

//Index is a quick description of the service
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
	<!DOCTYPE HTML5>
	<head></head>
	<body>
	<h1>
		Welcome to the FBV interplanetary weather forecast service.
	</h1>
	<div>
		<p>To check the whole forecasts for this year, try issuing a request to <a href=https://astroweather.appspot.com/forecast>/forecast</a>.</p>
		<p>To query a given day, you can play with the query parameter of this link: <a href=https://astroweather.appspot.com/forecast?day=335>/forecast?day=${day}</a>.</p>
		<p>To query them by a given weather, you may use: <a href=https://astroweather.appspot.com/forecast?weather=optimal>/forecast?weather=${weather}</a>.
		Currently known weather values are 'drought', 'optimal', 'rain' and 'normal'.</p>
		<p>If you want to delete all entries and run the simulation again send a POST to /forecast/predict.
		Please don't use that knowledge to DoS this service, three civilizations depend on it.</p>
		<p>Happy forecasting.</p>
	</div>
	</body>
	`))
}

//GetForecast is the handler function for `clima` endpoint.
func GetForecast(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	ctx := appengine.NewContext(r)

	var atomic bool
	var forecasts []model.Forecast

	q := datastore.NewQuery("Forecast")

	day, err := strconv.Atoi(r.FormValue("day"))

	//We change modify the query depending on the present parameters
	if err == nil {
		atomic = true
		q = q.Filter("Day =", normalizeDay(day)).Order("Day").Limit(1)
	} else if weather := r.FormValue("weather"); weather != "" {
		q = q.Filter("Weather =", weather)
	} else {
		q = q.Order("Day")
	}

	log.Infof(ctx, "Getting Forecast")

	if _, err := q.GetAll(ctx, &forecasts); err != nil {

		log.Errorf(ctx, "%v", err)
		w.WriteHeader(http.StatusBadGateway)

		return
	}

	log.Infof(ctx, "Forecasts: %v", forecasts)

	//If the query was by day, we return the atomic result instead of a list of len 1. We also change the queried day to match the one in the DB (as we only store 360 records)
	if atomic {
		forecasts[0].Day = day
		enc.Encode(forecasts[0])

		return
	}

	if forecasts != nil {
		//I couldn't order by day using the gcloud datastore SDK whenever the filter was by weather; so we sort it manually here:
		sort.Slice(forecasts, func(i, j int) bool { return forecasts[i].Day < forecasts[j].Day })
		enc.Encode(forecasts)
	}

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

	//We get the forecast for days 0 to 359, and insert only these, as the weather is periodic at 360 days.
	forecasts := make([]model.Forecast, 360)

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

	log.Debugf(ctx, "Forecasts are: \n%v", forecasts)

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
			return
		}
	}
	log.Infof(ctx, "All records deleted.")

	log.Infof(ctx, "Inserting records")
	for _, f := range forecasts {
		k := datastore.NewIncompleteKey(ctx, "Forecast", nil)
		if k, err = datastore.Put(ctx, k, &f); err != nil {
			log.Errorf(ctx, "Error inserting forecast: %v", err)
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
		return (360 + (i % 360)) % 360 //Go's implementation of % yields negative numbers if the operand is negative; we need to normalize this again after shifting them from the range (-359, 0) to (0, 359)
	}
	return i % 360
}
