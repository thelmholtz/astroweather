package planets

import (
    "github.com/thelmholtz/astroweather/geometry"
    "math"
)

type Planet struct {
    W       int     //Angular velocity in degrees per day, clockwise-positive.
    R       int     //Radius from the sun in kilometers.
}

func (planet Planet) Location(day int) geometry.Polar {
    return geometry.Polar{planet.W * day % 360, planet.R}
}

func (planet Planet) LocationCartesian(day int) geometry.Point {
    return geometry.AsPoint(planet.Location(day))
}

func (planet Planet) AxisAngle(day int) int {
    if planet.Location(day).A < 0 {
        return -planet.Location(day).A % 180
    }
    return planet.Location(day).A % 180
}


//This wouldn't be posible if Ferengi's speed was not 1 (or at least a common integer divisor between 3 and -5)
var Ferengi     Planet      =  Planet{1,500}
var Betasoide   Planet      =  Planet{3,2000}
var Vulcano     Planet      =  Planet{-5, 1000}

func RadiallyAligned(day int) bool {
    //TODO Initialize planets golbally
    planets := make([]Planet, 3)
    planets[0] = Ferengi; planets[1] = Betasoide; planets[2] = Vulcano

    for _, p := range planets {
        for _, other := range planets {
            if p != other && p.AxisAngle(day) != other.AxisAngle(day)  {
                return false
            }
        }
    }
    return true
}

func isInfinite(f float64) bool { return (f == math.Inf(1) || f == math.Inf(-1)) }

func Aligned(day int) bool {
    //TODO Initialize planets properly
    planets := make([]Planet, 3)
    planets[0] = Ferengi; planets[1] = Betasoide; planets[2] = Vulcano

    //TODO Document

    tolerance := 0.05

    p1, p2, p3 := planets[0].LocationCartesian(day), planets[1].LocationCartesian(day), planets[2].LocationCartesian(day)

    m1 := (p1.Y - p2.Y)/(p1.X - p2.X)
    m2 := (p1.Y - p3.Y)/(p1.X - p3.X)

    switch {
    case isInfinite(m1) && isInfinite(m2):
        return true
    case isInfinite(m1) || isInfinite(m2):
        return false
    case m1 == m2:
        return true
    default:
        dispersion := math.Abs(m1/m2)
        if dispersion > 1.0 - tolerance && dispersion < 1.0 + tolerance {
            return true
        }
        return false
    }
}
