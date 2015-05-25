package do

import "testing"

func TestRange(t *testing.T) {
	for i, v := range Range(5) {
		if i != v {
			t.Error("Error expected", i, "got", v)
		}
	}

	for i, v := range Range(5, 10) {
		if i+5 != v {
			t.Error("Error expected", i+5, "got", v)
		}
	}

	for i, v := range Range(3, 12, 3) {
		if 3+i*3 != v {
			t.Error("Error expected", 3+i*3, "got", v)
		}
	}

}
