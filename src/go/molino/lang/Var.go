package lang

import (
	"fmt"
	"reflect"
)

type Var struct {
	ns          Namespace
	sym         Symbol
	root        interface{}
	threadBound bool
	//dynamic bool
}

type TBox struct {
	val interface{}
}

type Unbound struct {
	Var
}

type Frame struct {
	bindings Associative
	prev     *Frame
}

var dvals Frame = Frame{bindings: PersistentHashMap{}, prev: nil}

func (this Var) intern(ns Namespace, sym Symbol, root interface{}, replaceRoot bool) Var {
	var dvout Var = ns.intern(sym)
	if !dvout.hasRoot() || replaceRoot {
		dvout.bindroot(root)
	}
	ns.updatemapping(sym, dvout)
	return dvout
}

func (v Var) create() Var {
	ret := Var{}
	ret.root = Unbound{ret}
	return ret
}

func (this *Var) bindroot(root interface{}) {
	//  oldroot := this.root
	this.root = root
}

func (v Var) invoke(arg1 interface{}) interface{} {
	fn := reflect.ValueOf(v.root)
	w := []reflect.Value{reflect.ValueOf(arg1)}
	x := fn.Call(w)
	return x
}

func (v Var) isBound() bool {
	return v.hasRoot()
}

func (v Var) get() interface{} {
	if !v.threadBound {
		return v.root
	}
	return v.deref()
}

func (v Var) deref() interface{} {
	var b TBox = v.getThreadBinding()
	if b.val != nil {
		return b.val
	}
	return v.root
}

func (v Var) set(val interface{}) interface{} {
	//
	var b TBox = v.getThreadBinding()
	if b.val != nil {
		b.val = val
		return val
	}
	panic(fmt.Sprintf("Can't change/establish root binding of: %s with set", v.sym))
}

func (v Var) hasRoot() bool {
	_, ok := v.root.(Unbound)
	return !ok
}

func (_ Var) pushThreadBinding(bindings Associative) {
	var f Frame = dvals
	var bmap Associative = f.bindings
	for bs := bindings.seq(); bs != nil; bs = bs.next() {
		var e IMapEntry = bs.first().(IMapEntry)
		var v Var = e.key().(Var)
		//if !v.dynamic {
		//	panic(fmt.Sprintf("Can't dynamically bind non-dynamic var: %s/%s", v.ns, v.sym))
		//}
		//
		v.threadBound = true
		bmap = bmap.assoc(v, TBox{val: e.val()})
	}
	dvals = Frame{bindings: bmap, prev: &f}
}

func (v Var) getThreadBinding() TBox {
	if v.threadBound {
		var e IMapEntry = dvals.bindings.entryAt(v)
		if e != nil {
			return e.val().(TBox)
		}
	}
	return TBox{val: nil}
}
