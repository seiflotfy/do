package do

import "reflect"

/*
Reversed returns the values of a given slice or array in reverse order.
*/
func Reversed(iterable interface{}) interface{} {
	iterValue := getIterable(iterable)
	iterLen := iterValue.Len()
	returnValue := reflect.MakeSlice(iterValue.Type(), iterLen, iterLen)

	for i := iterLen - 1; i > -1; i-- {
		returnValue.Index(iterLen - 1 - i).Set(iterValue.Index(i))
	}

	return returnValue.Interface()
}
