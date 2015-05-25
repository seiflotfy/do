package do

import "testing"

func TestAny(t *testing.T) {
	v := Any([]bool{false, false, false, true})
	if !v {
		t.Error("Error expected", true, "got", v)
	}

	v = Any([]bool{false, false, false, false})
	if v {
		t.Error("Error expected", false, "got", v)
	}
}
