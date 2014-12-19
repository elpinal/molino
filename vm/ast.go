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

  VectorExpression struct {
    Exprs []Expression
  }

  MapExpression struct {
    Map map[Expression]Expression
  }

  CallExpression struct {
    Expr Expression
    Args []Expression
  }
)

func (x *NumberExpression)       expression() {}
func (x *IdentifierExpression)   expression() {}
func (x *VectorExpression)       expression() {}
func (x *MapExpression)          expression() {}
func (x *CallExpression)         expression() {}
