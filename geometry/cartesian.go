package geometry

import "github.com/thelmholtz/astroweather/floats"

//Point is a representation of a point in cartesian coordinates (both of type float64)
type Point struct {
	X float64
	Y float64
}

//IsCollinearTo returns true if this point is collinear to all  in the given array of other points
func (p Point) IsCollinearTo(points ...Point) bool {

	//The Ferengi programmers realized a tolerance needs to be agreed upon to allow floating point comparisons.
	//A council of the best Betasoid astronomers and Vulcani metheorologists was summoned, and has agreed that 0.1 is a close enough.
	//tolerance := 0.1
	tolerance := 0.1

	m := make([]float64, len(points))

	for i, pi := range points {
		m[i] = (p.Y - pi.Y) / (p.X - pi.X)
	}

	switch {
	case floats.All(m, floats.IsInfinite):
		return true
	case floats.Any(m, floats.IsInfinite):
		return false
	case floats.AllEqual(m, tolerance):
		return true
	default:
		return false
	}
}
