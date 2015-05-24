package gogo

import (
	"fmt"
	"reflect"
)

/*
Reduce applies function of two arguments cumulatively to the items of iterable,
from left to right, so as to reduce the iterable to a single value adn return it.
*/
func Reduce(function interface{}, iterable interface{}) interface{} {
	iterValue := getIterable(iterable)
	funcValue := getFunction(function, iterValue.Type().Elem(), iterValue.Type().Elem(), 2)
	returnValue := reflect.New(iterValue.Type().Elem()).Elem()
	startValue := iterValue.Index(0)

	if !startValue.Type().AssignableTo(returnValue.Type()) && !startValue.Type().ConvertibleTo(returnValue.Type()) {
		panic(fmt.Errorf("second argument must be convertible to %s, not %T", returnValue.Type(), startValue))
	}

	if !startValue.Type().AssignableTo(returnValue.Type()) {
		startValue = startValue.Convert(returnValue.Type())
	}

	returnValue.Set(startValue)
	funcInput := make([]reflect.Value, 2)
	for i := 1; i < iterValue.Len(); i++ {
		funcInput[0], funcInput[1] = returnValue, iterValue.Index(i)
		res := funcValue.Call(funcInput)
		returnValue.Set(res[0])
	}
	return returnValue.Interface()
}
