package lang

type PersistentList struct {
	ASeq
	_first interface{}
	_rest IPersistentList
	_count int
}

type EmptyList struct {
	Obj
}

/*
func (l PersistentList) String() string {
	return print(l)
}
*/

/*
func (l PersistentList) listprint() string {
	var ret []rune
	ret = append(ret, '(')
	ret = append(ret, ')')
	return string(ret)
}

func (l PersistentList) printInnerSeq() []rune {
	var ret []rune
	ret = append(ret, '(')
	for s := l; s != nil; s = s.next() {
		ret =
	}
	ret = append(ret, ')')
	return string(ret)
}
*/

func (l PersistentList) create(init []interface{}) IPersistentList {
	var ret IPersistentList = EmptyList{}
	for i := len(init) - 1; i >= 0; i-- {
		ret = ret.(ISeq).cons(init[i]).(IPersistentList)
	}
	return ret
}

func (l PersistentList) equiv(obj interface{}) bool {
	//
	switch obj.(type) {
	case Sequential:
	default:
		return false
	}
	var ms ISeq = seq(obj)
	for s := l.seq(); &s != (*ISeq)(nil); s, ms = s.next(), ms.next() {
		if ms == nil {
			return false
		}
	}
	return ms == nil
}

func (l PersistentList) first() interface{} {
	return l._first
}

func (l PersistentList) next() ISeq {
	if l._count == 1 {
		return nil
	}
	return l._rest.(PersistentList)
}

func (l PersistentList) more() ISeq {
	var s = l.next()
	if s == nil {
		return EmptyList{}
	}
	return s
}

func (l PersistentList) count() int {
	return l._count
}

func (l PersistentList) cons(o interface{}) ISeq {
	return PersistentList{_first: o, _rest: l, _count: l._count + 1}
}

func (l PersistentList) empty() IPersistentCollection {
	return EmptyList{}
}

func (l PersistentList) peek() interface{} {
	return l._first
}

func (l PersistentList) pop() IPersistentStack { //IPersistentList
	return l._rest
}

func (l PersistentList) seq() ISeq {
	return l
}


func (e EmptyList) equals(o interface{}) bool {
	return o == nil
}

func (e EmptyList) equiv(o interface{}) bool {
	return e.equals(o)
}

func (e EmptyList) first() interface{} {
	return nil
}
func (e EmptyList) next() ISeq {
	return nil
}
func (e EmptyList) more() ISeq {
	return e
}
func (e EmptyList) cons(o interface{}) ISeq {
	return PersistentList{_first: o, _rest: nil, _count: 1}
}

func (e EmptyList) empty() IPersistentCollection {
	return e
}

func (e EmptyList) peek() interface{} {
	return nil
}

func (e EmptyList) pop() IPersistentStack {
	panic("Can't pop empty list")
}

func (e EmptyList) count() int {
	return 0
}

func (e EmptyList) seq() ISeq {
	return nil
}
