package lang

import (
	"testing"
)

func TestVarIntern(t *testing.T) {
	var v Var
	var ns Namespace = FindOrCreate(intern("testnamespace"))
	var sym Symbol = intern("testsymbol")

	var ret Var = v.intern(ns, sym, 4, true)	
	if ret.root != 4 {
		t.Errorf("Expect: %s, but %s", ret.root, 4)
	}
}

func TestEmptyList(t *testing.T) {
	var e EmptyList
	var _ ISeq = e
	var _ IPersistentList = e
}

func TestPersistentList(t *testing.T) {
	var l PersistentList
	//var _ IPersistentList = l
	var _ ISeq = l
	//x := l.equiv(l)
	/*
	if true {
		t.Errorf("%#v", x)
	}
*/

}

func TestASeq(t *testing.T) {
	var a ASeq
	var _ Obj = a.Obj
	var _ ISeq = a
}

func TestList(t *testing.T) {
	var i []interface{}
	var _ List = i

	var x interface{} = []interface{}{}
	var _ List = x.([]interface{})

	var l List
	var _ interface{} = l
	var _ Iterable = l
	var _ Iterator = l.iterator()
}
