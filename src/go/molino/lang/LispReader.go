package lang

import (
	"regexp"
	"strconv"
)

var QUOTE Symbol = intern("quote")

var macros = map[rune]ReaderFn {
	'"': StringReader{},
	';': CommentReader{},
	//'\'': WrappingReader{QUOTE},
	//'@':  WrappingReader{DEREF},
	//'^':  MetaReader{},
	//'`':  SyntaxQuoteReader{},
	//'~':  UnquoteReader{},
	'(': ListReader{},
	')': UnmatchedDelimiterReader{},
	//'[': VectorReader{},
	//']': UnmatchedDelimiterReader{},
	//'{': MapReader{},
	//'}': UnmatchedDelimiterReader{},
	//'\\': CharacterReader{},
	//'%':  ArgReader{},
	//'#':  DispatchReader{},
}

type StringReader struct {
}
type CommentReader struct {
}
type UnmatchedDelimiterReader struct {
}
type ListReader struct {
}

var symbolPat *regexp.Regexp = regexp.MustCompile("^[:]?([^/0-9].*/)?(/|[^/0-9][^/]*)$")
var intPat *regexp.Regexp = regexp.MustCompile("^([-+]?)(?:(0)|([1-9][0-9]*)|0[xX]([0-9A-Fa-f]+)|0([0-7]+)|([1-9][0-9]?)[rR]([0-9A-Za-z]+)|0[0-9]+)$")

type Position struct {
	Line   int
	Column int
}

type Reader struct {
	src      []rune // source
	offset   int    //
	lineHead int    //
	line     int    //
}

func (r *Reader) Init(src string) {
	r.src = []rune(src)
}

func (r *Reader) Read() (interface{}, bool) { //(tok int, lit string, pos Position)
	for {
		//pos = r.position()
		ch := r.read()
		for isWhitespace(ch) {
			ch = r.read()
		}
		if ch == -1 {
			return ch, true
		}
		if isDigit(ch) {
			var n int64 = r.readNumber(ch)
			return n, false
		}
		macroFn, ismacro := getMacro(ch)
		if ismacro {
			ret := macroFn.invoke(r, ch)
			//
			if ret == r {
				continue
			}
			return ret, false
		}
		if ch == '+' || ch == '-' {
			ch2 := r.read()
			if isDigit(ch2) {
				r.unread()
				var n int64 = r.readNumber(ch)
				return n, false
			}
			r.unread()
		}
		var token string = r.readToken(ch)
		return interpretToken(token), false
	}
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == ','
}

func getMacro(ch rune) (ReaderFn, bool) {
	m, n := macros[ch]
	return m, n
}

func isMacro(ch rune) bool {
	_, n := macros[ch]
	return n
}

func isTerminatingMacro(ch rune) bool {
	return ch != '#' && ch != '\'' && ch != '%' && isMacro(ch)
}

func (r *Reader) read() rune {
	if !r.reachEOF() {
		ch := r.src[r.offset]
		r.offset++
		if ch == '\n' {
			r.lineHead = r.offset
			r.line++
		}
		return ch
	}
	r.offset++
	return -1
}

func (r *Reader) unread() {
	r.offset--
}

func (r *Reader) reachEOF() bool {
	return len(r.src) <= r.offset
}

func (r *Reader) position() Position {
	return Position{Line: r.line + 1, Column: r.offset - r.lineHead + 1}
}

func (r *Reader) readToken(initch rune) string {
	var ret []rune = []rune{initch}
	for {
		if ch := r.read(); ch == -1 || isWhitespace(ch) || isTerminatingMacro(ch) {
			r.unread()
			return string(ret)
		} else {
			ret = append(ret, ch)
		}
	}
}

func (r *Reader) readNumber(initch rune) int64 {
	var ret []rune = []rune{initch}
	for {
		if ch := r.read(); ch == -1 || isWhitespace(ch) || isMacro(ch) {
			r.unread()
			break
		} else {
			ret = append(ret, ch)
		}
	}
	n, notmatch := matchNumber(string(ret))
	if !notmatch {
		panic("Invalid number: " + string(ret))
	}
	return n //strconv.FormatInt(n, 10)
}

