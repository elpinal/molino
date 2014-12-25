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
