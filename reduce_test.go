package do

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	iterable := []int{9, 8, 7, 6}
	values := Reduce(func(v1 int, v2 int) int { return v1 * v2 }, iterable)
	if !reflect.DeepEqual(values, 9*8*7*6) {
		t.Error("Error expected", 9*8*7*6, "got", values)
	}
}
