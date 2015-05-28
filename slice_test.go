package do

import (
	"reflect"
	"testing"
)

func TestSliceStruct(t *testing.T) {

	type uniqueStruct struct {
		val int
	}

	iterable := []uniqueStruct{
		uniqueStruct{8},
		uniqueStruct{9},
		uniqueStruct{8},
		uniqueStruct{5},
		uniqueStruct{7},
		uniqueStruct{6},
		uniqueStruct{5},
	}

	result := []uniqueStruct{
		uniqueStruct{9},
		uniqueStruct{5},
		uniqueStruct{6},
	}

	values := Slice(iterable, 1, 6, 2)
	if !reflect.DeepEqual(values, result) {
		t.Error("Error expected", result, "got", values)
	}
}
