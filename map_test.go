package gogo

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	iterable := []int{9, 8, 7, 6}
	values := Map(func(v int) bool { return v%2 == 0 }, iterable)
	fmt.Println(values)
}
