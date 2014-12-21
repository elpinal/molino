package lang

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
			default:
				if isDigit(ch) {
					//ch =
					//if {
					//	panic("Octal escape sequence must be in range [0, 377].")
					//}
				} else {
					panic("Unsupported escape character: \\" + string(ch))
				}
			}
		}
		ret = append(ret, ch)
	}
	return string(ret)
}
