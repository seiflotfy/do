package do

import (
	"reflect"
	"testing"
)

func TestUniquify(t *testing.T) {
	iterable := []int{8, 9, 8, 5, 7, 6, 5}
	values := Uniquify(func(v int) int { return v }, iterable)
	if !reflect.DeepEqual(values, []int{8, 9, 5, 7, 6}) {
		t.Error("Error expected", []int{8, 9, 5, 7, 6}, "got", values)
	}
}

func TestUniquifyStruct(t *testing.T) {

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
		uniqueStruct{8},
		uniqueStruct{9},
		uniqueStruct{5},
		uniqueStruct{7},
		uniqueStruct{6},
	}

	values := Uniquify(func(v uniqueStruct) int { return v.val }, iterable)
	if !reflect.DeepEqual(values, result) {
		t.Error("Error expected", result, "got", values)
	}
}
