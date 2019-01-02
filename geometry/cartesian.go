package geometry

import (
	"math"

	"github.com/thelmholtz/astroweather/floats"
)

//Point is a representation of a point in cartesian coordinates (both of type float64)
type Point struct {
	X float64
	Y float64
}

//DistanceTo returns the euclidean distance to another point using the pytagorean equation
func (p Point) DistanceTo(other Point) float64 {
	return math.Sqrt(math.Pow(p.X-other.Y, 2.0) + math.Pow(p.Y-other.Y, 2.0))
}

//CollinearTo returns true if this point is collinear to all  in the given array of other points
func (p Point) CollinearTo(points ...Point) bool {

	//The Ferengi programmers realized a tolerance needs to be agreed upon to allow floating point comparisons
	//The Betasoid astronomers argued that it's very unlikely the planets will ever be PERFECTLY aligned
	//A council of the best Vulcani metheorologists was summoned, and they agreed that a 0.3 tolerance is the one that yields the most accurate model in relation to the historic record.
	tolerance := 0.3

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

//InsidePolygon recives a series of ordered vertex (clockwise or counter-clockwise) and returns true if Point p is inside the polygon they describe.
//This extension to n sided polygons is based on the method for triangles described in:
//		https://stackoverflow.com/questions/2049582/how-to-determine-if-a-point-is-in-a-2d-triangle
//		https://math.stackexchange.com/questions/51326/determining-if-an-arbitrary-point-lies-inside-a-triangle-defined-by-three-points
//For len(vertices) > 3; each vertex must be ordered, as the same set of points can define multiple polygons depending on the order in which they are joined;
//Luckily; that's not the case for triangles, and that is the intended use case.
func (p Point) InsidePolygon(vertices ...Point) bool {
	if len(vertices) < 3 {
		panic("geometry.Point.InsidePolygon called with less than three vertices")
	}

	offsets := make([]float64, len(vertices))

	for i := range vertices {
		if i+1 < len(vertices) {
			offsets[i] = edgeOffset(p, vertices[i], vertices[i+1])
		} else {
			//Edge joining last point to starting point
			offsets[i] = edgeOffset(p, vertices[i], vertices[0])
		}
	}

	hasNegative := floats.Any(offsets, func(f float64) bool { return f < 0 })
	hasPositive := floats.Any(offsets, func(f float64) bool { return f > 0 })

	return !(hasNegative && hasPositive)
}

//edgeOffset makes an edge from x to y and returns the signed distance from it (one side/half-plane is positive while the other is negative)
//note that the sign will be affected by the order in which the point's are picked ( that is, edgeOffset(p, x, y) = -1 * edgeOffset(p, y, x) )
func edgeOffset(p, x, y Point) float64 {
	return (p.X-y.X)*(x.Y-y.Y) - (x.X-y.X)*(p.Y-y.Y)
}
