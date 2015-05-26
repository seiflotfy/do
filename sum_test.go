package do

import "testing"

func TestSum(t *testing.T) {
	arr1 := []int{1, 2, 4, 8, 16, 32, 64, 128}
	arr2 := []int{-1, -2, -4, -8, -16, -32, -64, -128}

	result := Sum(arr1)
	if result != 255 {
		t.Error("Error expected", 255, "got", result)
	}

	result = Sum(arr2)
	if result != -255 {
		t.Error("Error expected", -255, "got", result)
	}
}

func TestSum32(t *testing.T) {
	arr1 := []int32{1, 2, 4, 8, 16, 32, 64, 128}
	arr2 := []int32{-1, -2, -4, -8, -16, -32, -64, -128}

	result := Sum32(arr1)
	if result != 255 {
		t.Error("Error expected", 255, "got", result)
	}

	result = Sum32(arr2)
	if result != -255 {
		t.Error("Error expected", -255, "got", result)
	}
}

func TestSum64(t *testing.T) {
	arr1 := []int64{1, 2, 4, 8, 16, 32, 64, 128}
	arr2 := []int64{-1, -2, -4, -8, -16, -32, -64, -128}

	result := Sum64(arr1)
	if result != 255 {
		t.Error("Error expected", 255, "got", result)
	}

	result = Sum64(arr2)
	if result != -255 {
		t.Error("Error expected", -255, "got", result)
	}
}
