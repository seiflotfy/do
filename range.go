package gogo

import "fmt"

/*
Range returns a list of ints in a range
*/
func Range(args ...int) []int {
	start, end := 0, 0
	step := 1
	if len(args) >= 1 {
		end = args[0]
	}
	if len(args) >= 2 {
		start = args[0]
		end = args[1]
		if end <= start {
			panic(fmt.Errorf("start %d is larger than end %d", start, end))
		}
	}
	if len(args) == 3 {
		step = args[2]
	}

	values := make([]int, ((end - start) / step))
	for i := 0; i < len(values); i++ {
		v := start + i*step
		if v >= end {
			return values
		}
		values[i] = v
	}

	return values
}
