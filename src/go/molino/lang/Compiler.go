package lang

type Compiler struct {}

type Expr interface {
	eval() interface{}
}

type NilExpr struct {}
type BoolExpr struct {
	val bool
}

func eval(form interface{}) interface{} {
	//
	expr := analyze(form)
	return expr.eval()
}

func analyze(form interface{}) Expr {
	//
	if form == nil {
		return NilExpr{}
	} else if form == true {
		return BoolExpr{true}
	} else if form == false {
		return BoolExpr{false}
	}
	//
	return nil //
	//
}

func (_ Compiler) load(rdr *Reader) (interface{}, error) {
	var ret interface{}
	for r, eof, err := rdr.Read(); !eof; r, eof, err = rdr.Read() {
		if err != nil {
			return nil, err
		}
		ret = eval(r)
	}
	return ret, nil
}


func (_ NilExpr) eval() interface{} {
	return nil
}

func (e BoolExpr) eval() interface{} {
	if e.val {
		return true
	}
	return false
}
