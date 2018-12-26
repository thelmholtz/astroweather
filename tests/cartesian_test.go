package tests

import (
	"testing"

	"github.com/thelmholtz/astroweather/geometry"
)

func TestIsCollinearTo(t *testing.T) {
	p1, p2, p3, p4 := geometry.Point{X: 1.0, Y: 1.0}, geometry.Point{X: 2.0, Y: 2.0}, geometry.Point{X: -3.5, Y: -3.5}, geometry.Point{X: 95.46352, Y: 95.463521}
	if !p1.CollinearTo(p2, p3, p4) {
		t.Fatalf("12")
	}
	p5, p6, p7, p8 := geometry.Point{X: -1.0, Y: 1.0}, geometry.Point{X: -2.0, Y: 2.0}, geometry.Point{X: 3.5, Y: -3.5}, geometry.Point{X: -95.46352, Y: 95.463521}
	if !p5.CollinearTo(p6, p7, p8) {
		t.Fatalf("16")
	}
	if p1.CollinearTo(p2, p3, p4, p5) {
		t.Fatalf("19")
	}
}

func TestInsidePolygon(t *testing.T) {
	origin := geometry.Point{X: 0.0, Y: 0.0}
	antiOrigin := geometry.Point{X: -0.0, Y: -0.0}

	p1, p2, p3 := geometry.Point{X: 1.0, Y: 0.0}, geometry.Point{X: 0.0, Y: 1.0}, geometry.Point{X: -1.0, Y: -1.0}

	if !origin.InsidePolygon(p1, p2, p3) {
		t.Fatalf("30")
	}
	if !antiOrigin.InsidePolygon(p1, p2, p3) {
		t.Fatalf("33")
	}
	if !origin.InsidePolygon(p3, p1, p2) {
		t.Fatalf("36, not commutative")
	}
	p4, p5, p6 := geometry.Point{X: 0.0, Y: 1.0}, geometry.Point{X: -2.0, Y: -2.0}, geometry.Point{X: -3.0, Y: -2.0}
	if origin.InsidePolygon(p4, p5, p6) {
		t.Fatalf("40")
	}
	v, b, f := geometry.Point{X: -449.397, Y: -219.1855}, geometry.Point{X: -415.823, Y: -1956.295}, geometry.Point{X: 642.788, Y: 766.040}

	if !origin.InsidePolygon(f, b, v) {
		t.Fatalf("45")
	}
	if !origin.InsidePolygon(b, f, v) {
		t.Fatalf("48")
	}
	if !origin.InsidePolygon(v, b, f) {
		t.Fatalf("51")
	}
	if !origin.InsidePolygon(f, v, b) {
		t.Fatalf("54")
	}
	if !origin.InsidePolygon(b, v, f) {
		t.Fatalf("57")
	}
}
