package lang

type Keyword struct {
	sym Symbol
}

func (k Keyword) intern(sym Symbol) Keyword {
	var kw Keyword = Keyword{sym: sym}
	//
	return kw
}
