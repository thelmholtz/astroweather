package tests

import (
	"testing"

	"github.com/thelmholtz/astroweather/planets"
)

//TestIsSunInside does tests
func TestIsSunInside(t *testing.T) {
	if !planets.SunInsideTriangle(566) {
		t.Fatalf("10")
	}
}

func TestRadiallyAligned(t *testing.T) {
	if planets.RadiallyAligned(19) {
		t.Fatalf("19")
	}
}

func TestAligned(t *testing.T) {
	if !planets.Aligned(19) {
		t.Fatalf("19")
	}
}
