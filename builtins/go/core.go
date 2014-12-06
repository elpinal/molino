package core

import (
  "../../vm"
  "reflect"
//  "errors"
)

func Import(env *vm.Env) {
  env.Define("cons", reflect.ValueOf(func(x, seq reflect.Value) (reflect.Value, error) {
    a := []interface{}{x.Interface()}
    s := seq.Interface().([]interface{})
    for _, v := range s {
      a = append(a, v)
    }
    return reflect.ValueOf(a), nil
  }))
}
