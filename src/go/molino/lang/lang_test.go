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
		t.Errorf(`"%s" (type Symbol) must be "test"`, s)
	}

	var n Symbol = intern("London/Prelude")
	if n.String() != "London/Prelude" {
		t.Errorf(`"%s" (type Symbol) must be "London/Prelude"`, n)
	}

	var k Keyword = Keyword.intern(Keyword{}, s)
	if k.String() != ":test" {
		t.Errorf(`"%s" (type Keyword) must be ":test"`, k)
	}

	var l Keyword = Keyword.intern(Keyword{}, n)
	if l.String() != ":London/Prelude" {
		t.Errorf(`"%s" (type Keyword) must be ":London/Prelude"`, l)
	}

	var v Var = Var{}.intern(MOLINO_NS, intern("test-var"), "hoge", true)
	if v.String() != "#'molino.core/test-var" {
		t.Errorf("%v (type %T) must be #'molino.core/test-var", v, v)
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
		t.Errorf("Expect 9001, but %v", am1.array[1])
	}
	if am2.array[1] != 10001 {
		t.Errorf("Expect 10001, but %v", am2.array[1])
	}
	if am3.array[3] != 9001 {
		t.Errorf("Expect 9001, but %v", am3.array[3])
	}
	if m.array[1] != 9001 {
		t.Errorf("Expect 9001, but %v", m.array[1])
	}
}

func TestBitCount(t *testing.T) {
	var i = make([]int, 12)

	i = []int{0, 1, 2, 10, 64, 99, 100, 999, 9000, 39201, 666666, -90000000}
	for _, n := range i {
		r := bitCount(n)
		switch {
		case n == i[0] && r != 0:
			t.Errorf("Expect 0, but %v", r)
		case n == i[1] && r != 1:
			t.Errorf("Expect 1, but %v", r)
		case n == i[2] && r != 1:
			t.Errorf("Expect 1, but %v", r)
		case n == i[3] && r != 2:
			t.Errorf("Expect 2, but %v", r)
		case n == i[4] && r != 1:
			t.Errorf("Expect 1, but %v", r)
		case n == i[5] && r != 4:
			t.Errorf("Expect 4, but %v", r)
		case n == i[6] && r != 3:
			t.Errorf("Expect 3, but %v", r)
		case n == i[7] && r != 8:
			t.Errorf("Expect 8, but %v", r)
		case n == i[8] && r != 5:
			t.Errorf("Expect 5, but %v", r)
		case n == i[9] && r != 6:
			t.Errorf("Expect 6, but %v", r)
		case n == i[10] && r != 8:
			t.Errorf("Expect 8, but %v", r)
		case n == i[11] && r != 15:
			t.Errorf("Expect 15, but %v", r)
		}
	}
}

func TestPersistentVector(t *testing.T) {
	var a []interface{} = []interface{}{1, 2, 4, 8, 16, 32, 64, 128}
	result := LazilyPersistentVector{}.create(a)
	if result.(PersistentVector).cnt != 8 {
		t.Errorf("%v: %v should be 8", result.(PersistentVector), result.(PersistentVector).cnt)
	}

	f := seq(result).first()
	if f != 1 {
		t.Errorf("%#v should be 1", f)
	}
}

func TestSeqable(t *testing.T) {
	var i = []interface{}{1, 3, 9, 27, 81, 243}
	m := PersistentArrayMap{}.createWithCheck(i)
	var _ Seqable = m
}

func TestMap(t *testing.T) {
	var i = []interface{}{1, 3, 9, 27, 81, 243}
	m := PersistentArrayMap{}.createWithCheck(i)
	var _ Seq = m.seq().(Seq)

	var e MapEntry = m.seq().first().(MapEntry)
	if e.key() != 1 {
		t.Errorf("%v should be 1", e.key())
	}
	if e.val() != 3 {
		t.Errorf("%v should be 3", e.val())
	}

	var es MapEntry = m.seq().next().first().(MapEntry)
	if es.key() != 9 {
		t.Errorf("%v should be 9", es.key())
	}
	if es.val() != 27 {
		t.Errorf("%v should be 27", es.val())
	}

	var ess MapEntry = m.seq().next().next().first().(MapEntry)
	if ess.key() != 81 {
		t.Errorf("%v should be 81", ess.key())
	}
	if ess.val() != 243 {
		t.Errorf("%v should be 243", ess.val())
	}
}

func TestFmix(t *testing.T) {
	a := fmix(12000, 8)
	if a != 1259081277 {
		t.Errorf("%s should be 1259081277", a)
	}
}

func TestMixK1(t *testing.T) {
	a := mixK1(24)
	if a != 166749150 {
		t.Errorf("%s should be 166749150", a)
	}

	b := mixK1(0)
	if b != 0 {
		t.Errorf("%s should be 0", b)
	}
}

