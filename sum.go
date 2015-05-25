package do

/*
Sum ...
Sums all the elements of an array and returns the sum.
*/
func Sum(arr []int) int {
	var result int = 0
	for _, value := range arr {
		result += value
	}
	return result
}

/*
Sum32 ...
Sums all the elements of an int32 array and returns the sum as an int32.
*/
func Sum32(arr []int32) int32 {
	var result int32 = 0
	for _, value := range arr {
		result += value
	}
	return result
}

/*
Sum64 ...
Sums all the elements of an int64 array and returns the sum as an int64.
*/
func Sum64(arr []int64) int64 {
	var result int64 = 0
	for _, value := range arr {
		result += value
	}
	return result
}

