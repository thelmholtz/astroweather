package model

import (
	"log"
	"strconv"

	"github.com/thelmholtz/astroweather/planets"
	"github.com/thelmholtz/except"
)

//Forecast represents a day-weather value pair and it's json mapping.
type Forecast struct {
	D       string `bson:"day" json:"dia"`
	Weather string `bson:"weather" json:"clima"`
}

//Predict sets the Weather for this Day with an accurate prediction. Overwrites any previous value.
func (forecast *Forecast) Predict() except.E {

	key, err := strconv.Atoi(forecast.D)

	if err != nil {
		log.Print(err)
		return except.New("TYPE", "El dia tiene que ser numerico")
	}

	switch {
	case key < 0:
		err := except.New("VALUE", "El dia tiene que ser positivo")
		return err
	case isDry(key):
		forecast.Weather = "sequia"
		return nil
	case isOptimal(key):
		forecast.Weather = "optimo"
		return nil
	case isRaining(key):
		forecast.Weather = "lluvia"
		return nil
	default:
		forecast.Weather = "normal"
		return nil
	}
}

//Helpers
func isRaining(day int) bool {
	return planets.SunInsideTriangle(day)
}

func isDry(day int) bool {
	return planets.RadiallyAligned(day)
}

func isOptimal(day int) bool {
	return planets.Aligned(day)
}
