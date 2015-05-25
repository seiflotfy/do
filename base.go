package do

import (
	"fmt"
	"reflect"
)

func getIterable(iterable interface{}) reflect.Value {
	v := reflect.ValueOf(iterable)
	if !v.IsValid() {
		panic(fmt.Errorf("%T is of of invalid type", iterable))
	} else if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic(fmt.Errorf("%T is not of type array or slice, but of type %T", iterable, v.Kind()))
	}
	return v
}

func getFunction(function interface{}, argType reflect.Type, retType reflect.Type, numArgs int) reflect.Value {
	v := reflect.ValueOf(function)
	if !v.IsValid() {
		panic(fmt.Errorf("%T is of invalid type", function))
	} else if v.Kind() != reflect.Func {
		panic(fmt.Errorf("%T is not of type func(%s) %s", function, argType.Elem(), retType))
	} else if v.Type().NumIn() != numArgs {
		panic(fmt.Errorf("%T is not of type func(%s) %s: number of input != 1", function, argType, retType))
	} else if v.Type().NumOut() != 1 {
		panic(fmt.Errorf("%T is not of type func(%s) %s: number of output != 1", function, argType, retType))
	} else if retType != nil && v.Type().Out(0) != retType {
		panic(fmt.Errorf("%T is not of type func(%s) %s: return not of type %s", function, argType, retType, retType))
	}

	for i := 0; i < numArgs; i++ {
		if v.Type().In(i) != argType {
			panic(fmt.Errorf("%T is not of type func(%s) %s: input not of type %s", function, argType, retType, argType))
		}
	}

	return v
}
