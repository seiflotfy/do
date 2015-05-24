package gogo

import (
	"reflect"
	"sync"
)

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

/*
MapParallel ...
Apply function to every item of iterable and return a list of the results. numParallel defines the number of items
the function will be applied to in parallel
*/
func MapParallel(function interface{}, iterable interface{}, numParallel int) interface{} {

	if numParallel < 1 {
		numParallel = 2
	}

	iterValue := getIterable(iterable)
	funcValue := getFunction(function, iterValue.Type().Elem(), nil, 1)
	returnValue := reflect.MakeSlice(reflect.SliceOf(funcValue.Type().Out(0)), iterValue.Len(), iterValue.Len())
	funcInput := make([]reflect.Value, 1)

	var wg sync.WaitGroup
	defer wg.Wait()

	var pFunc = func(j int) {
		defer wg.Done()
		funcInput[0] = iterValue.Index(j)
		res := funcValue.Call(funcInput)
		returnValue.Index(j).Set(res[0])
	}

	for i := 0; i < iterValue.Len(); i++ {
		wg.Add(1)
		go pFunc(i)
		if (i+1)%numParallel == 0 {
			wg.Wait()
		}
	}

	return returnValue.Interface()
}
