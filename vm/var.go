package vm

import (
  _ "fmt"
  "reflect"
)

type Var struct {
  ns   Namespace
  sym  Symbol
//  this.threadBound
  root interface{}
}

type Unbound struct {
  *Var
}

func (this Var) intern(ns Namespace, sym Symbol, root interface{}, replaceRoot bool) Var {
  var dvout Var = ns.intern(sym)
  /*
  if replaceRoot {
    dvout.bindroot(root)
  }
  */
  dvout.root = root
  updatemapping(sym, dvout)
  return dvout
}

/*
func (this Var) bindroot(root interface{}) {
//  oldroot := this.root
  this.root = root
}
*/

func (v Var) invoke(arg1 interface{}) interface{} {
  fn := reflect.ValueOf(v.root)
  w := []reflect.Value{reflect.ValueOf(arg1)}
  x := fn.Call(w)
  return x
}

/*
func (this Var) set() {
  //
}
*/
