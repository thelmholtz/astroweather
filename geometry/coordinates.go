package geometry

import "math"

//Polar is a representation of a point in polar coordinates (both of type int)
type Polar struct {
	A int
	R int
}

//Point is a representation of a point in cartesian coordinates (both of type float64)
type Point struct {
	X float64
	Y float64
}

//Point returns the cartesian coordinates (float64) representation of a Polar (int) point
func (p Polar) Point() Point {
	return Point{float64(p.R) * math.Cos(180.0/(math.Pi*float64(p.A))), float64(p.R) * math.Sin(180.0/(math.Pi*float64(p.A)))}
}

//AxisAngle returns the angle normalized between 0 - 180 degrees
func (p Polar) AxisAngle() int {
	if p.A < 0 {
		return -p.A % 180
	}
	return p.A % 180
}
