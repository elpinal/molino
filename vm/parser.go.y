%{
package vm

import (
  "log"
)

type Token struct {
  tok int
  lit string
  pos Position
}

type Fn struct {
  Args  Args
  Exprs []Expression
}

type Args struct {
  Args   []string
  Vararg bool
  More   string
}

%}

%union{
  statements []Statement
  statement  Statement
  exprs      []Expression
  expr       Expression
  tok        Token
  expr_pairs map[Expression]Expression
  fn         Fn
  fns        []Fn
  args       Args
  idents     []string
  bool       bool
}

%type<statements> statements
%type<statement>  statement
%type<exprs>      exprs
%type<expr>       expr
%type<expr_pairs> expr_pairs
%type<fn>         fn
%type<fns>        fns
%type<args>       args
%type<idents>     idents
%type<bool>       bool

%token<tok> IDENT NUMBER KEYWORD STRING DEF IF TRUE FALSE NIL FN QUOTE

%left '+' '-'
%left '*' '/' '%'
%right UNARY

%%

statements
  :
  {
    $$ = nil
    if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
      l.statements = $$
    }
  }
  | statement statements
  {
    $$ = append([]Statement{$1}, $2...)
    if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
      l.statements = $$
    }
  }

statement
  : expr
  {
    $$ = &ExpressionStatement{Expr: $1}
  }

exprs
  : 
  {
    $$ = []Expression{}
  }
  | expr exprs
  {
    $$ = append([]Expression{$1}, $2...)
  }

expr  : NUMBER
  {
    $$ = &NumberExpression{Lit: $1.lit}
  }
  | IDENT
  {
    $$ = &IdentifierExpression{Lit: $1.lit}
  }
  | bool
  {
    $$ = &BoolExpression{Bool: $1}
  }
  | NIL
  {
    $$ = &NilExpression{}
  }
  | '-' expr      %prec UNARY
  {
    $$ = &UnaryMinusExpression{SubExpr: $2}
  }
  | KEYWORD
  {
    $$ = &UnaryKeywordExpression{Lit: $1.lit}
  }
  | STRING
  {
    $$ = &StringExpression{Lit: $1.lit}
  }
  | '[' exprs ']'
  {
    $$ = &VectorExpression{Exprs: $2}
  }
  | '{' expr_pairs '}'
  {
    $$ = &MapExpression{Map: $2}
  }
  | '(' FN fns ')'
  {
    $$ = &FnExpression{Fns: $3}
  }
  | '(' FN args exprs ')'
  {
    $$ = &FnExpression{Fns: []Fn{Fn{Args: $3, Exprs: $4}}}
  }
  | '(' expr exprs ')'
  {
    $$ = &CallExpression{Expr: $2, Args: $3}
  }
  | '(' DEF IDENT expr ')'
  {
    $$ = &DefExpression{VarName: $3.lit, Expr: $4}
  }
  | '(' IF expr expr ')'
  {
    $$ = &IfExpression{Expr: $3, True: $4}
  }
  | '(' IF expr expr expr ')'
  {
    $$ = &IfExpression{Expr: $3, True: $4, False: $5}
  }
  | '(' QUOTE expr exprs ')'
  {
    $$ = &ConstantExpression{Expr: $3}
  }
  | '(' '=' exprs ')'
  {
    $$ = &EqualExpression{HS: $3}
  }
  | '(' '+' exprs ')'
  { $$ = &BinOpExpression{HS: $3, Operator: int('+')} }
  | '(' '-' exprs ')'
  { $$ = &BinOpExpression{HS: $3, Operator: int('-')} }
  | '(' '*' exprs ')'
  { $$ = &BinOpExpression{HS: $3, Operator: int('*')} }
  | '(' '/' exprs ')'
  { $$ = &BinOpExpression{HS: $3, Operator: int('/')} }
  | '(' '%' exprs ')'
  { $$ = &BinOpExpression{HS: $3, Operator: int('%')} }

expr_pairs
  :
  {
    $$ = map[Expression]Expression{}
  }
  | expr_pairs expr expr
  {
    $$ = $1
    $$[$2] = $3
  }

fn
  : '(' args exprs ')'
  {
    $$ = Fn{Args: $2, Exprs: $3}
  }

fns
  : fn
  {
    $$ = []Fn{$1}
  }
  | fns fn
  {
    $$ = append($1, $2)
  }

args
  : '[' idents ']'
  { $$ = Args{Args: $2} }
  | '[' idents '&' IDENT ']'
  { $$ = Args{Args: $2, Vararg: true, More: $4.lit} }

idents
  : { $$ = []string{} }
  | idents IDENT
  { $$ = append($1, $2.lit) }

bool
  : TRUE
  { $$ = true  }
  | FALSE
  { $$ = false }

%%

type LexerWrapper struct {
  s          *Scanner
  recentLit  string
  recentPos  Position
  statements []Statement
}

func (l *LexerWrapper) Lex(lval *yySymType) int {
  tok, lit, pos := l.s.Scan()
  if tok == EOF {
    return 0
  }
  lval.tok = Token{tok: tok, lit: lit, pos: pos}
  l.recentLit = lit
  l.recentPos = pos
  return tok
}

func (l *LexerWrapper) Error(e string) {
  log.Fatalf("Line %d, Column %d: %q %s",
    l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(s *Scanner) []Statement {
  l := LexerWrapper{s: s}
  if yyParse(&l) != 0 {
    panic("Parse error")
  }
  return l.statements
}
