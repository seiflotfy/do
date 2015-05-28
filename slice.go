package do

import "reflect"

/*
Slice returns a list of elements from iterable from start to stop in step length
*/
func Slice(iterable interface{}, start int, stop int, step int) interface{} {
	iterValue := getIterable(iterable)
	returnValue := reflect.MakeSlice(iterValue.Type(), 0, 0)
	for i := start; i < stop; i++ {
		if (i+1)%(step) == 0 {
			returnValue = reflect.Append(returnValue, iterValue.Index(i))
		}
	}
	return returnValue.Interface()
}
