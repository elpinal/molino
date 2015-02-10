package lang

import (
	"fmt"
)

type Compiler struct {}

type Expr interface {
	eval() interface{}
}

type NilExpr struct {}
type BoolExpr struct {
	val bool
}
type NumberExpr struct {
	n int64
}
type VectorExpr struct {
	args IPersistentVector
}

func eval(form interface{}) interface{} {
	//
	expr := analyze(form)
	return expr.eval()
}

func analyze(form interface{}) Expr {
	//
	if form == nil {
		return NilExpr{}
	} else if form == true {
		return BoolExpr{true}
	} else if form == false {
		return BoolExpr{false}
	}
	switch form.(type) {
//	case Symbol:
//		return analyzeSymbol(form.(Symbol))
	case int64:
		return NumberExpr{form.(int64)}
	case ISeq:
		return analyzeSeq(form.(ISeq))
	case IPersistentVector:
		return VectorExpr{}.parse(form.(IPersistentVector))
	}
	//
	//return nil //
	panic(fmt.Sprintf("Can't analyze: %s", form))
	//
}

func (_ Compiler) load(rdr *Reader) (interface{}, error) {
	var ret interface{}
	for r, eof, err := rdr.Read(); !eof; r, eof, err = rdr.Read() {
		if err != nil {
			return nil, err
		}
		ret = eval(r)
	}
	return ret, nil
}

/*
func analyzeSymbol(sym Symbol) Expr {
	//
	return //
}
*/

func analyzeSeq(form ISeq) Expr {
	op := first(form)
	if op == nil {
		panic("Can't call nil")
	}
	//
	return nil //
	//
}

func (_ NilExpr) eval() interface{} {
	return nil
}

func (e BoolExpr) eval() interface{} {
	if e.val {
		return true
	}
	return false
}

func (e NumberExpr) eval() interface{} {
	return e.n
}

func (_ VectorExpr) parse(form IPersistentVector) Expr {
	var args IPersistentVector = PersistentVector_EMPTY
	for i := 0; i < form.count(); i++ {
		var v Expr = analyze(form.nth(i))
		args = args.cons(v)
	}
	//
	var ret Expr = VectorExpr{args}
	return ret
	//
}

func (e VectorExpr) eval() interface{} {
	var ret IPersistentVector = PersistentVector_EMPTY
	for i := 0; i < e.args.count(); i++ {
		ret = ret.cons(e.args.nth(i).(Expr).eval())
	}
	return ret
	//
}
