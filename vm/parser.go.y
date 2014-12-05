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
  args  []string
  stmts []Statement
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
  args       []string
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

%token<tok> IDENT NUMBER KEYWORD STRING VAR IF TRUE FALSE FN

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
  | '(' VAR IDENT expr ')'
  {
    $$ = &VarDefStatement{VarName: $3.lit, Expr: $4}
  }
  | '(' IF expr expr ')'
  {
    $$ = &IfStatement{Expr: $3, True: $4}
  }

exprs
  : expr
  {
    $$ = []Expression{$1}
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
  | '(' FN args statements ')'
  {
    $$ = &FnExpression{Fns: []Fn{Fn{args: $3, stmts: $4}}}
  }
  | '(' expr exprs ')'
  {
    $$ = &CallExpression{Expr: $2, Args: $3}
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
  | '(' '=' exprs ')'
  { $$ = &BinOpExpression{HS: $3, Operator: int('=')} }

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
  : '(' args statements ')'
  {
    $$ = Fn{args: $2, stmts: $3}
  }

fns
  : fn
  {
    $$ = []Fn{$1}
  }
  | fn fns
  {
    $$ = append([]Fn{$1}, $2...)
  }

args
  : '[' idents ']'
  { $$ = $2 }

idents
  : { $$ = nil }
  | IDENT idents
  { $$ = append([]string{$1.lit}, $2...)}

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