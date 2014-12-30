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
	reader := new(Reader)
	reader.Init(string(body))
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
	case Iterable:
		ret, _ := IteratorSeq.create(IteratorSeq{}, coll.(Iterable).iterator())
		return ret
	case []interface{}:
		return createFromObject(coll.([]interface{}))
	}
	//
	panic("Don't know how to create ISeq from: " + fmt.Sprintf("%T\n", coll))
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


func RT_map(init []interface{}) IPersistentMap {
	if init == nil {
		return PersistentArrayMap{}
	} else if len(init) <= HASHTABLE_THRESHOLD {
		return PersistentArrayMap.createWithCheck(PersistentArrayMap{}, init)
	}
	return PersistentHashMap.createWithCheck(PersistentHashMap{}, init)
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
