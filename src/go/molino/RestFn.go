package molino

import (
	"reflect"
)

type RestFn reflect.Value

func NewRestFn(f interface{}, requiredArity int) IFn {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		panic(fv.Kind().String() + " cannot be cast IFn")
	}
	return RestFn(fv)
}

func (a RestFn) invoke(args ...interface{}) interface{} {
	var x []reflect.Value
	for _, arg := range args {
		x = append(x, reflect.ValueOf(arg))
	}
	ret := reflect.Value(a).Call(x)
	return ret[0].Interface()
}

func (a RestFn) applyTo(args ISeq) interface{} {
	// FIXME
	return a.invoke(args)
}
