package do

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	iterable := []int{9, 8, 7, 6}
	values := Filter(func(v int) bool { return v%2 == 0 }, iterable)
	if !reflect.DeepEqual(values, []int{8, 6}) {
		t.Error("Error expected", []int{8, 6}, "got", values)
	}
}
