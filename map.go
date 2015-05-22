package gogo

import "reflect"

/*
Map ...
Apply function to every item of iterable and return a list of the results..
*/
func Map(function interface{}, iterable interface{}) interface{} {
	iterValue := getIterable(iterable)
	funcValue := getFunction(function, iterValue.Type().Elem(), nil, 1)
	returnValue := reflect.MakeSlice(reflect.SliceOf(funcValue.Type().Out(0)), iterValue.Len(), iterValue.Len())
	funcInput := make([]reflect.Value, 1)
	for i := 0; i < iterValue.Len(); i++ {
		funcInput[0] = iterValue.Index(i)
		res := funcValue.Call(funcInput)
		returnValue.Index(i).Set(res[0])
	}
	return returnValue.Interface()
}
