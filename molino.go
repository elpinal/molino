package main

import (
  "./vm"
  "./builtins/go"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "reflect"
)

func main() {
  env := vm.NewEnv()
  core.Import(env)
  for _, arg := range os.Args[1:] {
    scanner := new(vm.Scanner)
    body, err := ioutil.ReadFile(arg)
    if err != nil {
      log.Fatal(err)
    }
    scanner.Init(string(body))
    for _, statement := range vm.Parse(scanner) {
      s, err := vm.Run(statement, env)
      if err != nil {
        log.Fatal(err)
      }
      //fmt.Println(s)
      var keyword vm.Keyword
      o, ok := s.Interface().(fmt.Stringer)
      switch {
      case s.Kind() != reflect.String && ok:
        fmt.Println(o)
      case s.Kind() == reflect.Slice:
        fmt.Println(s.Interface())
      case s.Kind() == reflect.Map:
        fmt.Println(s.Interface())
      case s.Type() == reflect.TypeOf(keyword):
        fmt.Printf(":%v\n", s.Interface())
      default:
        fmt.Printf("%#v\n", s.Interface())
      }
    }
  }
}
