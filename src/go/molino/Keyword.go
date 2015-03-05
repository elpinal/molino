package molino

type Keyword struct {
	sym  Symbol
	_str string
}

func (k Keyword) intern(sym Symbol) Keyword {
	var kw Keyword = Keyword{sym: sym}
	//
	return kw
}

func (k Keyword) internFromString(nsname string) Keyword {
	return k.intern(intern(nsname))
}

func (k Keyword) String() string {
	if k._str == "" {
		k._str = ":" + k.sym.String()
	}
	return k._str
}

func (k Keyword) hashCode() int {
	var ret int64 = int64(Util.hash(k.sym.String())) + int64(0x9e3779b9)
	if (ret >> 31) != 0 {
		return -int(^ret) - 1
	} else {
		return int(ret)
	}
}
