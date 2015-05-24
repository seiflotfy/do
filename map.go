package gogo

import (
	"reflect"
	"sync"
)

/*
Map apply function to every item of iterable and return a list of the results..
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
MapParallel apply function to every item of iterable and return a list of the results. numParallel defines the number of items
the function will be applied to in parallel
*/
func MapParallel(function interface{}, iterable interface{}, numParallel int) interface{} {

	if numParallel < 1 {
		numParallel = 2
	}

	iterValue := getIterable(iterable)
	funcValue := getFunction(function, iterValue.Type().Elem(), nil, 1)
	returnValue := reflect.MakeSlice(reflect.SliceOf(funcValue.Type().Out(0)), iterValue.Len(), iterValue.Len())

	var wg sync.WaitGroup
	defer wg.Wait()
	resChan := make(chan parallelStruct)

	var pFunc = func(j int) {
		defer wg.Done()
		funcInput := make([]reflect.Value, 1)
		funcInput[0] = iterValue.Index(j)
		res := funcValue.Call(funcInput)
		resChan <- parallelStruct{j, res[0]}
	}
	for i := 0; i < iterValue.Len(); i++ {
		wg.Add(1)
		go pFunc(i)
		if (i+1)%numParallel == 0 {
			for j := 0; j < numParallel; j++ {
				pS := <-resChan
				returnValue.Index(pS.pos).Set(pS.val)
			}
		}
	}
	for j := 0; j < iterValue.Len()%numParallel; j++ {
		pS := <-resChan
		returnValue.Index(pS.pos).Set(pS.val)
	}
	return returnValue.Interface()
}

type parallelStruct struct {
	pos int
	val reflect.Value
}
