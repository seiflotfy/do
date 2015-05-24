package gogo

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	iterable := []int{9, 8, 7, 6}
	values := Map(func(v int) bool { return v%2 == 0 }, iterable)
	if !reflect.DeepEqual(values, []bool{false, true, false, true}) {
		t.Error("Error expected", []bool{false, true, false, true}, "got", values)
	}
}

func TestMapParallel(t *testing.T) {
	iterable := []int{9, 8, 7, 6, 5, 6, 6, 1}
	values := MapParallel(func(v int) bool { return v%2 == 0 }, iterable, 2)
	if !reflect.DeepEqual(values, []bool{false, true, false, true, false, true, true, false}) {
		t.Error("Error expected", []bool{false, true, false, true, false, true, true, false}, "got", values)
	}
}
