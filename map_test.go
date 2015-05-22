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
