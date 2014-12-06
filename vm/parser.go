
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

//line parser.go.y:206


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

const yyNprod = 36
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 186

var yyAct = []int{

	47, 34, 35, 74, 46, 37, 36, 36, 73, 70,
	29, 3, 28, 3, 67, 19, 37, 52, 66, 27,
	38, 39, 41, 42, 43, 44, 45, 68, 32, 65,
	63, 69, 40, 62, 61, 54, 53, 19, 60, 56,
	59, 49, 50, 51, 58, 57, 72, 31, 1, 7,
	55, 15, 33, 30, 2, 0, 0, 71, 0, 0,
	64, 6, 5, 9, 10, 16, 17, 13, 14, 18,
	20, 21, 22, 23, 24, 0, 26, 0, 11, 0,
	12, 0, 25, 6, 5, 9, 10, 0, 0, 13,
	14, 18, 20, 21, 22, 23, 24, 0, 26, 0,
	11, 0, 12, 0, 25, 6, 5, 9, 10, 0,
	0, 13, 14, 0, 0, 8, 0, 0, 0, 0,
	26, 0, 11, 0, 12, 48, 6, 5, 9, 10,
	0, 0, 13, 14, 0, 0, 8, 0, 0, 0,
	0, 26, 0, 11, 0, 12, 6, 5, 9, 10,
	0, 0, 13, 14, 0, 0, 8, 0, 0, 0,
	0, 4, 0, 11, 0, 12, 6, 5, 9, 10,
	0, 0, 13, 14, 0, 0, 0, 0, 0, 0,
	0, 26, 0, 11, 0, 12,
}
var yyPact = []int{

	142, -1000, 142, -1000, 57, -1000, -1000, -1000, 122, -1000,
	-1000, 122, -1000, -1000, -1000, -1000, 43, 122, -14, 122,
	122, 122, 122, 122, 122, 122, 79, -1000, -18, 122,
	101, 122, 122, -3, 122, -1000, -1000, -15, 25, 24,
	162, 20, 18, 14, 13, 10, -1000, -1000, -1000, 122,
	9, -2, -1000, -1000, -6, 5, 122, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 42,
	-1000, -12, -19, -1000, -1000,
}
var yyPgo = []int{

	0, 48, 54, 0, 10, 53, 2, 52, 1, 50,
	49,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 3, 3, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 5, 5, 6, 7, 7,
	8, 8, 9, 9, 10, 10,
}
var yyR2 = []int{

	0, 0, 2, 1, 5, 5, 0, 2, 1, 1,
	1, 2, 1, 1, 3, 3, 4, 5, 4, 4,
	4, 4, 4, 4, 4, 0, 3, 4, 1, 2,
	3, 5, 0, 2, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -4, 19, 5, 4, -10, 14, 6,
	7, 21, 23, 10, 11, -1, 8, 9, 12, -4,
	13, 14, 15, 16, 17, 25, 19, -4, -3, -4,
	-5, 4, -4, -7, -8, -6, 21, 19, -3, -3,
	-4, -3, -3, -3, -3, -3, 22, -3, 24, -4,
	-4, -4, 20, -6, -3, -9, -8, 20, 20, 20,
	20, 20, 20, 20, -4, 20, 20, 20, 22, 26,
	4, -3, 4, 20, 22,
}
var yyDef = []int{

	1, -2, 1, 3, 0, 8, 9, 10, 0, 12,
	13, 6, 25, 34, 35, 2, 0, 0, 0, 6,
	6, 6, 6, 6, 6, 6, 0, 11, 0, 6,
	0, 0, 0, 0, 6, 28, 32, 0, 0, 0,
	-2, 0, 0, 0, 0, 0, 14, 7, 15, 0,
	0, 0, 16, 29, 0, 0, 6, 18, 19, 20,
	21, 22, 23, 24, 26, 4, 5, 17, 30, 0,
	33, 0, 0, 27, 31,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 17, 26, 3,
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
		//line parser.go.y:82
		{
	    yyVAL.statement = &VarDefStatement{VarName: yyS[yypt-2].tok.lit, Expr: yyS[yypt-1].expr}
	  }
	case 5:
		//line parser.go.y:86
		{
	    yyVAL.statement = &IfStatement{Expr: yyS[yypt-2].expr, True: yyS[yypt-1].expr}
	  }
	case 6:
		//line parser.go.y:92
		{
	    yyVAL.exprs = []Expression{}
	  }
	case 7:
		//line parser.go.y:96
		{
	    yyVAL.exprs = append([]Expression{yyS[yypt-1].expr}, yyS[yypt-0].exprs...)
	  }
	case 8:
		//line parser.go.y:101
		{
	    yyVAL.expr = &NumberExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 9:
		//line parser.go.y:105
		{
	    yyVAL.expr = &IdentifierExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 10:
		//line parser.go.y:109
		{
	    yyVAL.expr = &BoolExpression{Bool: yyS[yypt-0].bool}
	  }
	case 11:
		//line parser.go.y:113
		{
	    yyVAL.expr = &UnaryMinusExpression{SubExpr: yyS[yypt-0].expr}
	  }
	case 12:
		//line parser.go.y:117
		{
	    yyVAL.expr = &UnaryKeywordExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 13:
		//line parser.go.y:121
		{
	    yyVAL.expr = &StringExpression{Lit: yyS[yypt-0].tok.lit}
	  }
	case 14:
		//line parser.go.y:125
		{
	    yyVAL.expr = &VectorExpression{Exprs: yyS[yypt-1].exprs}
	  }
	case 15:
		//line parser.go.y:129
		{
	    yyVAL.expr = &MapExpression{Map: yyS[yypt-1].expr_pairs}
	  }
	case 16:
		//line parser.go.y:133
		{
	    yyVAL.expr = &FnExpression{Fns: yyS[yypt-1].fns}
	  }
	case 17:
		//line parser.go.y:137
		{
	    yyVAL.expr = &FnExpression{Fns: []Fn{Fn{Args: yyS[yypt-2].args, Exprs: yyS[yypt-1].exprs}}}
	  }
	case 18:
		//line parser.go.y:141
		{
	    yyVAL.expr = &CallExpression{Expr: yyS[yypt-2].expr, Args: yyS[yypt-1].exprs}
	  }
	case 19:
		//line parser.go.y:150
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('+')} }
	case 20:
		//line parser.go.y:152
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('-')} }
	case 21:
		//line parser.go.y:154
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('*')} }
	case 22:
		//line parser.go.y:156
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('/')} }
	case 23:
		//line parser.go.y:158
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('%')} }
	case 24:
		//line parser.go.y:160
		{ yyVAL.expr = &BinOpExpression{HS: yyS[yypt-1].exprs, Operator: int('=')} }
	case 25:
		//line parser.go.y:164
		{
	    yyVAL.expr_pairs = map[Expression]Expression{}
	  }
	case 26:
		//line parser.go.y:168
		{
	    yyVAL.expr_pairs = yyS[yypt-2].expr_pairs
	    yyVAL.expr_pairs[yyS[yypt-1].expr] = yyS[yypt-0].expr
	  }
	case 27:
		//line parser.go.y:175
		{
	    yyVAL.fn = Fn{Args: yyS[yypt-2].args, Exprs: yyS[yypt-1].exprs}
	  }
	case 28:
		//line parser.go.y:181
		{
	    yyVAL.fns = []Fn{yyS[yypt-0].fn}
	  }
	case 29:
		//line parser.go.y:185
		{
	    yyVAL.fns = append(yyS[yypt-1].fns, yyS[yypt-0].fn)
	  }
	case 30:
		//line parser.go.y:191
		{ yyVAL.args = Args{Args: yyS[yypt-1].idents} }
	case 31:
		//line parser.go.y:193
		{ yyVAL.args = Args{Args: yyS[yypt-3].idents, Vararg: true, More: yyS[yypt-1].tok.lit} }
	case 32:
		//line parser.go.y:196
		{ yyVAL.idents = []string{} }
	case 33:
		//line parser.go.y:198
		{ yyVAL.idents = append(yyS[yypt-1].idents, yyS[yypt-0].tok.lit) }
	case 34:
		//line parser.go.y:202
		{ yyVAL.bool = true  }
	case 35:
		//line parser.go.y:204
		{ yyVAL.bool = false }
	}
	goto yystack /* stack new state and value */
}
