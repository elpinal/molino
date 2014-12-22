package lang

import (
	"strconv"
)

type Fn interface {
	invoke(*Reader, rune) interface{}
}

type StringReader struct {
}

func (f StringReader) invoke(r *Reader,doublequote rune) interface{} {
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
