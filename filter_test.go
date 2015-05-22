package gogo

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	iterable := []int{9, 8, 7, 6}
	values := Filter(func(v int) bool { return v%2 == 0 }, iterable)
	fmt.Println(values)
}
