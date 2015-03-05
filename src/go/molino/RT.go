package molino

import (
	"fmt"
	"io/ioutil"
	"log"
)

var (
	MOLINO_NS Namespace = FindOrCreate(intern("molino.core"))

	TAG_KEY Keyword = Keyword{}.internFromString("tag")
	DOC_KEY Keyword = Keyword{}.internFromString("doc")

	IN_NAMESPACE Symbol = intern("in-ns")
	NAMESPACE    Symbol = intern("ns")

	CURRENT_NS Var = Var{}.intern(MOLINO_NS, intern("*ns*"), MOLINO_NS, true)
	NS_VAR     Var = Var{}.intern(MOLINO_NS, intern("ns"), false, true)
	IN_NS_VAR  Var = Var{}.intern(MOLINO_NS, intern("in-ns"), false, true)
)

var inNamespace IFn = NewAFn(func(arg1 interface{}) interface{} {
	var nsname Symbol = arg1.(Symbol)
	var ns Namespace = FindOrCreate(nsname)
	CURRENT_NS.set(ns)
	return ns
})

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
	case nil:
		return nil
	case Iterable:
		ret, _ := IteratorSeq.create(IteratorSeq{}, coll.(Iterable).iterator())
		return ret
	case []interface{}:
		return createFromObject(coll.([]interface{}))
	}
	//
	panic("Don't know how to create ISeq from: " + fmt.Sprintf("%T\n", coll))
}

func meta(x interface{}) IPersistentMap {
	if m, ok := x.(IMeta); ok {
		return m.meta()
	}
	return nil
}

func count(o interface{}) int {
	if c, ok := o.(Counted); ok {
		return c.count()
	}
	return countFrom(o)
}

func countFrom(o interface{}) int {
	switch o.(type) {
	case nil:
		return 0
	case IPersistentCollection:
		i := 0
		for s := seq(o); s != nil; s = s.next() {
			if c, ok := s.(Counted); ok {
				return i + c.count()
			}
			i++
		}
		return i
	}
	panic(fmt.Sprintf("count not supported on this type: %T\n", o))
}

func conj(coll IPersistentCollection, x interface{}) IPersistentCollection {
	if coll == nil {
		return PersistentList{_first: x, _count: 1}
	}
	return coll.(ISeq).cons(x)
}

func cons(x, coll interface{}) ISeq {
	if coll == nil {
		return PersistentList{_first: x, _count: 1}
	} else if c, ok := coll.(ISeq); ok {
		return Cons{_first: x, _more: c}
	}
	return Cons{_first: x, _more: seq(coll)}
}

func first(x interface{}) interface{} {
	var seq ISeq = seq(x)
	if seq == nil {
		return nil
	}
	return seq.first()
}

func second(x interface{}) interface{} {
	return first(next(x))
}

func third(x interface{}) interface{} {
	return first(next(next(x)))
}

func fourth(x interface{}) interface{} {
	return first(next(next(next(x))))
}

func next(x interface{}) ISeq {
	var seq ISeq = seq(x)
	if seq == nil {
		return nil
	}
	return seq.next()
}

func get(coll, key interface{}) interface{} {
	if c, ok := coll.(ILookup); ok {
		return c.valAt(key)
	}
	return getFrom(coll, key)
}

func getFrom(coll, key interface{}) interface{} {
	if coll == nil {
		return nil
	}
	//
	panic(fmt.Sprintf("FIXME: Can't get from %T", coll))
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

func mapUniqueKeys(init ...interface{}) IPersistentMap {
	if init == nil {
		return PersistentArrayMap{}
	} else if len(init) <= HASHTABLE_THRESHOLD {
		return PersistentArrayMap{array: init}
	}
	return PersistentHashMap{}.create(init)
}

func list(arg ...interface{}) ISeq {
	l := len(arg)
	switch l {
	case 0:
		return nil
	case 1:
		return PersistentList{_first: arg[0], _count: 1}
	default:
		return listStar(nil, arg...)
	}
}

func listStar(rest ISeq, arg ...interface{}) ISeq {
	var ret ISeq = rest
	for i := len(arg) - 1; i >= 0; i-- {
		ret = cons(arg[i], ret)
	}
	return ret
}

func length(list ISeq) int {
	i := 0
	for c := list; c != nil; c = c.next() {
		i++
	}
	return i
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

func booleanCast(x interface{}) bool {
	if b, ok := x.(bool); ok {
		return b
	}
	return x != nil
}
