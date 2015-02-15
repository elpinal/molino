package lang

import (
	"io/ioutil"
	"log"
	"reflect"
	"fmt"
)

var MOLINO_NS Namespace = FindOrCreate(intern("molino.core"))
var IN_NAMESPACE Symbol = intern("in-ns")
var NAMESPACE Symbol = intern("ns")

var CURRENT_NS Var = Var{}.intern(MOLINO_NS, intern("*ns*"), MOLINO_NS, true)

var NS_VAR Var = Var{}.intern(MOLINO_NS, intern("ns"), false, true)
var IN_NS_VAR Var = Var{}.intern(MOLINO_NS, intern("in-ns"), false, true)

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
	reader := new(Reader)
	reader.Init(string(body))
	var ret interface{}
	ret, err = Compiler.load(Compiler{}, reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
}


////////////// Collections support /////////////////////////////////

func seq(coll interface{}) ISeq {
	if s, ok := coll.(ISeq); ok {
		return s
	}
	return seqFrom(coll)
}

func seqFrom(coll interface{}) ISeq {
	switch coll.(type) {
	case Seqable:
		return coll.(Seqable).seq()
	case Iterable:
		ret, _ := IteratorSeq.create(IteratorSeq{}, coll.(Iterable).iterator())
		return ret
	case []interface{}:
		return createFromObject(coll.([]interface{}))
	}
	//
	panic("Don't know how to create ISeq from: " + fmt.Sprintf("%T\n", coll))
}

func conj(coll IPersistentCollection, x interface{}) IPersistentCollection {
	if coll == nil {
		return PersistentList{_first: x, _count: 1}
	}
	return coll.(ISeq).cons(x)
}

func first(x interface{}) interface{} {
	var seq ISeq = seq(x)
	if seq == nil {
		return nil
	}
	return seq.first()
}

func next(x interface{}) ISeq {
	var seq ISeq = seq(x)
	if seq == nil {
		return nil
	}
	return seq.next()
}

func get(coll, key interface{}) interface{} {
	//
	return getFrom(coll, key)
}

func getFrom(coll, key interface{}) interface{} {
	if coll == nil {
		return nil
	}
	//
	return nil
}

func assoc(coll, key, val interface{}) Associative {
	if coll == nil {
		return PersistentArrayMap{array: []interface{}{key, val}}
	}
	return coll.(Associative).assoc(key, val)
}

func RT_map(init []interface{}) IPersistentMap {
	if init == nil {
		return PersistentArrayMap{}
	} else if len(init) <= HASHTABLE_THRESHOLD {
		return PersistentArrayMap{}.createWithCheck(init)
	}
	return PersistentHashMap{}.createWithCheck(init)
}

func print(x interface{}) string {
	switch x.(type) {
	case nil:
		return "nil"
	case string:
		return x.(string)
	case ISeq:
		return "ISeq"
	default:
		return "Unknown"
	}
}
