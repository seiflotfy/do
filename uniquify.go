package do

import "reflect"

/*
Uniquify returns a list of unique elements from iterable.
*/
func Uniquify(function interface{}, iterable interface{}) interface{} {
	iterValue := getIterable(iterable)
	funcValue := getFunction(function, iterValue.Type().Elem(), nil, 1)
	returnValue := reflect.MakeSlice(iterValue.Type(), 0, 0)
	funcInput := make([]reflect.Value, 1)

	uniques := make(map[interface{}]interface{})

	for i := 0; i < iterValue.Len(); i++ {
		funcInput[0] = iterValue.Index(i)
		res := funcValue.Call(funcInput)

		_, ok := uniques[res[0].Interface()]
		if !ok {
			uniques[res[0].Interface()] = funcInput[0]
			returnValue = reflect.Append(returnValue, funcInput[0])
		}

	}
	return returnValue.Interface()
}
