package geometry

import (
	"math"
)

//Polar is a representation of a point in polar coordinates (both of type int)
type Polar struct {
	A int
	R int
}

//Point returns the cartesian coordinates (float64) representation of a Polar (int) point
func (p Polar) Point() Point {
	return Point{float64(p.R) * math.Cos((math.Pi*float64(p.A))/180), float64(p.R) * math.Sin((float64(p.A)*math.Pi)/180)}
}

//AxisAngle returns the angle normalized between 0 - 180 degrees
func (p Polar) AxisAngle() int {
	if p.A < 0 {
		return p.A%180 + 180
	}
	return p.A % 180
}
