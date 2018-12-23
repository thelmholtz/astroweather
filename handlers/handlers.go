/*Package handlers is a package that groups all service handlers. Each handler translates from http to an addecuate entity,
  and then operates upon this entities */
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/thelmholtz/astroweather/weather"
)

//Forecast is the handler function for `clima` endpoint.
func Forecast(w http.ResponseWriter, r *http.Request) {

	forecast := weather.Forecast{D: r.FormValue("dia"), Weather: "normal"}

	log.Print("Queried weather for day: " + forecast.D)

	enc := json.NewEncoder(w)

	if err := forecast.Predict(); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		enc.Encode(err) //err type should be json encodable to allow this; see github.com/thelmholtz/except
		return
	}

	enc.Encode(forecast)

}

//Predict returns a list of all the days and their prediction
func Predict(w http.ResponseWriter, r *http.Request) {

	forecasts := make([]weather.Forecast, 360)

	enc := json.NewEncoder(w)

	for i := range forecasts {
		f := weather.Forecast{D: strconv.Itoa(i), Weather: "normal"}
		if err := f.Predict(); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			enc.Encode(err)
			return
		}

		forecasts[i] = f
	}

	enc.Encode(forecasts)
}
