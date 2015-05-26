package do

/*
Sum ...
Sums all the elements of an array and returns the sum.
*/
func Sum(arr []int) int {
	result := 0
	for _, value := range arr {
		result += value
	}
	return result
}

/*
Sum32 ...
Sums all the elements of an int32 array and returns the sum as an int32.
*/
func Sum32(arr []int32) (result int32) {
	result = 0
	for _, value := range arr {
		result += value
	}
	return result
}

/*
Sum64 ...
Sums all the elements of an int64 array and returns the sum as an int64.
*/
func Sum64(arr []int64) (result int64) {
	result = 0
	for _, value := range arr {
		result += value
	}
	return result
}
