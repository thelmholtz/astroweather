package weather

import (
	"log"
	"strconv"

	"github.com/thelmholtz/astroweather/planets"
	"github.com/thelmholtz/except"
)

//Day represents a day-weather value pair and it's json mapping.
type Day struct {
	D       string `json:"dia"`
	Weather string `json:"clima"`
}

//Forecast sets the Weather for this Day with an accurate prediction. Overwrites any previous value.
func (day *Day) Forecast() except.E {

	log.Print("Queried weather for day: " + day.D)

	key, err := strconv.Atoi(day.D)

	if err != nil {
		log.Print(err)
		return except.New("TYPE", "El dia tiene que ser numerico")
	}

	switch {
	case key < 0:
		err := except.New("VALUE", "El dia tiene que ser positivo")
		return err
	case isRaining(key):
		day.Weather = "lluvia"
		return nil
	case isOptimal(key):
		day.Weather = "optimal"
		return nil
	case isDry(key):
		day.Weather = "sequia"
		return nil
	default:
		day.Weather = "normal"
		return nil
	}
}

//Helpers
func isRaining(day int) bool {
	return false
}

func isDry(day int) bool {
	return planets.RadiallyAligned(day)
}

func isOptimal(day int) bool {
	return planets.Aligned(day)
}
