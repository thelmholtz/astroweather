package geometry

import (
	"testing"
)

func TestAxisAngle(t *testing.T) {

	p1, p2, p3, p4 := Polar{R: 1, A: 45}, Polar{R: 1, A: 135}, Polar{R: 1, A: 225}, Polar{R: 1, A: 315}
	p5, p6, p7, p8 := Polar{R: 1, A: -45}, Polar{R: 1, A: -135}, Polar{R: 1, A: 405}, Polar{R: 1, A: -405}

	if p1.AxisAngle() != 45 {
		t.Fatalf("15")
	}
	if p2.AxisAngle() != 135 {
		t.Fatalf("18")
	}
	if p3.AxisAngle() != 45 {
		t.Fatalf("21")
	}
	if p4.AxisAngle() != 135 {
		t.Fatalf("24")
	}
	if p5.AxisAngle() != 135 {
		t.Fatalf("27")
	}
	if p6.AxisAngle() != 45 {
		t.Fatalf("30")
	}
	if p7.AxisAngle() != 45 {
		t.Fatalf("33")
	}
	if p8.AxisAngle() != 135 {
		t.Fatalf("36")
	}
}
