package gogo

import "reflect"

/*
Filter returns a list of elements from iterable for which function returns true.
*/
func Filter(function interface{}, iterable interface{}) interface{} {
	iterValue := getIterable(iterable)
	funcValue := getFunction(function, iterValue.Type().Elem(), reflect.TypeOf(true), 1)
	returnValue := reflect.MakeSlice(iterValue.Type(), 0, 0)
	funcInput := make([]reflect.Value, 1)
	for i := 0; i < iterValue.Len(); i++ {
		funcInput[0] = iterValue.Index(i)
		res := funcValue.Call(funcInput)
		if res[0].Bool() {
			returnValue = reflect.Append(returnValue, funcInput[0])
		}
	}
	return returnValue.Interface()
}
