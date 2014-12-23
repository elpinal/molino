package lang

import (
	_ "fmt"
	_ "reflect"
)

type Namespace struct {
	name     Symbol
	mappings map[Symbol]Var
}

//var mappings   = make(map[Symbol]Var)
//  var aliases    =

var namespaces = make(map[Symbol]Namespace)

func FindOrCreate(name Symbol) Namespace {
	ns, isexist := namespaces[name]
	if isexist {
		return ns
	}
	var newns Namespace = Namespace{name: name, mappings: make(map[Symbol]Var)}
	namespaces[name] = newns
	return newns
}

func (this Namespace) intern(sym Symbol) Var {
	if sym.ns != "" {
		panic("ns is not empty!")
	}
	a, isexist := this.mappings[sym]
	if isexist {
		return a
	}
	var v Var = Var{ns: this, sym: sym}
	unbound := Unbound{v}
	v.root = unbound
	this.mappings[sym] = v
	return v
}

func (this Namespace) refer(sym Symbol, v Var) Var {
	if sym.ns != "" {
		panic("ns is not empty!")
	}
	a, isexist := this.mappings[sym]
	if isexist {
		return a
	}
	this.mappings[sym] = v
	return v
}

func (this Namespace) getmapping(name Symbol) (Var, bool) {
	v, isexist := this.mappings[name]
	return v, isexist
}

func (this Namespace) updatemapping(name Symbol, newval Var) {
	if _, isexist := this.mappings[name]; isexist {
		this.mappings[name] = newval
	}
}

func findNamespace(name Symbol) (Namespace, bool) {
	v, isexist := namespaces[name]
	return v, isexist
}

func (this Namespace) findInternedVar(sym Symbol) (Var, bool) {
	v, isexist := this.mappings[sym]
	if isexist && v.ns.name == this.name {
		return v, true
	}
	return v, false
}
