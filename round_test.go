package gogo

import (
	"testing"
)

func TestRound(t *testing.T) {
	a := 5.08
	b := 1.0001
	c := 7.8
	d := 8.5645

	a = Round(a, 1)
	b = Round(b, 1)
	c = Round(c, 0)
	d = Round(d, 2)

	if a != 5.1 {
		t.Error("Error expected", 5.1, "got", a)
	}

	if b != 1.0 {
		t.Error("Error expected", 1.0, "got", b)
	}

	if c != 8.0 {
		t.Error("Error expected", 8.0, "got", c)
	}

	if d != 8.56 {
		t.Error("Error expected", 8.56, "got", d)
	}
}
