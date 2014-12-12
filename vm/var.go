package vm

import (
  _ "fmt"
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
  if replaceRoot {
		dvout.bindroot(root)
  }
	return dvout
}

func (this Var) bindroot(root interface{}) {
//  oldroot := this.root
  this.root = root
}

/*
func (this Var) set() {
  //
}
*/