func TestMixH1(t *testing.T) {
	a := mixH1(0, 17)
	if a != -429978780 {
		t.Errorf("%s should be -429978780", a)
	}

	b := mixH1(0, 166749150)
	if b != 616509850 {
		t.Errorf("%s should be 616509850", b)
	}

	c := mixH1(616509850, 0)
	if c != 1700053591 {
		t.Errorf("%s should be 1700053591", c)
	}
}

func TestHash(t *testing.T) {
	a := hash(1)
	if a != 1392991556 {
		t.Errorf("%s should be 1392991556", a)
	}

	b := hash(128)
	if b != 292862370 {
		t.Errorf("%s should be 292862370", b)
	}

	c := hash(2147483647)
	if c != 1819228606 {
		t.Errorf("%s should be 1819228606", c)
	}

	//	d := hash(-3)
	//	if d != -1797448787 {
	//		t.Errorf("%s should be -1797448787", d)
	//	}

	e := hash(22)
	if e != 270085581 {
		t.Errorf("%s should be 270085581", e)
	}

	f := hash(64)
	if f != 875635954 {
		t.Errorf("%s should be 875635954", f)
	}

	g := hash(32)
	if g != 483896201 {
		t.Errorf("%s should be 483896201", g)
	}

	h := hash(33)
	if h != 1904410925 {
		t.Errorf("%s should be 1904410925", h)
	}

	i := hash(37)
	if i != -1168698466 {
		t.Errorf("%s should be -1168698466", i)
	}
}

func TestPersistentHashMap(t *testing.T) {
	var i = []interface{}{1, 3, 9, 27, 81, 243}
	m := PersistentHashMap{}.createWithCheck(i)
	var _ IPersistentMap = m
	if !true {
		t.Errorf("%v", m)
	}

	var b ITransientMap = PersistentHashMap{}.asTransient()
	b = b.assoc(5, 10)
	if !true {
		t.Errorf("%v", b)
	}
}

func TestAFn(t *testing.T) {
	var a IFn = NewAFn(func(x int) int { return x*x + 2 })
	v := a.invoke(11)
	if v != 123 {
		t.Errorf("%v should be 123", v)
	}
}

func TestGet(t *testing.T) {
	var i = []interface{}{1, 3, 9, 27, 81, 243}
	am := PersistentArrayMap{}.createWithCheck(i)
	v1 := get(am, 9)
	if v1 != 27 {
		t.Errorf("%v should be 27", v1)
	}

	hm := PersistentHashMap{}.createWithCheck(i)
	v2 := get(hm, 81)
	if v2 != 243 {
		t.Errorf("%v should be 243", v2)
	}
}

func TestSeq(t *testing.T) {
	var i = []interface{}{1, 3, 9, 27, 81, 243, 2, 4, 6, 8, 10, 12, 14, 16}
	hm := PersistentHashMap{}.createWithCheck(i)
	var v ISeq = seq(hm)
	var _ NodeSeq = v.(NodeSeq)
	var f IMapEntry = v.first().(IMapEntry)
	if f.key() != 1 {
		t.Errorf("%v should be 1", f.key())
	}

	var s ISeq = v.next()
	var sf IMapEntry = s.first().(IMapEntry)
	if sf.key() != 6 {
		t.Errorf("%v should be 6", sf.key())
	}

	var ss ISeq = s.next()
	var ssf IMapEntry = ss.first().(IMapEntry)
	if ssf.val() != 4 {
		t.Errorf("%v should be 4", ssf.val())
	}
}

func TestBitmapIndexedNode(t *testing.T) {
	var b INode = BitmapIndexedNode{}
	b1 := b.assocWithEdit(false, 0, hash(1), 1, 5, &Box{})
	b1 = b1.assocWithEdit(false, 0, hash(22), 22, 6, &Box{})
	//t.Errorf("%#v", b1)

	b2 := b.assocWithEdit(false, 0, hash(64), 64, 32, &Box{})
	b2 = b2.assocWithEdit(false, 0, hash(33), 33, 37, &Box{})
	if b2.nodeSeq().first().(IMapEntry).val() != 37 {
		t.Errorf("%v should be 37", b2.nodeSeq().first().(IMapEntry).val())
	}
}

func TestUtilHash(t *testing.T) {
	var s = "mol"
	x := Util.hash(s)
	if x != 108298 {
		t.Errorf("%v should be 108298", x)
	}
}

func TestHashFromVar(t *testing.T) {
	var v Var = Var{}.intern(MOLINO_NS, intern("test-var"), "hoge", true)
	a := hash(v)
	b := hash(v)
	if a != b {
		t.Errorf("%v should equals %v", a, b)
	}
}

func TestHashCode(t *testing.T) {
	var v Var
	var _ IHashCode = v
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
