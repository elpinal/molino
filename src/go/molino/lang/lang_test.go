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

func TestPrint(t *testing.T) {
	var s Symbol = intern("test")
	if s.String() != "test" {
		t.Errorf("\"%s\" (type Symbol) must be \"test\"\n", s)
	}

	var n Symbol = intern("London/Prelude")
	if n.String() != "London/Prelude" {
		t.Errorf("\"%s\" (type Symbol) must be \"London/Prelude\"\n", n)
	}

	var k Keyword = Keyword.intern(Keyword{}, s)
	if k.String() != ":test" {
		t.Errorf("\"%s\" (type Keyword) must be \":test\"\n", k)
	}

	var l Keyword = Keyword.intern(Keyword{}, n)
	if l.String() != ":London/Prelude" {
		t.Errorf("\"%s\" (type Keyword) must be \":London/Prelude\"\n", l)
	}
}

func TestPersistentArrayMap(t *testing.T) {
	s1, s2, s3, s4 := intern(":k1"), intern(":k2"), intern(":k3"), intern(":k4")
	k1, k2, k3, k4 :=
		Keyword.intern(Keyword{}, s1), Keyword.intern(Keyword{}, s2), Keyword.intern(Keyword{}, s3), Keyword.intern(Keyword{}, s4)

	var i = []interface{}{k1, 9001, k2, 9002, k3, 9003}
	m := PersistentArrayMap.createWithCheck(PersistentArrayMap{}, i)
	am1, am2, am3 :=
		m.assoc(k1, 9001).(PersistentArrayMap), m.assoc(k1, 10001).(PersistentArrayMap), m.assoc(k4, 9004).(PersistentArrayMap)

	if am1.array[1] != 9001 {
		t.Errorf("Expect 9001, but %v\n", am1.array[1])
	}
	if am2.array[1] != 10001 {
		t.Errorf("Expect 10001, but %v\n", am2.array[1])
	}
	if am3.array[3] != 9001 {
		t.Errorf("Expect 9001, but %v\n", am3.array[3])
	}
	if m.array[1] != 9001 {
		t.Errorf("Expect 9001, but %v\n", m.array[1])
	}
}

func TestBitCount(t *testing.T) {
	var i = make([]int, 12)

	i = []int{0, 1, 2, 10, 64, 99, 100, 999, 9000, 39201, 666666, -90000000}
	for _, n := range i {
		r := bitCount(n)
		switch {
		case n == i[0] && r != 0:
			t.Errorf("Expect 0, but %v\n", r)
		case n == i[1] && r != 1:
			t.Errorf("Expect 1, but %v\n", r)
		case n == i[2] && r != 1:
			t.Errorf("Expect 1, but %v\n", r)
		case n == i[3] && r != 2:
			t.Errorf("Expect 2, but %v\n", r)
		case n == i[4] && r != 1:
			t.Errorf("Expect 1, but %v\n", r)
		case n == i[5] && r != 4:
			t.Errorf("Expect 4, but %v\n", r)
		case n == i[6] && r != 3:
			t.Errorf("Expect 3, but %v\n", r)
		case n == i[7] && r != 8:
			t.Errorf("Expect 8, but %v\n", r)
		case n == i[8] && r != 5:
			t.Errorf("Expect 5, but %v\n", r)
		case n == i[9] && r != 6:
			t.Errorf("Expect 6, but %v\n", r)
		case n == i[10] && r != 8:
			t.Errorf("Expect 8, but %v\n", r)
		case n == i[11] && r != 15:
			t.Errorf("Expect 15, but %v\n", r)
		}
	}
}

func TestPersistentVector(t *testing.T) {
	var a []interface{} = []interface{}{1, 2, 4, 8, 16, 32, 64, 128}
	result := LazilyPersistentVector{}.create(a)
	if result.(PersistentVector).cnt != 8 {
		t.Errorf("%v: %v should be 8", result.(PersistentVector), result.(PersistentVector).cnt)
	}
}

func BenchmarkPersistentVector(b *testing.B) {
	var a []interface{} = []interface{}{1, 2, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x ISeq = seq(a)
		var y = PersistentVector.create(PersistentVector{}, x)
		var _ = y
	}
}
