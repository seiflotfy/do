package gogo

import "testing"

func TestAll(t *testing.T) {
	v := All([]bool{false, false, false, true})
	if v {
		t.Error("Error expected", false, "got", v)
	}

	v = All([]bool{false, false, false, false})
	if v {
		t.Error("Error expected", false, "got", v)
	}

	v = All([]bool{true, true, true, true})
	if !v {
		t.Error("Error expected", true, "got", v)
	}
}
