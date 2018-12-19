package geometry

import "math"

type Polar struct {
    A   int
    R   int
}

type Point struct {
    X   float64
    Y   float64
}

func AsPoint(p Polar) Point {
   return Point{float64(p.R) * math.Cos(180.0/(math.Pi*float64(p.A))), float64(p.R) * math.Sin(180.0/(math.Pi*float64(p.A)))}
}
