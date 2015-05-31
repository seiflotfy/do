package do

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{5, 4, 3, 2, 1}

	values := Reverse(arr1)
	if !reflect.DeepEqual(values, []int{5, 4, 3, 2, 1}) {
		t.Error("Error expected", []int{5, 4, 3, 2, 1}, "got", values)
	}

	values = Reverse(arr2)
	if !reflect.DeepEqual(values, []int{1, 2, 3, 4, 5}) {
		t.Error("Error expected", []int{1, 2, 3, 4, 5}, "got", values)
	}
}
