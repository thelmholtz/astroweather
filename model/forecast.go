package model

import (
	"github.com/thelmholtz/astroweather/planets"
	"github.com/thelmholtz/except"
)

//Forecast represents a day-weather value pair and it's json mapping.
type Forecast struct {
	Day     int    `bson:"day" json:"day"`
	Weather string `bson:"weather" json:"weather"`
	Peak    bool   `bson:"peak" json:"peak,omitempty"`
}

//Predict sets the Weather for this Day with an accurate prediction. Overwrites any previous value.
func (f *Forecast) Predict() except.E {

	switch {
	case f.Day < 0:
		err := except.New("VALUE", "Day must be positive")
		return err
	case isDry(f.Day):
		f.Weather = "drought"
		return nil
	case isOptimal(f.Day):
		f.Weather = "optimal"
		return nil
	case isRaining(f.Day):
		f.Weather = "rain"
		if planets.MaxPerimeter(f.Day) {
			f.Peak = true
		}
		return nil
	default:
		f.Weather = "normal"
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
