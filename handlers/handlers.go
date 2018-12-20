/*Package handlers is a package that groups all service handlers. Each handler translates from http to an addecuate entity,
  and then operates upon this entities */
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/thelmholtz/astroweather/weather"
)

//Weather is the handler function for `clima` endpoint.
func Weather(w http.ResponseWriter, r *http.Request) {

	day := weather.Day{D: r.FormValue("dia"), Weather: "undefined"}

	enc := json.NewEncoder(w)

	if err := day.Forecast(); err != nil {
		log.Print(err)
		enc.Encode(err) //err type should be json encodable to allow this; see github.com/thelmholtz/except
		return
	}

	enc.Encode(day)

}
