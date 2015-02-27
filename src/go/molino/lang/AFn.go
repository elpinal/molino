package lang

import (
	"reflect"
)

type AFn reflect.Value

func NewAFn(f interface{}) IFn {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		panic(fv.Kind().String() + " cannot be cast IFn")
	}
	return AFn(fv)
}

func (a AFn) invoke(args ...interface{}) interface{} {
	var x []reflect.Value
	for _, arg := range args {
		x = append(x, reflect.ValueOf(arg))
	}
	ret := reflect.Value(a).Call(x)
	return ret[0].Interface()
}

func (a AFn) applyTo(arglist ISeq) interface{} {
	//
	return a.invoke(arglist.first())
}
