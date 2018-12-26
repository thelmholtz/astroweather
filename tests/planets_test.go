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
