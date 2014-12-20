package lang

import (
  "fmt"
  "reflect"
  "io/ioutil"
  "log"
  "bytes"
  "errors"
)

var MOLINO_NS    Namespace = FindOrCreate(intern("molino.core"))
var IN_NAMESPACE Symbol    = intern("in-ns")
var NAMESPACE    Symbol    = intern("ns")

var VAR Var
var CURRENT_NS Var = VAR.intern(MOLINO_NS, intern("*ns*"), MOLINO_NS, true)

var inNamespace = func(arg1 reflect.Value) (Namespace, error) {
    var nsname Symbol = arg1.Interface().(Symbol)
    var ns Namespace = FindOrCreate(nsname)
//    CURRENT_NS.set(ns)
    //CURRENT_NS.bindroot(ns)
    CURRENT_NS.root = ns
    return ns, nil
}


func Runtime() {
  //fmt.Println(MOLINO_NS, NAMESPACE, IN_NAMESPACE.name)
  var v Var
  //var s Symbol = intern("user")
  v = v.intern(MOLINO_NS, IN_NAMESPACE, inNamespace, true)
  doInit()
  //v.invoke(reflect.ValueOf(s))
}

func doInit() {
  load("molino/core")
}

func load(scriptbase string) {
  body, err := ioutil.ReadFile("src/mln/" + scriptbase + ".mln")
  if err != nil {
    log.Fatal(err)
  }
  env := NewEnv()
  refImport(env)
  reader := new(Reader)
  reader.Init(string(body))
  for _, statement := range Parse(reader) {
    fmt.Printf("%#v\n", statement.(*ExpressionStatement).Expr)
    _, err := Run(statement, env)
    if err != nil {
      fmt.Printf("%s: ", scriptbase)
      log.Fatal(err)
    }
  }
}


//===============


func refImport(env *Env) {
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
        buffer.WriteString(themolinoprint(s[i]))
      } else {
        buffer.WriteString(themolinoprint(s[i]))
      }
    }
    fmt.Println(buffer.String())
    return reflect.ValueOf(nil), nil
  }))

  env.Define("pr-on", reflect.ValueOf(func(s reflect.Value) (reflect.Value, error) {
    fmt.Print(themolinoprint(s))
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
    v, ok := x.Interface().([]interface{})
    if !ok {
      return reflect.ValueOf(v), errors.New(fmt.Sprint(x.Kind(), " cannot be cast to vector"))
    }
    if len(v) > 1 {
      var a = make([]interface{}, len(v) - 1)
      for i := 1; i < len(v); i++ {
        a[i-1] = v[i]
      }
      return reflect.ValueOf(a), nil
    }
    return reflect.ValueOf(nil), nil
  }))

  env.Define("applyTo", reflect.ValueOf(func(f, x reflect.Value) (reflect.Value, error) {
    if f.Kind() != reflect.Func {
      return reflect.ValueOf(f), errors.New(fmt.Sprint("Unknown Function: ", f.Interface()))
    }
    v, ok := x.Interface().([]interface{})
    if !ok {
      return reflect.ValueOf(f), errors.New(fmt.Sprint("Don't know how to create vector from: ", x.Interface()))
    }
    var arg = make([]reflect.Value, len(v))
    for i, a := range v {
      arg[i] = reflect.ValueOf(reflect.ValueOf(a))
    }
    //var arg = []reflect.Value{reflect.ValueOf(x)}
    r := f.Call(arg)
    if r[1].Interface() != nil {
      return r[0], r[1].Interface().(error)
    }
    return r[0].Interface().(reflect.Value), nil
  }))
  env.Define("multiply", reflect.ValueOf(func(x, y reflect.Value) (reflect.Value, error) {
    xx := x.Int()
    yy := y.Int()
    return reflect.ValueOf(xx * yy), nil
  }))
  env.Define("minus", reflect.ValueOf(func(x, y reflect.Value) (reflect.Value, error) {
    xx := x.Int()
    yy := y.Int()
    return reflect.ValueOf(xx - yy), nil
  }))
  env.Define("add", reflect.ValueOf(func(x, y reflect.Value) (reflect.Value, error) {
    xx := x.Int()
    yy := y.Int()
    return reflect.ValueOf(xx + yy), nil
  }))
  env.Define("equiv", reflect.ValueOf(func(x, y reflect.Value) (reflect.Value, error) {
    xx, yy := x.Interface(), y.Interface()
    return reflect.ValueOf(xx == yy), nil
  }))
}

func themolinoprint(s reflect.Value) string {
  var keyword Keyword
  switch {
  case s == reflect.ValueOf(nil):
    return fmt.Sprint("nil")
  case s.Type() == reflect.TypeOf(keyword):
    return fmt.Sprintf(":%v", s.Interface())
  }
  return fmt.Sprint(s.Interface())
}
