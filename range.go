package gogo

/*
Range
Returns a slice of ints in a range
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
	}
	if len(args) == 3 {
		step = args[2]
	}

	values := make([]int, ((end-start)/step)+1)
	for i := 0; i < len(values); i++ {
		values[i] = start + i*step
	}

	return values
}
