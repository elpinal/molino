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

  VarDefStatement struct {
    VarName string
    Expr    Expression
  }

  IfStatement struct {
    Expr Expression 
    True Expression 
  }
)

func (x *ExpressionStatement) statement() {}
func (x *VarDefStatement)     statement() {}
func (x *IfStatement)         statement() {}

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

  UnaryMinusExpression struct {
    SubExpr Expression
  }

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

  CallExpression struct {
    Expr Expression
    Args []Expression
  }

  BinOpExpression struct {
    HS       []Expression
    Operator int
  }
)

func (x *NumberExpression)       expression() {}
func (x *IdentifierExpression)   expression() {}
func (x *UnaryMinusExpression)   expression() {}
func (x *UnaryKeywordExpression) expression() {}
func (x *StringExpression)       expression() {}
func (x *VectorExpression)       expression() {}
func (x *MapExpression)          expression() {}
func (x *FnExpression)           expression() {}
func (x *CallExpression)         expression() {}
func (x *BinOpExpression)        expression() {}
func (x *BoolExpression)         expression() {}
