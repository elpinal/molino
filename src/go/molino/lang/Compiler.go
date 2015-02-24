package lang

import (
	"fmt"
)

type Compiler struct {}

type LocalBinding struct {
	idx int
	sym Symbol
	tag Symbol
	init Expr
	isArg bool
}

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
type MapExpr struct {
	keyvals IPersistentVector
}
type VarExpr struct {
	v Var
}
type LocalBindingExpr struct {
	b LocalBinding
	tag Symbol
}

var LOCAL_ENV Var = Var{}
var CONSTANTS Var = Var{}.create()
var CONSTANT_IDS Var = Var{}.create()
var VARS Var = Var{}.create()
var NS Symbol = intern("ns")
var IN_NS Symbol = intern("in-ns")

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
	case Symbol:
		return analyzeSymbol(form.(Symbol))
	case int64:
		return NumberExpr{form.(int64)}
	case ISeq:
		return analyzeSeq(form.(ISeq))
	case IPersistentVector:
		return VectorExpr{}.parse(form.(IPersistentVector))
	case IPersistentMap:
		return MapExpr{}.parse(form.(IPersistentMap))
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

func analyzeSymbol(sym Symbol) Expr {
	var tag Symbol = tagOf(sym)
	if sym.ns == "" {
		var b LocalBinding = referenceLocal(sym)
		if b.sym.name != "" {
			return LocalBindingExpr{b: b, tag: tag}
		}
		//
	} else {
		if namespaceFor(currentNS(), sym).name.name == "" {
			var nsSym Symbol = intern(sym.ns)
			_ = nsSym
			//
		}
	}
	var o interface{} = resolve(sym)
	switch o.(type) {
	case Var:
		var v Var = o.(Var)
		//
		registerVar(v)
		return VarExpr{v}
	case Symbol:
		//
	}
	//
	panic("Unable to resolve symbol: " + sym.String() + " in this context")
}

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

func (_ MapExpr) parse(form IPersistentMap) Expr {
	var keyvals IPersistentVector = PersistentVector_EMPTY
	//
	for s := seq(form); s != nil; s = s.next() {
		var e IMapEntry = s.first().(IMapEntry)
		var k Expr = analyze(e.key())
		var v Expr = analyze(e.val())
		keyvals = keyvals.cons(k)
		keyvals = keyvals.cons(v)
		//
	}
	var ret Expr = MapExpr{keyvals}
	//
	return ret
	//
}

func (e MapExpr) eval() interface{} {
	ret := make([]interface{}, 0, e.keyvals.count())
	for i := 0; i < e.keyvals.count(); i++ {
		ret = append(ret, e.keyvals.nth(i).(Expr).eval())
	}
	return RT_map(ret)
}

func (e VarExpr) eval() interface{} {
	return e.v.deref()
}

func (e LocalBindingExpr) eval() interface{} {
	panic("Can't eval locals")
}

func referenceLocal(sym Symbol) LocalBinding {
	var b LocalBinding = get(LOCAL_ENV.deref(), sym).(LocalBinding)
	if b.sym.name != "" {
		//
	}
	return b
}

func tagOf(o interface{}) Symbol {
	tag := get(meta(o), TAG_KEY)
	switch tag.(type) {
	case Symbol:
		return tag.(Symbol)
	case string:
		return intern(tag.(string))
	}
	return Symbol{}
}

func namespaceFor(inns Namespace, sym Symbol) Namespace {
	var nsSym Symbol = intern(sym.ns)
	var ns Namespace
	//
	ns = find(nsSym)
	//
	return ns
}

func resolve(sym Symbol) interface{} {
	return resolveIn(currentNS(), sym)
}

func resolveIn(n Namespace, sym Symbol) interface{} {
	if sym.ns != "" {
		var ns Namespace = namespaceFor(n, sym)
		if ns.name.name == "" {
			panic("No such namespace: " + sym.ns)
		}
		v, exist := ns.findInternedVar(intern(sym.name))
		if !exist {
			panic("No such var: " + sym.String())
		} //
		return v
	} else if sym == NS {
		return NS_VAR
	} else if sym == IN_NS {
		return IN_NS_VAR
	} else {
		o, exist := n.getmapping(sym)
		if !exist {
			//
			panic("Unable to resolve symbol: " + sym.String() + " in this context")
		}
		return o
	}
	//
}

func currentNS() Namespace {
	return CURRENT_NS.deref().(Namespace)
}

func registerContent(o interface{}) int {
	var v PersistentVector = CONSTANTS.deref().(PersistentVector)
	var ids map[interface{}]int = CONSTANT_IDS.deref().(map[interface{}]int)
	i, ok := ids[o]
	if ok {
		return i
	}
	CONSTANTS.set(conj(v, o))
	ids[o] = v.count()
	return v.count()
}

func registerVar(v Var) {
	if !VARS.isBound() {
		return
	}
	var varsMap IPersistentMap = VARS.deref().(IPersistentMap)
	id := getFrom(varsMap, v)
	if id == nil {
		VARS.set(assoc(varsMap, v, registerContent(v)))
	}
}
