package vm

type (
  Statement interface {
    statement()
  }

  Expression interface {
    expression()
  }
)

type (
  ExpressionStatement struct {
    Expr Expression
  }
)

func (x *ExpressionStatement) statement() {}

type (
  NumberExpression struct {
    Lit string
  }

  IdentifierExpression struct {
    Lit string
  }

  BoolExpression struct {
    Bool bool
  }

  NilExpression struct {
  }

  /*
  UnaryMinusExpression struct {
    SubExpr Expression
  }
  */

  UnaryKeywordExpression struct {
    Lit string
  }

  StringExpression struct {
    Lit string
  }

  VectorExpression struct {
    Exprs []Expression
  }

  MapExpression struct {
    Map map[Expression]Expression
  }

  FnExpression struct {
    Fns []Fn
  }

  DefExpression struct {
    VarName string
    Expr    Expression
  }

  CallExpression struct {
    Expr Expression
    Args []Expression
  }

  IfExpression struct {
    Expr  Expression
    True  Expression
    False Expression
  }

  ConstantExpression struct {
    Expr Expression
  }

/*
  EqualExpression struct {
    HS       []Expression
  }
 */
/*
  BinOpExpression struct {
    HS       []Expression
    Operator int
  }
 */
)

func (x *NumberExpression)       expression() {}
func (x *IdentifierExpression)   expression() {}
//func (x *UnaryMinusExpression)   expression() {}
func (x *UnaryKeywordExpression) expression() {}
func (x *StringExpression)       expression() {}
func (x *VectorExpression)       expression() {}
func (x *MapExpression)          expression() {}
func (x *FnExpression)           expression() {}
func (x *CallExpression)         expression() {}
func (x *DefExpression)          expression() {}
func (x *IfExpression)           expression() {}
//func (x *BinOpExpression)        expression() {}
func (x *BoolExpression)         expression() {}
func (x *NilExpression)          expression() {}
//func (x *EqualExpression)        expression() {}
func (x *ConstantExpression)     expression() {}
