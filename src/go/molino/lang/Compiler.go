package lang

import (
	"fmt"
)

type Compiler struct{}

type LocalBinding struct {
	idx   int
	sym   Symbol
	tag   Symbol
	init  Expr
	isArg bool
}

type IParser interface {
	parse(interface{}) Expr
}

type Expr interface {
	eval() interface{}
}
type LiteralExpr interface {
	Expr
	val() interface{}
}

type NilExpr struct{}
type BoolExpr struct {
	val bool
}
type NumberExpr struct {
	n int64
}
type ConstantExpr struct {
	v  interface{}
	id int
}
type StringExpr struct {
	str string
}
type KeywordExpr struct {
	k Keyword
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
	b   LocalBinding
	tag Symbol
}
type InvokeExpr struct {
	fexpr Expr
	args  IPersistentVector
}

var (
	LOCAL_ENV    Var    = Var{}
	CONSTANTS    Var    = Var{}.create()
	CONSTANT_IDS Var    = Var{}.create()
	KEYWORDS     Var    = Var{}.create()
	VARS         Var    = Var{}.create()
	NS           Symbol = intern("ns")
	IN_NS        Symbol = intern("in-ns")
)

var specials IPersistentMap = PersistentHashMap{}.create(
	QUOTE, ConstantExpr{},
)

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
	case Keyword:
		return registerKeyword(form.(Keyword))
	case int64:
		return NumberExpr{form.(int64)}
	case string:
		return StringExpr{form.(string)} //.intern()
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
	Var{}.pushThreadBinding(mapUniqueKeys(&CURRENT_NS, CURRENT_NS.deref()))
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
		var b, ok = referenceLocal(sym)
		if ok {
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
	if p := specials.valAt(op); p != nil {
		return p.(IParser).parse(form)
	}
	return InvokeExpr{}.parse(form)
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

func (e ConstantExpr) eval() interface{} {
	return e.v
}

func (e ConstantExpr) parse(form interface{}) Expr {
	v := second(form)
	switch v.(type) {
	case nil:
		return NilExpr{}
	case bool:
		return BoolExpr{v.(bool)}
	case int64:
		return NumberExpr{v.(int64)}
	case string:
		return StringExpr{v.(string)}
		//case IPersistentCollection :
	}
	return ConstantExpr{v: v, id: registerConstant(v)}
}

func (e StringExpr) eval() interface{} {
	return e.str
}

func (e KeywordExpr) eval() interface{} {
	return e.k
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
	var keysConstant, valsConstant bool = false, false
	//
	for s := seq(form); s != nil; s = s.next() {
		var e IMapEntry = s.first().(IMapEntry)
		var k Expr = analyze(e.key())
		var v Expr = analyze(e.val())
		keyvals = keyvals.cons(k)
		keyvals = keyvals.cons(v)
		// FIXME:
		//if _, ok := k.(LiteralExpr); ok {
		//}
		//
	}
	var ret Expr = MapExpr{keyvals}
	//
	if keysConstant {
		if valsConstant {
			var m IPersistentMap = PersistentHashMap{}
			for i := 0; i < keyvals.length(); i += 2 {
				m = m.assoc(keyvals.nth(i).(Expr).eval(), keyvals.nth(i+1).(Expr).eval())
			}
			return ConstantExpr{v: m, id: registerConstant(m)}
		}
	}
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

func (e InvokeExpr) parse(form ISeq) Expr {
	var fexpr Expr = analyze(form.first())
	//
	//
	var args PersistentVector = PersistentVector_EMPTY
	for s := seq(form.next()); s != nil; s = s.next() {
		args = args.cons(analyze(s.first())).(PersistentVector)
	}
	return InvokeExpr{fexpr: fexpr, args: args}
}

func (e InvokeExpr) eval() interface{} {
	var fn IFn = e.fexpr.eval().(IFn)
	var argvs IPersistentVector = PersistentVector_EMPTY
	for i := 0; i < e.args.count(); i++ {
		argvs = argvs.cons(e.args.nth(i).(Expr).eval()).(PersistentVector)
	}
	return fn.applyTo(seq(Util.ret1(argvs.(Seqable).seq(), nil)))
}

func referenceLocal(sym Symbol) (LocalBinding, bool) {
	if !LOCAL_ENV.isBound() {
		return LocalBinding{}, false
	}
	var b = get(LOCAL_ENV.deref(), sym)
	if b != nil {
		//
		panic("FIXME: referenceLocal")
	}
	return LocalBinding{}, false
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
		v, ok := ns.findInternedVar(intern(sym.name))
		if !ok {
			panic("No such var: " + sym.String())
		} //
		return v
	} else if sym == NS {
		return NS_VAR
	} else if sym == IN_NS {
		return IN_NS_VAR
	} else {
		o, ok := n.getMapping(sym)
		if !ok {
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

func registerConstant(o interface{}) int {
	if !CONSTANTS.isBound() {
		return -1
	}
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

func registerKeyword(keyword Keyword) KeywordExpr {
	if !KEYWORDS.isBound() {
		return KeywordExpr{k: keyword}
	}
	var keywordsMap IPersistentMap = KEYWORDS.deref().(IPersistentMap)
	var id = get(keywordsMap, keyword)
	if id == nil {
		KEYWORDS.set(assoc(keywordsMap, keyword, registerConstant(keyword)))
	}
	return KeywordExpr{k: keyword}
}

func registerVar(v Var) {
	if !VARS.isBound() {
		return
	}
	var varsMap IPersistentMap = VARS.deref().(IPersistentMap)
	id := get(varsMap, v)
	if id == nil {
		VARS.set(assoc(varsMap, v, registerConstant(v)))
	}
}
