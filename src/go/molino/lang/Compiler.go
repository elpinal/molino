package lang

type Expr interface {
	eval() interface{}
}

func eval(form interface{}) interface{} {
	//
	expr := analyze(form)
	return expr.eval()
}

func analyze(form interface{}) Expr {
	//
	if form == nil {
		return NIL_EXPR
	} else if form == true {
		return TRUE_EXPR
	} else if form == false {
		return FALSE_EXPR
	}
	//
	return nil //
	//
}
