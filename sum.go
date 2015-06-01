package do

import (
	"errors"
	"reflect"
)

// Taken from http://stackoverflow.com/questions/19325152/golang-go-re-post-generic-sum-in-go
func add(a, b interface{}) (interface{}, error) {
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)

	switch valueA.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return valueA.Int() + valueB.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return valueA.Uint() + valueB.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return valueA.Float() + valueB.Float(), nil
	case reflect.String:
		return valueA.String() + valueB.String(), nil
	default:
		return nil, errors.New("Type does not support addition.")
	}
}

/*
Sum sums all the elements of an array and returns the sum as an interface that should be casted as an int64 or float64.
*/
func Sum(iterable interface{}) interface{} {
	iterValue := getIterable(iterable)
	if iterValue.Len() == 0 {
		return 0
	}

	currentValue := iterValue.Index(0).Interface()

	for i := 1; i < iterValue.Len(); i++ {
		newValue, err := add(currentValue, iterValue.Index(i).Interface())
		if err != nil {
			panic(err)
		}
		currentValue = newValue
	}
	return currentValue
}
