package do

import "testing"

func TestSum(t *testing.T) {
	arr1 := []float32{1.1, 2, 4, 8, 16, 32, 64, 128}
	arr2 := []int{-1, -2, -4, -8, -16, -32, -64, -128}

	result := float32(Sum(arr1).(float64))
	if result != 255.1 {
		t.Error("Error expected", 255.1, "got", result)
	}

	result2 := int(Sum(arr2).(int64))
	if result2 != -255 {
		t.Error("Error expected", -255, "got", result2)
	}
}
