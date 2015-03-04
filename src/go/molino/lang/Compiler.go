package lang

import (
	"fmt"
	"strings"
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

type DefExpr struct {
	v            Var
	init         Expr
	meta         Expr
	initProvided bool
	isDynamic    bool
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
	DEF          Symbol  = intern("def")
	LOCAL_ENV    Var     = Var{}
	CONSTANTS    Var     = Var{}.create()
	CONSTANT_IDS Var     = Var{}.create()
	KEYWORDS     Var     = Var{}.create()
	VARS         Var     = Var{}.create()
	arglistsKey  Keyword = Keyword{}.internFromString("arglists")
	dynamicKey   Keyword = Keyword{}.internFromString("dynamic")
	NS           Symbol  = intern("ns")
	IN_NS        Symbol  = intern("in-ns")
)

var specials IPersistentMap = PersistentHashMap{}.create(
	DEF, DefExpr{},
	QUOTE, ConstantExpr{},
)

func eval(form interface{}) interface{} {
	//
	expr := analyze(form)
	return expr.eval()
}

func analyze(form interface{}) Expr {
	return analyze1(form, "")
}

func analyze1(form interface{}, name string) Expr {
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
		return analyzeSeq(form.(ISeq), name)
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
	defer Var{}.popThreadBinding()
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

func analyzeSeq(form ISeq, name string) Expr {
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

func (e DefExpr) parse(form interface{}) Expr {
	var docstring string
	if count(form) == 4 {
		if s, ok := third(form).(string); ok {
			docstring = s
			form = list(first(form), second(form), fourth(form))
		}
	}
	if count(form) > 3 {
		panic("Too many arguments to def")
	} else if count(form) < 2 {
		panic("Too few arguments to def")
	}
	sym, ok := second(form).(Symbol)
	if !ok {
		panic("First argument to def must be a Symbol")
	}
	v, ok := lookupVar(sym, true, true)
	if !ok {
		panic("Can't refer to qualified var that doesn't exist")
	}
	if v.ns.String() != currentNS().String() {
		if sym.ns == "" {
			v = currentNS().intern(sym)
			registerVar(v)
		} else {
			panic("Can't create defs outside of current ns")
		}
	}
	var mm IPersistentMap = sym.meta()
	var isDynamic bool = booleanCast(get(mm, dynamicKey))
	if isDynamic {
		v.setDynamic()
	} else if !isDynamic && strings.HasPrefix(sym.name, "*") && strings.HasSuffix(sym.name, "*") && len(sym.name) > 2 {
		panic(fmt.Sprintf("Warning: %v not declared dynamic and thus is not dynamically rebindable, "+
			"but its name suggests otherwise. Please either indicate ^:dynamic %v or change the name.",
			sym, sym))
	}
	if booleanCast(get(mm, arglistsKey)) {
		var vm IPersistentMap = v.meta()
		vm = assoc(vm, arglistsKey, second(mm.valAt(arglistsKey)))
		v.setMeta(vm)
	}
	//
	if docstring != "" {
		mm = assoc(mm, DOC_KEY, docstring)
	}
	//
	var meta Expr
	if mm.count() != 0 {
		meta = analyze(mm)
	}
	return DefExpr{v: v, init: analyze1(third(form), v.sym.name), meta: meta, initProvided: count(form) == 3, isDynamic: isDynamic}
}

func (e DefExpr) eval() interface{} {
	if e.initProvided {
		e.v.bindRoot(e.init.eval())
	}
	if e.meta != nil {
		var metaMap IPersistentMap = e.meta.eval().(IPersistentMap)
		e.v.setMeta(metaMap)
	}
	return e.v.setDynamicTo(e.isDynamic)
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

func lookupVar(sym Symbol, internNew, registerMacro bool) (Var, bool) {
	var v Var
	var ok bool = true
	if sym.ns != "" {
		var ns Namespace = namespaceFor(currentNS(), sym)
		if ns.name.name == "" {
			return Var{}, false
		}
		var name Symbol = intern(sym.name)
		if internNew && (ns.String() == currentNS().String()) {
			v = currentNS().intern(name)
		} else {
			v, ok = ns.findInternedVar(name)
		}
	} else if sym == NS {
		v = NS_VAR
	} else if sym == IN_NS {
		v = IN_NS_VAR
	} else {
		o, ok := currentNS().getMapping(sym)
		if !ok {
			if internNew {
				v = currentNS().intern(intern(sym.name))
			}
			//} else if ov, ok := o.(Var); ok {
			//v = ov
		} else {
			v = o
			//panic(fmt.Sprintf("Expecting var, but %v is mapped to %v", sym, o))
		}
	}
	if !v.isMacro() || registerMacro {
		registerVar(v)
	}
	return v, ok
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
