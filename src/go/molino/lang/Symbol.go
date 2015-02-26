package lang

import (
	"strings"
)

type Symbol struct {
	ns      string
	name    string
	_hasheq int
	_str    string
}

func (s Symbol) String() string {
	if s._str == "" {
		if s.ns != "" {
			s._str = s.ns + "/" + s.name
		} else {
			s._str = s.name
		}
	}
	return s._str
}

func intern(nsname string) Symbol {
	var i int = strings.Index(nsname, "/")
	if i == -1 || strings.EqualFold(nsname, "/") {
		return Symbol{name: nsname}
	} else {
		ns := nsname[0:i]
		name := nsname[i+1:]
		return Symbol{ns: ns, name: name}
	}
}

func (s Symbol) hasheq() int {
	if s._hasheq == 0 {
		s._hasheq = Util.hashCombine(hashUnencodedChars(s.name), Util.hash(s.ns))
	}
	return s._hasheq
}
