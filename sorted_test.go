package do

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func generateRandomIntSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func TestSortedInt(t *testing.T) {
	iterable := []int{3, 1, 4, 1, 5, 9}
	expected := []int{1, 1, 3, 4, 5, 9}
	values, err := Sorted(iterable)
	if err != nil {
		t.Error("Error occurred:", err)
	}
	if !reflect.DeepEqual(values, expected) {
		t.Error("Error expected", expected, "got", values)
	}
}

func TestSortedFloat64(t *testing.T) {
	iterable := []float64{3.1, 1.4, 4.1, 5.9}
	expected := []float64{1.4, 3.1, 4.1, 5.9}
	values, err := Sorted(iterable)
	if err != nil {
		t.Error("Error occurred:", err)
	}
	if !reflect.DeepEqual(values, expected) {
		t.Error("Error expected", expected, "got", values)
	}
}

func TestSortedString(t *testing.T) {
	iterable := []string{"apple", "banana", "cherry"}
	expected := []string{"apple", "banana", "cherry"}
	values, err := Sorted(iterable)
	if err != nil {
		t.Error("Error occurred:", err)
	}
	if !reflect.DeepEqual(values, expected) {
		t.Error("Error expected", expected, "got", values)
	}
}

func TestSortedSpeed(t *testing.T) {
	// Generate a large random slice of integers
	iterable := generateRandomIntSlice(10000)

	// Measure the time taken by Sorted function
	start := time.Now()
	values1, err := Sorted(iterable)
	if err != nil {
		t.Error("Error occurred:", err)
	}
	elapsed1 := time.Since(start)

	// For comparison, sort the slice using Go's built-in sort for int slices
	start = time.Now()
	values2 := make([]int, len(iterable))
	copy(values2, iterable)
	sort.Ints(values2)
	elapsed2 := time.Since(start)

	// Check if both sorted slices are equal
	if !reflect.DeepEqual(values1, values2) {
		t.Error("Error expected", values2, "got", values1)
	}

	// Check if custom Sorted function is not excessively slower than Go's built-in sort
	if elapsed1 > 2*elapsed2 {
		t.Error("Error expected Sorted to be not excessively slower than Go's built-in sort, got", elapsed1, ">", 2*elapsed2)
	}
}
