package main

import (
  "./vm"
  go_core "./builtins/go"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  _ "reflect"
)

var fs = flag.NewFlagSet(os.Args[0], 1)
var e  = fs.String("e", "", "One line of program")

func main() {
  fs.Parse(os.Args[1:])
  env := vm.NewEnv()
  var body []byte
  var source string
  if *e != "" {
    body  = []byte(*e)
    source = "argument"
  } else {
    var err error
    arg := os.Args[1]
    source = fs.Arg(0)
    body, err = ioutil.ReadFile(arg)
    if err != nil {
      log.Fatal(err)
    }
  }
  os.Args = fs.Args()

  go_core.Import(env)

  vm.Runtime()

  scanner := new(vm.Scanner)
  scanner.Init(string(body))
  for _, statement := range vm.Parse(scanner) {
    _, err := vm.Run(statement, env)
    if err != nil {
      fmt.Printf("%s: ", source)
      log.Fatal(err)
    }
    //fmt.Println(s)
/*
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
 */
  }
}
