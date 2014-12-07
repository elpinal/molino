
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
  Args  Args
  Exprs []Expression
}

type Args struct {
  Args   []string
  Vararg bool
  More   string
}


//line parser.go.y:27
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
  args       Args
  idents     []string
  bool       bool
}

const IDENT = 57346
const NUMBER = 57347
const KEYWORD = 57348
const STRING = 57349
const DEF = 57350
const IF = 57351
const TRUE = 57352
const FALSE = 57353
const NIL = 57354
const FN = 57355
const UNARY = 57356

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"KEYWORD",
	"STRING",
	"DEF",
	"IF",
	"TRUE",
	"FALSE",
	"NIL",
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

//line parser.go.y:211


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
	-1, 45,
	25, 4,
	-2, 10,
}

const yyNprod = 38
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 175

var yyAct = []int{

	19, 3, 75, 3, 36, 74, 68, 39, 51, 17,
	32, 70, 37, 65, 22, 64, 63, 62, 61, 60,
	59, 34, 18, 66, 56, 42, 38, 38, 45, 67,
	39, 76, 31, 40, 73, 50, 43, 44, 46, 47,
	48, 49, 57, 58, 55, 41, 1, 53, 52, 16,
	6, 54, 35, 20, 2, 0, 0, 0, 0, 72,
	0, 0, 0, 0, 0, 0, 69, 5, 4, 9,
	10, 23, 24, 14, 15, 7, 21, 26, 27, 28,
	29, 30, 0, 11, 0, 12, 0, 13, 0, 25,
	5, 4, 9, 10, 0, 0, 14, 15, 7, 0,
	0, 8, 0, 0, 0, 0, 11, 0, 12, 0,
	13, 71, 5, 4, 9, 10, 0, 0, 14, 15,
	7, 0, 0, 8, 0, 0, 0, 0, 11, 0,
	12, 33, 13, 5, 4, 9, 10, 0, 0, 14,
	15, 7, 0, 0, 8, 0, 0, 0, 0, 11,
	0, 12, 0, 13, 5, 4, 9, 10, 0, 0,
	14, 15, 7, 0, 0, 0, 0, 0, 0, 0,
	11, 0, 12, 0, 13,
}
var yyPact = []int{

	129, -1000, 129, -1000, -1000, -1000, -1000, -1000, 129, -1000,
	-1000, 129, -1000, 63, -1000, -1000, -1000, -1000, 11, 129,
	108, 6, 129, 41, 129, 129, 129, 129, 129, 129,
	129, -1000, -1000, -1000, 129, -17, 129, -1000, -1000, 7,
	-1, 129, 129, -5, -6, 150, -7, -8, -9, -10,
	-1000, -1000, -1000, -12, 2, 129, -1000, -14, 86, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 30, -1000, -20,
	-1000, -1000, -23, 10, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 46, 54, 10, 0, 53, 12, 52, 4, 51,
	50,
}
var yyR1 = []int{

	0, 1, 1, 2, 3, 3, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 5, 5, 6,
	7, 7, 8, 8, 9, 9, 10, 10,
}
var yyR2 = []int{

	0, 0, 2, 1, 0, 2, 1, 1, 1, 1,
	2, 1, 1, 3, 3, 4, 5, 4, 5, 5,
	6, 4, 4, 4, 4, 4, 4, 0, 3, 4,
	1, 2, 3, 5, 0, 2, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -4, 5, 4, -10, 12, 15, 6,
	7, 20, 22, 24, 10, 11, -1, -4, -3, -4,
	-5, 13, -4, 8, 9, 26, 14, 15, 16, 17,
	18, 21, -3, 23, -4, -7, -8, -6, 20, 24,
	-3, 4, -4, -3, -3, -4, -3, -3, -3, -3,
	-4, 25, -6, -3, -9, -8, 25, -4, -4, 25,
	25, 25, 25, 25, 25, 25, 21, 27, 4, -3,
	25, 25, -4, 4, 25, 25, 21,
}
var yyDef = []int{

	1, -2, 1, 3, 6, 7, 8, 9, 0, 11,
	12, 4, 27, 0, 36, 37, 2, 10, 0, 4,
	0, 0, 4, 0, 0, 4, 4, 4, 4, 4,
	4, 13, 5, 14, 0, 0, 4, 30, 34, 0,
	0, 0, 0, 0, 0, -2, 0, 0, 0, 0,
	28, 15, 31, 0, 0, 4, 17, 0, 0, 21,
	22, 23, 24, 25, 26, 16, 32, 0, 35, 0,
	18, 19, 0, 0, 29, 20, 33,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 18, 27, 3,
	24, 25, 16, 14, 3, 15, 3, 17, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 26, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 20, 3, 21, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 22, 3, 23,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 19,
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
		//line parser.go.y:62
		{
	    yyVAL.statements = nil
	    if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
	      l.statements = yyVAL.statements
	    }
	  }
	case 2:
		//line parser.go.y:69
		{
	    yyVAL.statements = append([]Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
	    if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
	      l.statements = yyVAL.statements
	    }
	  }
	case 3:
		//line parser.go.y:78
		{
	    yyVAL.statement = &ExpressionStatement{Expr: yyS[yypt-0].expr}
	  }
	case 4:
		//line parser.go.y:84
		{
	    yyVAL.exprs = []Expression{}
	  }
	case 5:
		//line parser.go.y:88
		{
	    yyVAL.exprs = append([]Expression{yyS[yypt-1].expr}, yyS[yypt-0].exprs...)
	  }
	case 6:
		//line parser.go.y:93
		{
	    yyVAL.expr = &NumberExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 7:
		//line parser.go.y:97
		{
	    yyVAL.expr = &IdentifierExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 8:
		//line parser.go.y:101
		{
	    yyVAL.expr = &BoolExpression{Bool: yyS[yypt-0].bool}
	  }
	case 9:
		//line parser.go.y:105
		{
	    yyVAL.expr = &NilExpression{}
	  }
	case 10:
		//line parser.go.y:109
		{
	    yyVAL.expr = &UnaryMinusExpression{SubExpr: yyS[yypt-0].expr}
	  }
	case 11:
		//line parser.go.y:113
		{
	    yyVAL.expr = &UnaryKeywordExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 12:
		//line parser.go.y:117
		{
	    yyVAL.expr = &StringExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 13:
		//line parser.go.y:121
		{
	    yyVAL.expr = &VectorExpression{Exprs: yyS[yypt-1].exprs}
	  }
	case 14:
		//line parser.go.y:125
		{
	    yyVAL.expr = &MapExpression{Map: yyS[yypt-1].expr_pairs}
	  }
	case 15:
		//line parser.go.y:129
		{
	    yyVAL.expr = &FnExpression{Fns: yyS[yypt-1].fns}
	  }
	case 16:
		//line parser.go.y:133
		{
	    yyVAL.expr = &FnExpression{Fns: []Fn{Fn{Args: yyS[yypt-2].args, Exprs: yyS[yypt-1].exprs}}}
	  }
	case 17:
		//line parser.go.y:137
		{
	    yyVAL.expr = &CallExpression{Expr: yyS[yypt-2].expr, Args: yyS[yypt-1].exprs}
	  }
	case 18:
		//line parser.go.y:141
		{
	    yyVAL.expr = &DefExpression{VarName: yyS[yypt-2].tok.lit, Expr: yyS[yypt-1].expr}
	  }
	case 19:
		//line parser.go.y:145
		{
	    yyVAL.expr = &IfExpression{Expr: yyS[yypt-2].expr, True: yyS[yypt-1].expr}
	  }
	case 20:
		//line parser.go.y:149
		{
	    yyVAL.expr = &IfExpression{Expr: yyS[yypt-3].expr, True: yyS[yypt-2].expr, False: yyS[yypt-1].expr}
	  }
	case 21:
		//line parser.go.y:153
		{
	    yyVAL.expr = &EqualExpression{HS: yyS[yypt-1].exprs}
	  }
	case 22:
		//line parser.go.y:157
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('+')} }
	case 23:
		//line parser.go.y:159
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('-')} }
	case 24:
		//line parser.go.y:161
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('*')} }
	case 25:
		//line parser.go.y:163
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('/')} }
	case 26:
		//line parser.go.y:165
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('%')} }
	case 27:
		//line parser.go.y:169
		{
	    yyVAL.expr_pairs = map[Expression]Expression{}
	  }
	case 28:
		//line parser.go.y:173
		{
	    yyVAL.expr_pairs = yyS[yypt-2].expr_pairs
	    yyVAL.expr_pairs[yyS[yypt-1].expr] = yyS[yypt-0].expr
	  }
	case 29:
		//line parser.go.y:180
		{
	    yyVAL.fn = Fn{Args: yyS[yypt-2].args, Exprs: yyS[yypt-1].exprs}
	  }
	case 30:
		//line parser.go.y:186
		{
	    yyVAL.fns = []Fn{yyS[yypt-0].fn}
	  }
	case 31:
		//line parser.go.y:190
		{
	    yyVAL.fns = append(yyS[yypt-1].fns, yyS[yypt-0].fn)
	  }
	case 32:
		//line parser.go.y:196
		{ yyVAL.args = Args{Args: yyS[yypt-1].idents} }
	case 33:
		//line parser.go.y:198
		{ yyVAL.args = Args{Args: yyS[yypt-3].idents, Vararg: true, More: yyS[yypt-1].tok.lit} }
	case 34:
		//line parser.go.y:201
		{ yyVAL.idents = []string{} }
	case 35:
		//line parser.go.y:203
		{ yyVAL.idents = append(yyS[yypt-1].idents, yyS[yypt-0].tok.lit) }
	case 36:
		//line parser.go.y:207
		{ yyVAL.bool = true  }
	case 37:
		//line parser.go.y:209
		{ yyVAL.bool = false }
	}
	goto yystack /* stack new state and value */
}
