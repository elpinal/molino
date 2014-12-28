package lang

type Keyword struct {
	sym  Symbol
	_str string
}

func (k Keyword) intern(sym Symbol) Keyword {
	var kw Keyword = Keyword{sym: sym}
	//
	return kw
}

func (k Keyword) String() string {
	if k._str == "" {
		k._str = ":" + k.sym.String()
	}
	return k._str
}
