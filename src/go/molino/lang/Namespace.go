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

func (this Namespace) intern(sym Symbol) Var {
	if sym.ns != "" {
		panic("Can't intern namespace-qualified symbol")
	}
	a, ok := this.mappings[sym]
	if ok {
		return a
	}
	var v Var = Var{ns: this, sym: sym}
	unbound := Unbound{v}
	v.root = unbound
	this.mappings[sym] = v
	return v
}

func (this Namespace) refer(sym Symbol, v Var) Var {
	if sym.ns != "" {
		panic("ns is not empty!")
	}
	a, ok := this.mappings[sym]
	if ok {
		return a
	}
	this.mappings[sym] = v
	return v
}

func (this Namespace) getmapping(name Symbol) (Var, bool) {
	v, ok := this.mappings[name]
	return v, ok
}

func (this Namespace) updatemapping(name Symbol, newval Var) {
	if _, ok := this.mappings[name]; ok {
		this.mappings[name] = newval
	}
}

func findNamespace(name Symbol) (Namespace, bool) {
	v, ok := namespaces[name]
	return v, ok
}

func (this Namespace) findInternedVar(sym Symbol) (Var, bool) {
	v, ok := this.mappings[sym]
	if ok && v.ns.name == this.name {
		return v, true
	}
	return v, false
}
