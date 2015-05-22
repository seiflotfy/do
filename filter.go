package gogo

import "reflect"

/*
Filter ...
Return true if any elements of the iterable are true, the iterable is empty, return false.
*/
func Filter(function interface{}, iterable interface{}) interface{} {
	iterValue := getIterable(iterable)
	funcValue := getFunction(function, iterValue.Type().Elem(), reflect.TypeOf(true))
	newIterable := reflect.MakeSlice(iterValue.Type(), 0, 0)
	returnValue := make([]reflect.Value, 1)
	for i := 0; i < iterValue.Len(); i++ {
		returnValue[0] = iterValue.Index(i)
		res := funcValue.Call(returnValue)
		if res[0].Bool() {
			newIterable = reflect.Append(newIterable, returnValue[0])
		}
	}
	return newIterable.Interface()
}
