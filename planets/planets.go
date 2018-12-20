package planets

import (
	"math"

	"github.com/thelmholtz/astroweather/geometry"
)

//Planet represents a planet through it's orbital speed in degrees per day (W) and it's orbital radius in km (R)
type Planet struct {
	W int //Angular velocity in degrees per day, clockwise-positive.
	R int //Radius from the sun in kilometers.
}

//Ferengi ..
var Ferengi = Planet{1, 500}

//Betasoide ...
var Betasoide = Planet{3, 2000}

//Vulcano ...
var Vulcano = Planet{-5, 1000}

//Location returns a planet location's Polar coordinates. They are assumed to be on position w=0 on day 0.
func (planet Planet) Location(day int) geometry.Polar {
	return geometry.Polar{A: planet.W * day % 360, R: planet.R}
}

//LocationCartesian returns a planet location's as a Point (cartesian coordinates).
func (planet Planet) LocationCartesian(day int) geometry.Point {
	return planet.Location(day).Point()
}

//RadiallyAligned returns true if the planets are both colinear between themselves and the origin at the given day
func RadiallyAligned(day int) bool {
	//TODO Initialize planets golbally
	planets := make([]Planet, 3)
	planets[0] = Ferengi
	planets[1] = Betasoide
	planets[2] = Vulcano

	for _, p := range planets {
		for _, other := range planets {
			if p != other && p.Location(day).AxisAngle() != other.Location(day).AxisAngle() {
				return false
			}
		}
	}
	return true
}

//Aligned returns true if the planets are colinear among themselves at the given day
func Aligned(day int) bool {

	//TODO Initialize planets properly
	planets := make([]Planet, 3)
	planets[0] = Ferengi
	planets[1] = Betasoide
	planets[2] = Vulcano

	//Some tolerance needs to be agreed upon to allow floating point comparisons.
	//A council of the best Betasoid astronomers and Vulcani metheorologists has agreed upon 0.1 as a close enough value.
	tolerance := 0.1

	p1, p2, p3 := planets[0].LocationCartesian(day), planets[1].LocationCartesian(day), planets[2].LocationCartesian(day)

	m1 := (p1.Y - p2.Y) / (p1.X - p2.X)
	m2 := (p1.Y - p3.Y) / (p1.X - p3.X)

	switch {
	case isInfinite(m1) && isInfinite(m2):
		return true
	case isInfinite(m1) || isInfinite(m2):
		return false
	case m1 == m2:
		return true
	default:
		dispersion := math.Abs(m1 / m2)
		if dispersion > 1.0-tolerance && dispersion < 1.0+tolerance {
			return true
		}
		return false
	}
}

//Helpers
func isInfinite(f float64) bool { return (f == math.Inf(1) || f == math.Inf(-1)) }
