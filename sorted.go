package do

import (
	"errors"
	"sort"
)

// Sorted takes an iterable (slice) of elements and returns a new slice containing
// the elements sorted in ascending order.
// Example:
//
//	sortedInts, err := Sorted([]int{3, 1, 4, 1, 5, 9})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// sortedInts will be []int{1, 1, 3, 4, 5, 9}
func Sorted(iterable interface{}) (interface{}, error) {
	switch v := iterable.(type) {
	case []int:
		sortedSlice := make([]int, len(v))
		copy(sortedSlice, v)
		sort.Ints(sortedSlice)
		return sortedSlice, nil
	case []float64:
		sortedSlice := make([]float64, len(v))
		copy(sortedSlice, v)
		sort.Float64s(sortedSlice)
		return sortedSlice, nil
	case []string:
		sortedSlice := make([]string, len(v))
		copy(sortedSlice, v)
		sort.Strings(sortedSlice)
		return sortedSlice, nil
	default:
		return nil, errors.New("Unsupported type")
	}
}
