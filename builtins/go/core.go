package core

import (
  "../../vm"
  "reflect"
//  "errors"
)

func Import(env *vm.Env) {
  env.Define("cons", reflect.ValueOf(func(x reflect.Value, seq reflect.Value) (reflect.Value, error) {
    var a []interface{}
    a = append(a, x.Interface())
    s := seq.Interface().([]interface{})
    for i, _ := range s {
      a = append(a, s[i])
    }
    return reflect.ValueOf(a), nil
  }))
}

//func Cons(x interface{}, seq []interface{}) []interface{} {
//  return append(seq, x)
//}

