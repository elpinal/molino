package vm

import (
  _ "fmt"
  _ "reflect"
)

type Namespace struct {
  name Symbol
}

var mappings   = make(map[Symbol]Var)
//  var aliases    = 
var namespaces = make(map[Symbol]Namespace)

func FindOrCreate(name Symbol) Namespace {
  ns, isexist := namespaces[name]
  if isexist {
    return ns
  }
  var newns Namespace = Namespace{name}
  namespaces[name] = newns
  return newns
}

func (this Namespace) intern(sym Symbol) Var {
  if sym.ns != (*string)(nil) {
    panic("ns is empty!")
  }
  a, isexist := mappings[sym]
  if isexist {
    return a
  }
  var v Var = Var{ns: this, sym: sym}
  unbound := &Unbound{&v}
  v.root = unbound
  mappings[sym] = v
  return v
}

func (this Namespace) refer(sym Symbol, v Var) Var {
  if sym.ns != (*string)(nil) {
    panic("ns is empty!")
  }
  a, isexist := mappings[sym]
  if isexist {
    return a
  }
  mappings[sym] = v
  return v
}
