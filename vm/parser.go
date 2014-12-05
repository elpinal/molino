
//line parser.go.y:2
package vm
import __yyfmt__ "fmt"
//line parser.go.y:2
		
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


//line parser.go.y:21
type yySymType struct{
	yys int
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

const IDENT = 57346
const NUMBER = 57347
const KEYWORD = 57348
const STRING = 57349
const VAR = 57350
const IF = 57351
const TRUE = 57352
const FALSE = 57353
const FN = 57354
const UNARY = 57355

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"KEYWORD",
	"STRING",
	"VAR",
	"IF",
	"TRUE",
	"FALSE",
	"FN",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:193


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

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 40,
	20, 6,
	-2, 11,
}

const yyNprod = 35
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 178

var yyAct = []int{

	29, 3, 55, 3, 1, 19, 34, 15, 69, 27,
	33, 46, 37, 36, 36, 72, 68, 67, 32, 66,
	64, 63, 40, 6, 5, 9, 10, 19, 62, 13,
	14, 49, 50, 51, 61, 3, 60, 59, 26, 53,
	11, 58, 12, 52, 57, 37, 54, 56, 31, 7,
	65, 35, 30, 2, 0, 0, 0, 0, 3, 70,
	0, 0, 71, 6, 5, 9, 10, 16, 17, 13,
	14, 18, 20, 21, 22, 23, 24, 0, 26, 0,
	11, 0, 12, 0, 25, 6, 5, 9, 10, 0,
	0, 13, 14, 18, 20, 21, 22, 23, 24, 0,
	26, 0, 11, 0, 12, 0, 25, 6, 5, 9,
	10, 47, 0, 13, 14, 0, 0, 8, 0, 0,
	0, 0, 26, 28, 11, 0, 12, 48, 0, 0,
	0, 38, 39, 41, 42, 43, 44, 45, 6, 5,
	9, 10, 0, 0, 13, 14, 0, 0, 8, 0,
	0, 0, 0, 4, 0, 11, 0, 12, 6, 5,
	9, 10, 0, 0, 13, 14, 0, 0, 8, 0,
	0, 0, 0, 26, 0, 11, 0, 12,
}
var yyPact = []int{

	134, -1000, 134, -1000, 59, -1000, -1000, -1000, 154, -1000,
	-1000, 154, -1000, -1000, -1000, -1000, 44, 154, -7, 154,
	154, 154, 154, 154, 154, 154, 81, -1000, -11, 154,
	103, 154, 154, 23, 134, 26, 43, -8, 21, 17,
	19, 16, 14, 8, 1, 0, -1000, -1000, -1000, 154,
	-1, -3, -1000, -4, -1000, -14, 43, 134, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -5, -1000,
}
var yyPgo = []int{

	0, 4, 53, 111, 0, 52, 51, 10, 6, 2,
	49,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 3, 3, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 5, 5, 6, 7, 7,
	8, 9, 9, 10, 10,
}
var yyR2 = []int{

	0, 0, 2, 1, 5, 5, 1, 2, 1, 1,
	1, 2, 1, 1, 3, 3, 4, 5, 4, 4,
	4, 4, 4, 4, 4, 0, 3, 4, 1, 2,
	3, 0, 2, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -4, 19, 5, 4, -10, 14, 6,
	7, 21, 23, 10, 11, -1, 8, 9, 12, -4,
	13, 14, 15, 16, 17, 25, 19, -4, -3, -4,
	-5, 4, -4, -7, -8, -6, 21, 19, -3, -3,
	-4, -3, -3, -3, -3, -3, 22, -3, 24, -4,
	-4, -4, 20, -1, -7, -9, 4, -8, 20, 20,
	20, 20, 20, 20, 20, -4, 20, 20, 20, 22,
	-9, -1, 20,
}
var yyDef = []int{

	1, -2, 1, 3, 0, 8, 9, 10, 0, 12,
	13, 0, 25, 33, 34, 2, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 11, 0, 6,
	0, 0, 0, 0, 1, 28, 31, 0, 0, 0,
	-2, 0, 0, 0, 0, 0, 14, 7, 15, 0,
	0, 0, 16, 0, 29, 0, 31, 1, 18, 19,
	20, 21, 22, 23, 24, 26, 4, 5, 17, 30,
	32, 0, 27,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 17, 3, 3,
	19, 20, 15, 13, 3, 14, 3, 16, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 25, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 21, 3, 22, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 23, 3, 24,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 18,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:56
		{
	    yyVAL.statements = nil
	    if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
	      l.statements = yyVAL.statements
	    }
	  }
	case 2:
		//line parser.go.y:63
		{
	    yyVAL.statements = append([]Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
	    if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
	      l.statements = yyVAL.statements
	    }
	  }
	case 3:
		//line parser.go.y:72
		{
	    yyVAL.statement = &ExpressionStatement{Expr: yyS[yypt-0].expr}
	  }
	case 4:
		//line parser.go.y:76
		{
	    yyVAL.statement = &VarDefStatement{VarName: yyS[yypt-2].tok.lit, Expr: yyS[yypt-1].expr}
	  }
	case 5:
		//line parser.go.y:80
		{
	    yyVAL.statement = &IfStatement{Expr: yyS[yypt-2].expr, True: yyS[yypt-1].expr}
	  }
	case 6:
		//line parser.go.y:86
		{
	    yyVAL.exprs = []Expression{yyS[yypt-0].expr}
	  }
	case 7:
		//line parser.go.y:90
		{
	    yyVAL.exprs = append([]Expression{yyS[yypt-1].expr}, yyS[yypt-0].exprs...)
	  }
	case 8:
		//line parser.go.y:95
		{
	    yyVAL.expr = &NumberExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 9:
		//line parser.go.y:99
		{
	    yyVAL.expr = &IdentifierExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 10:
		//line parser.go.y:103
		{
	    yyVAL.expr = &BoolExpression{Bool: yyS[yypt-0].bool}
	  }
	case 11:
		//line parser.go.y:107
		{
	    yyVAL.expr = &UnaryMinusExpression{SubExpr: yyS[yypt-0].expr}
	  }
	case 12:
		//line parser.go.y:111
		{
	    yyVAL.expr = &UnaryKeywordExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 13:
		//line parser.go.y:115
		{
	    yyVAL.expr = &StringExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 14:
		//line parser.go.y:119
		{
	    yyVAL.expr = &VectorExpression{Exprs: yyS[yypt-1].exprs}
	  }
	case 15:
		//line parser.go.y:123
		{
	    yyVAL.expr = &MapExpression{Map: yyS[yypt-1].expr_pairs}
	  }
	case 16:
		//line parser.go.y:127
		{
	    yyVAL.expr = &FnExpression{Fns: yyS[yypt-1].fns}
	  }
	case 17:
		//line parser.go.y:131
		{
	    yyVAL.expr = &FnExpression{Fns: []Fn{Fn{args: yyS[yypt-2].args, stmts: yyS[yypt-1].statements}}}
	  }
	case 18:
		//line parser.go.y:135
		{
	    yyVAL.expr = &CallExpression{Expr: yyS[yypt-2].expr, Args: yyS[yypt-1].exprs}
	  }
	case 19:
		//line parser.go.y:139
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('+')} }
	case 20:
		//line parser.go.y:141
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('-')} }
	case 21:
		//line parser.go.y:143
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('*')} }
	case 22:
		//line parser.go.y:145
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('/')} }
	case 23:
		//line parser.go.y:147
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('%')} }
	case 24:
		//line parser.go.y:149
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('=')} }
	case 25:
		//line parser.go.y:153
		{
	    yyVAL.expr_pairs = map[Expression]Expression{}
	  }
	case 26:
		//line parser.go.y:157
		{
	    yyVAL.expr_pairs = yyS[yypt-2].expr_pairs
	    yyVAL.expr_pairs[yyS[yypt-1].expr] = yyS[yypt-0].expr
	  }
	case 27:
		//line parser.go.y:164
		{
	    yyVAL.fn = Fn{args: yyS[yypt-2].args, stmts: yyS[yypt-1].statements}
	  }
	case 28:
		//line parser.go.y:170
		{
	    yyVAL.fns = []Fn{yyS[yypt-0].fn}
	  }
	case 29:
		//line parser.go.y:174
		{
	    yyVAL.fns = append([]Fn{yyS[yypt-1].fn}, yyS[yypt-0].fns...)
	  }
	case 30:
		//line parser.go.y:180
		{ yyVAL.args = yyS[yypt-1].idents }
	case 31:
		//line parser.go.y:183
		{ yyVAL.idents = nil }
	case 32:
		//line parser.go.y:185
		{ yyVAL.idents = append([]string{yyS[yypt-1].tok.lit}, yyS[yypt-0].idents...)}
	case 33:
		//line parser.go.y:189
		{ yyVAL.bool = true  }
	case 34:
		//line parser.go.y:191
		{ yyVAL.bool = false }
	}
	goto yystack /* stack new state and value */
}
