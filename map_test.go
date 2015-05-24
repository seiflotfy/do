package gogo

import (
	"reflect"
	"testing"
	"time"
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

func TestSpeed(t *testing.T) {
	iterable := Range(1000)
	start := time.Now()
	values1 := Map(func(v int) bool {
		time.Sleep(1 * time.Millisecond)
		return v%2 == 0
	}, iterable)
	elapsed1 := time.Since(start)
	start = time.Now()
	values2 := MapParallel(func(v int) bool {
		time.Sleep(1 * time.Millisecond)
		return v%2 == 0
	}, iterable, 4)
	elapsed2 := time.Since(start)
	if !reflect.DeepEqual(values1, values2) {
		t.Error("Error expected", values1, "got", values2)
	}
	if elapsed1 <= elapsed2 {
		t.Error("Error expected Map to be slower than MapParallel, got", elapsed1, "<=", elapsed2)
	}
}
