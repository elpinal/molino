package lang

type Namespace struct {
	name     Symbol
	mappings map[Symbol]Var
}

//var mappings   = make(map[Symbol]Var)
//  var aliases    =

var namespaces = make(map[Symbol]Namespace)

func (n Namespace) String() string {
	return n.name.String()
}

func FindOrCreate(name Symbol) Namespace {
	ns, ok := namespaces[name]
	if ok {
		return ns
	}
	var newns Namespace = Namespace{name: name, mappings: make(map[Symbol]Var)}
	namespaces[name] = newns
	return newns
}

func find(name Symbol) Namespace {
	return namespaces[name]
}

func (n Namespace) intern(sym Symbol) Var {
	if sym.ns != "" {
		panic("Can't intern namespace-qualified symbol")
	}
	a, ok := n.mappings[sym]
	if ok {
		return a
	}
	var v Var = Var{ns: n, sym: sym}
	unbound := Unbound{v}
	v.root = unbound
	n.mappings[sym] = v
	return v
}

func (n Namespace) refer(sym Symbol, v Var) Var {
	if sym.ns != "" {
		panic("ns is not empty!")
	}
	a, ok := n.mappings[sym]
	if ok {
		return a
	}
	n.mappings[sym] = v
	return v
}

func (n Namespace) getMapping(name Symbol) (Var, bool) {
	v, ok := n.mappings[name]
	return v, ok
}

func (n Namespace) updateMapping(name Symbol, newval Var) {
	if _, ok := n.mappings[name]; ok {
		n.mappings[name] = newval
	}
}

func findNamespace(name Symbol) (Namespace, bool) {
	v, ok := namespaces[name]
	return v, ok
}

func (n Namespace) findInternedVar(sym Symbol) (Var, bool) {
	v, ok := n.mappings[sym]
	if ok && v.ns.name == n.name {
		return v, true
	}
	return v, false
}
