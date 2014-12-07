package core

import (
  "../../vm"
  "reflect"
  "fmt"
  "bytes"
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

  env.Define("println", reflect.ValueOf(func(s ...reflect.Value) (reflect.Value, error) {
    var buffer bytes.Buffer
    for i := 0; i < len(s) ; i++ {
      if i + 1 != len(s) {
        buffer.WriteString(molinoprint(s[i]))
      } else {
        buffer.WriteString(molinoprint(s[i]))
      }
    }
    fmt.Println(buffer.String())
    return reflect.ValueOf(nil), nil
  }))

  env.Define("pr-on", reflect.ValueOf(func(s reflect.Value) (reflect.Value, error) {
    fmt.Print(molinoprint(s))
    return reflect.ValueOf(nil), nil
  }))

  env.Define("first", reflect.ValueOf(func(x reflect.Value) (reflect.Value, error) {
    v := x.Interface().([]interface{})
    if len(v) > 0 {
      return reflect.ValueOf(v[0]), nil
    }
    return reflect.ValueOf(nil), nil
  }))

  env.Define("next", reflect.ValueOf(func(x reflect.Value) (reflect.Value, error) {
    v := x.Interface().([]interface{})
    if len(v) > 1 {
      var a = make([]interface{}, len(v) - 1)
      for i := 1; i < len(v); i++ {
        a[i-1] = v[i]
      }
      return reflect.ValueOf(a), nil
    }
    return reflect.ValueOf(nil), nil
  }))
}

func molinoprint(s reflect.Value) string {
  var keyword vm.Keyword
  switch {
  case s == reflect.ValueOf(nil):
    return fmt.Sprint("nil")
  case s.Type() == reflect.TypeOf(keyword):
    return fmt.Sprintf(":%v", s.Interface())
  }
  return fmt.Sprint(s.Interface())
}
