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
			//
		}
		ret = append(ret, ch)
	}
	return string(ret)
}