func matchNumber(s string) (int64, bool) {
	m := intPat.FindStringSubmatch(s)
	if m != nil {
		if m[2] != "" {
			return 0, true
		}
		var negate bool = m[1] == "-"
		radix := 10
		var n string
		if n = m[3]; n != "" {
			radix = 10
		} else if n = m[4]; n != "" {
			radix = 16
		}
		if n == "" {
			return -1, false
		}
		ret, err := strconv.ParseInt(n, radix, 64)
		if err != nil {
			panic("error!")
		}
		if negate {
			ret = -ret
		}
		return ret, true
	}
	return -1, false
}

func interpretToken(s string) interface{} {
	if s == "nil" {
		return nil
	} else if s == "true" {
		return true
	} else if s == "false" {
		return false
	}
	ret, isValid := matchSymbol(s)
	if isValid {
		return ret
	}

	panic("Invalid token: " + s)
}

func matchSymbol(s string) (interface{}, bool) {
	m := symbolPat.FindStringSubmatch(s)
	if m != nil {
		isKeyword := s[0] == ':'
		if isKeyword {
			sym := intern(s[1:])
			return /*Keyword.intern(s)*/ sym, true
		} else {
			sym := intern(s)
			return sym, true
		}
	}
	return nil, false
}


func (f StringReader) invoke(r *Reader, doublequote rune) interface{} {
	var ret []rune
	for ch := r.read(); ch != '"'; ch = r.read() {
		if ch == -1 {
			panic("EOF while reading string")
		}
		if ch == '\\' {
			ch = r.read()
			if ch == -1 {
				panic("EOF while reading string")
			}
			switch ch {
			case 't':
				ch = '\t'
			case 'r':
				ch = '\r'
			case 'n':
				ch = '\n'
			case '\\':
			case '"':
			case 'b':
				ch = '\b'
			case 'f':
				ch = '\f'
			case 'u':
				ch = r.read()
				if !( ( '0' <= ch && ch <= '9' ) || ( 'a' <= ch && ch <= 'f' ) ) {
					panic("Invalid unicode escape \\u" + string(ch))
				}
				ch = readUnicodeChar(r, ch, 16, 4, true)
			default:
				if isDigit(ch) {
					ch = readUnicodeChar(r, ch, 8, 3, false)
					if ch > 0377{
						panic("Octal escape sequence must be in range [0, 377].")
					}
				} else {
					panic("Unsupported escape character: \\" + string(ch))
				}
			}
		}
		ret = append(ret, ch)
	}
	return string(ret)
}

func (f CommentReader) invoke(r *Reader, semicolon rune) interface{} {
	var ch rune
	for ch != -1 && ch != '\n' && ch != '\r' {
		ch = r.read()
	}
	return r
}

func (f ListReader) invoke(r *Reader, leftparam rune) interface{} {
	var list []interface{} = readDelimitedList(')', r)
	if list == nil {
		return EmptyList{}
	}
	return PersistentList.create(PersistentList{}, list)
}

func (f UnmatchedDelimiterReader) invoke(r *Reader, rightdelim rune) interface{} {
	panic("Unmatched delimiter: " + string(rightdelim))
}

func readUnicodeChar(r *Reader, initch rune, base int, length int, exact bool) rune {
	uc64, err := strconv.ParseInt(string(initch), base, 0)
	if err != nil {
		panic("Invalid digit: " + string(initch))
	}
	uc := int(uc64)
	i := 1
	for ; i < length; i++ {
		ch := r.read()
		if ch == -1 || isWhitespace(ch) || isMacro(ch) {
			r.unread()
			break
		}
		d64, err := strconv.ParseInt(string(ch), base, 32)
		if err != nil {
			panic("Invalid digit: " + string(ch))
		}
		uc = uc * base + int(d64)
	}
	if i != length && exact {
		panic("Invalid character length: " + strconv.Itoa(i) + ", should be: " + strconv.Itoa(length))
	}
	return rune(uc)
}

func readDelimitedList(delim rune, r *Reader) []interface{} {
	var a []interface{}
	for {
		ch := r.read()
		for isWhitespace(ch) {
			ch = r.read()
		}
		if ch == -1 {
			panic("EOF while reading")
		}
		if ch == delim {
			break
		}
		macroFn, ismacro := getMacro(ch)
		if ismacro {
			mret := macroFn.invoke(r, ch)
			if mret != r {
				a = append(a, mret)
			}
		} else {
			r.unread()
			o, _ := r.Read()
			a = append(a, o)
		}
	}
	return a
}
