package planets

import (
	"github.com/thelmholtz/astroweather/geometry"
)

//Planet represents a planet through it's orbital speed in degrees per day (W) and it's orbital radius in km (R)
type Planet struct {
	W int //Angular velocity in degrees per day, clockwise-positive.
	R int //Radius from the sun in kilometers.
}

//PlanetarySystem returns the current solar system. This should be loaded from a DB if more planets are expected
func PlanetarySystem() []Planet {

	planets := make([]Planet, 3)
	planets[0] = Planet{1, 500}   //Ferengi
	planets[1] = Planet{3, 2000}  //Betasoide
	planets[2] = Planet{-5, 1000} //Vulcano

	return planets
}

//Aligned returns true if the planets are colinear among each other at the given day
func Aligned(day int) bool {

	planets := PlanetarySystem()

	locations := locations(planets, day)

	return locations[0].CollinearTo(locations[1:]...)

}

//RadiallyAligned returns true if the planets are both colinear among both each other and the sun (origin) at the given day.
func RadiallyAligned(day int) bool {

	planets := PlanetarySystem()

	for _, p := range planets {
		for _, other := range planets {
			if p != other && p.locationRadial(day).AxisAngle() != other.locationRadial(day).AxisAngle() {
				return false
			}
		}
	}
	return true
}

//SunInsideTriangle returns true if the sun is inside the triangle formed by the planets
//Panics if length of planets is other than 3; as no handler has been implemented for n-sided polygons.
func SunInsideTriangle(day int) bool {

	planets := PlanetarySystem()

	if len(planets) != 3 {
		panic("Either we have new neighbours and this method should be reimplemented for n-sided polygoons; or something went TERRIBLY wrong with our solar system and we should take a minute of mourning")
	}

	locations := locations(planets, day)

	origin := geometry.Point{X: 0.0, Y: 0.0}

	ret := origin.InsidePolygon(locations...)
	return ret
}

//MaxPerimeter returns true if the day has the highest perimeter for the year
func MaxPerimeter(day int) bool {
	perimeters := make([]float64, 360)
	for d := 0; d < 360; d++ {
		perimeters[d] = getPerimeter(d)
	}
	max := perimeters[0]
	index := 0
	for i, v := range perimeters {
		if v > max {
			max = v
			index = i
		}
	}
	if day != index {
		return false
	}
	return true
}

//location returns a planet location's Polar coordinates. They are assumed to be on position w=0 on day 0.
func (planet Planet) locationRadial(day int) geometry.Polar {
	return geometry.Polar{A: planet.W * day, R: planet.R}
}

//locationCartesian returns a planet location's as a Point (cartesian coordinates).
func (planet Planet) locationCartesian(day int) geometry.Point {
	return planet.locationRadial(day).Point()
}

//locations returns an array withe the location of each planet in planets
func locations(planets []Planet, day int) []geometry.Point {

	locations := make([]geometry.Point, len(planets))
	for i, p := range planets {
		locations[i] = p.locationCartesian(day)
	}

	return locations

}

//getPerimeter returns the triangle's perimeter for a given day
func getPerimeter(day int) float64 {
	locations := locations(PlanetarySystem(), day)
	var perimeter float64
	for i, l := range locations {
		if i+1 < len(locations) {
			perimeter += l.DistanceTo(locations[i+1])
		} else {
			perimeter += l.DistanceTo(locations[0])
		}
	}
	return perimeter
}
