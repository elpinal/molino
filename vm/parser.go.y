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

%}

%union{
  statements []Statement
  statement  Statement
  exprs      []Expression
  expr       Expression
  tok        Token
}

%type<statements> statements
%type<statement>  statement
%type<exprs>      exprs
%type<expr>       expr

%token<tok> IDENT NUMBER

%%

statements
  :
  {
    $$ = nil
    if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
      l.statements = $$
    }
  }
  | statements statement
  {
    $$ = append($1, $2)
    if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
      l.statements = $$
    }
  }

statement
  : expr
  {
    $$ = &ExpressionStatement{$1}
  }

exprs
  : 
  {
    $$ = []Expression{}
  }
  | exprs expr
  {
    $$ = append($1, $2)
  }

expr
  : NUMBER
  {
    $$ = &NumberExpression{$1.lit}
  }
  | IDENT
  {
    $$ = &IdentifierExpression{$1.lit}
  }
  | '[' exprs ']'
  {
    $$ = &VectorExpression{$2}
  }
  | '{' exprs '}'
  {
    $$ = &MapExpression{$2}
  }
  | '(' exprs ')'
  {
    $$ = &CallExpression{$2}
  }

%%
/*
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
*/
