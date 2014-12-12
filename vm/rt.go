package vm

import (
  "fmt"
)

var MOLINO_NS    Namespace = FindOrCreate(intern("molino.core"))
var IN_NAMESPACE Symbol    = intern("in-ns")
var NAMESPACE    Symbol    = intern("ns")

var VAR Var
var CURRENT_NS Var = VAR.intern(MOLINO_NS, intern("*ns*"), MOLINO_NS, true)

var inNamespace = func(arg1 Symbol) Namespace {
    var nsname Symbol = arg1
    var ns Namespace = FindOrCreate(nsname)
//    CURRENT_NS.set(ns)
    CURRENT_NS.bindroot(ns)
    return ns
}


func Runtime() {
  fmt.Println(MOLINO_NS, NAMESPACE, *IN_NAMESPACE.name)
  var v Var
  fmt.Println(v.intern(MOLINO_NS, IN_NAMESPACE, inNamespace, true))
}

/*
func doInit() {
  load("molino/core")
}
*/
