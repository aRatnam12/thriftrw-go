// @generated Code generated by yacc


//line thrift.y:2
package internal

import __yyfmt__ "fmt"

//line thrift.y:2
import "github.com/uber/thriftrw-go/ast"

//line thrift.y:7
type yySymType struct {
	yys int
	// Used to record line numbers when the line number at the start point is
	// required.
	line int

	// Holds the final AST for the file.
	prog *ast.Program

	// Other intermediate variables:

	str string
	i64 int64
	dub float64

	fieldType     ast.Type
	structType    ast.StructureType
	baseTypeID    ast.BaseTypeID
	fieldRequired ast.Requiredness

	field  *ast.Field
	fields []*ast.Field

	header  ast.Header
	headers []ast.Header

	enumItem  *ast.EnumItem
	enumItems []*ast.EnumItem

	definition  ast.Definition
	definitions []ast.Definition

	typeAnnotations []*ast.Annotation

	constantValue    ast.ConstantValue
	constantValues   []ast.ConstantValue
	constantMapItems []ast.ConstantMapItem
}

const IDENTIFIER = 57346
const LITERAL = 57347
const INTCONSTANT = 57348
const DUBCONSTANT = 57349
const NAMESPACE = 57350
const INCLUDE = 57351
const VOID = 57352
const BOOL = 57353
const BYTE = 57354
const I16 = 57355
const I32 = 57356
const I64 = 57357
const DOUBLE = 57358
const STRING = 57359
const BINARY = 57360
const MAP = 57361
const LIST = 57362
const SET = 57363
const ONEWAY = 57364
const TYPEDEF = 57365
const STRUCT = 57366
const UNION = 57367
const EXCEPTION = 57368
const EXTENDS = 57369
const THROWS = 57370
const SERVICE = 57371
const ENUM = 57372
const CONST = 57373
const REQUIRED = 57374
const OPTIONAL = 57375
const TRUE = 57376
const FALSE = 57377

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"LITERAL",
	"INTCONSTANT",
	"DUBCONSTANT",
	"NAMESPACE",
	"INCLUDE",
	"VOID",
	"BOOL",
	"BYTE",
	"I16",
	"I32",
	"I64",
	"DOUBLE",
	"STRING",
	"BINARY",
	"MAP",
	"LIST",
	"SET",
	"ONEWAY",
	"TYPEDEF",
	"STRUCT",
	"UNION",
	"EXCEPTION",
	"EXTENDS",
	"THROWS",
	"SERVICE",
	"ENUM",
	"CONST",
	"REQUIRED",
	"OPTIONAL",
	"TRUE",
	"FALSE",
	"'*'",
	"'='",
	"'{'",
	"'}'",
	"':'",
	"'<'",
	"','",
	"'>'",
	"'['",
	"']'",
	"'('",
	"')'",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	8, 56,
	9, 56,
	-2, 7,
	-1, 3,
	1, 1,
	-2, 56,
}

const yyNprod = 60
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 139

var yyAct = [...]int{

	43, 117, 28, 69, 60, 5, 7, 65, 61, 62,
	44, 44, 65, 61, 62, 98, 10, 23, 11, 65,
	61, 62, 100, 73, 12, 72, 71, 47, 46, 45,
	103, 44, 37, 101, 77, 74, 63, 64, 51, 50,
	68, 63, 64, 97, 22, 68, 67, 93, 63, 64,
	57, 67, 68, 95, 52, 66, 70, 106, 67, 108,
	109, 76, 79, 54, 55, 56, 9, 8, 92, 104,
	20, 115, 89, 85, 86, 87, 21, 83, 90, 80,
	49, 48, 42, 53, 66, 66, 94, 96, 41, 84,
	99, 40, 88, 39, 38, 91, 82, 81, 3, 105,
	6, 58, 14, 17, 18, 19, 66, 112, 110, 15,
	13, 102, 75, 2, 4, 59, 116, 78, 16, 119,
	66, 111, 118, 107, 24, 113, 1, 114, 29, 30,
	31, 32, 33, 34, 35, 36, 25, 26, 27,
}
var yyPact = [...]int{

	-1000, -1000, -1000, -1000, -1000, 58, -24, 79, 65, 40,
	-1000, -1000, -1000, 117, 117, 90, 89, -1000, -1000, -1000,
	-1000, 87, 84, 78, -35, -12, -13, -14, 77, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 76, 1, 0,
	-1000, -1000, 17, -1000, -1000, 117, 117, 117, -1000, -35,
	-1000, -1000, 7, -44, -16, -18, -20, -1000, -4, -5,
	-1000, -1000, -1000, -1000, -1000, -1000, 75, -1000, -1000, -1000,
	73, 117, -35, -35, -35, -24, 68, -35, -24, 62,
	-1000, 2, 14, 6, -28, -1000, -1000, -1000, -1000, -15,
	-1000, -1000, -7, -1000, -24, -1000, -10, 64, -35, -1000,
	51, 27, -1000, 7, -24, -1000, -35, 117, -1000, -1000,
	-24, -1000, -1000, 67, -1000, -36, -1000, 7, -35, -1000,
}
var yyPgo = [...]int{

	0, 2, 126, 17, 124, 123, 118, 117, 115, 114,
	113, 112, 101, 100, 98, 4, 97, 96, 83, 0,
	16,
}
var yyR1 = [...]int{

	0, 2, 10, 10, 9, 9, 9, 14, 14, 13,
	13, 13, 13, 6, 6, 6, 12, 12, 11, 11,
	8, 8, 7, 7, 5, 5, 5, 3, 3, 3,
	3, 3, 4, 4, 4, 4, 4, 4, 4, 4,
	15, 15, 15, 15, 15, 15, 15, 15, 16, 16,
	17, 17, 19, 19, 18, 18, 1, 20, 20, 20,
}
var yyR2 = [...]int{

	0, 2, 0, 2, 3, 4, 4, 0, 3, 6,
	5, 7, 7, 1, 1, 1, 0, 3, 3, 5,
	0, 3, 7, 9, 1, 1, 0, 2, 7, 5,
	5, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 3, 0, 3,
	0, 5, 0, 3, 0, 6, 0, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -2, -10, -14, -9, -1, -13, -1, 9, 8,
	-20, 42, 48, 31, 23, 30, -6, 24, 25, 26,
	5, 36, 4, -3, -4, 19, 20, 21, -1, 11,
	12, 13, 14, 15, 16, 17, 18, -3, 4, 4,
	4, 4, 4, -19, 46, 41, 41, 41, 4, 4,
	38, 38, 37, -18, -3, -3, -3, -19, -12, -8,
	-15, 6, 7, 34, 35, 5, -1, 44, 38, 47,
	-1, 42, 43, 43, 39, -11, -1, 39, -7, -1,
	4, -16, -17, 4, -3, -19, -19, -19, -20, 4,
	-19, -20, 6, 45, -15, 39, -15, 37, 43, -19,
	37, 40, -20, 40, 5, -19, 6, -5, 32, 33,
	-15, -20, -19, -3, -20, 4, -19, 37, -15, -19,
}
var yyDef = [...]int{

	2, -2, -2, -2, 3, 0, 59, 0, 0, 0,
	8, 57, 58, 56, 56, 0, 0, 13, 14, 15,
	4, 0, 0, 0, 52, 0, 0, 0, 0, 32,
	33, 34, 35, 36, 37, 38, 39, 0, 0, 0,
	5, 6, 0, 27, 54, 56, 56, 56, 31, 52,
	16, 20, 56, 56, 0, 0, 0, 10, 56, 56,
	9, 40, 41, 42, 43, 44, 0, 48, 50, 53,
	0, 56, 52, 52, 52, 59, 0, 52, 59, 0,
	45, 56, 56, 0, 0, 29, 30, 11, 17, 52,
	12, 21, 0, 46, 59, 47, 0, 0, 52, 18,
	0, 26, 49, 56, 59, 28, 52, 56, 24, 25,
	59, 55, 19, 0, 51, 52, 22, 56, 52, 23,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 47, 36, 3, 42, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 40, 48,
	41, 37, 43, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 44, 3, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 38, 3, 39,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
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
			yychar, yytoken = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yytoken {
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
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
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
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line thrift.y:85
		{
			yyVAL.prog = &ast.Program{}

			for _, header := range yyDollar[1].headers {
				yyVAL.prog.AddHeader(header)
			}

			for _, def := range yyDollar[2].definitions {
				yyVAL.prog.AddDefinition(def)
			}

			yylex.(*lexer).program = yyVAL.prog
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:106
		{
			yyVAL.headers = nil
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line thrift.y:107
		{
			yyVAL.headers = append(yyDollar[1].headers, yyDollar[2].header)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:112
		{
			yyVAL.header = &ast.Include{
				Path: yyDollar[3].str,
				Line: yyDollar[1].line,
			}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line thrift.y:119
		{
			yyVAL.header = &ast.Namespace{
				Scope: "*",
				Name:  yyDollar[4].str,
				Line:  yyDollar[1].line,
			}
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line thrift.y:127
		{
			yyVAL.header = &ast.Namespace{
				Scope: yyDollar[3].str,
				Name:  yyDollar[4].str,
				Line:  yyDollar[1].line,
			}
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:141
		{
			yyVAL.definitions = nil
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:142
		{
			yyVAL.definitions = append(yyDollar[1].definitions, yyDollar[2].definition)
		}
	case 9:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line thrift.y:148
		{
			yyVAL.definition = &ast.Constant{
				Name:  yyDollar[4].str,
				Type:  yyDollar[3].fieldType,
				Value: yyDollar[6].constantValue,
				Line:  yyDollar[1].line,
			}
		}
	case 10:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line thrift.y:157
		{
			yyVAL.definition = &ast.Typedef{
				Name:        yyDollar[4].str,
				Type:        yyDollar[3].fieldType,
				Annotations: yyDollar[5].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 11:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line thrift.y:166
		{
			yyVAL.definition = &ast.Enum{
				Name:        yyDollar[3].str,
				Items:       yyDollar[5].enumItems,
				Annotations: yyDollar[7].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line thrift.y:175
		{
			yyVAL.definition = &ast.Struct{
				Name:        yyDollar[3].str,
				Type:        yyDollar[2].structType,
				Fields:      yyDollar[5].fields,
				Annotations: yyDollar[7].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:187
		{
			yyVAL.structType = ast.StructType
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:188
		{
			yyVAL.structType = ast.UnionType
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:189
		{
			yyVAL.structType = ast.ExceptionType
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:193
		{
			yyVAL.enumItems = nil
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:194
		{
			yyVAL.enumItems = append(yyDollar[1].enumItems, yyDollar[2].enumItem)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:199
		{
			yyVAL.enumItem = &ast.EnumItem{Name: yyDollar[2].str, Annotations: yyDollar[3].typeAnnotations, Line: yyDollar[1].line}
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line thrift.y:201
		{
			value := int(yyDollar[4].i64)
			yyVAL.enumItem = &ast.EnumItem{
				Name:        yyDollar[2].str,
				Value:       &value,
				Annotations: yyDollar[5].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:213
		{
			yyVAL.fields = nil
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:214
		{
			yyVAL.fields = append(yyDollar[1].fields, yyDollar[2].field)
		}
	case 22:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line thrift.y:220
		{
			yyVAL.field = &ast.Field{
				ID:           int(yyDollar[2].i64),
				Name:         yyDollar[6].str,
				Type:         yyDollar[5].fieldType,
				Requiredness: yyDollar[4].fieldRequired,
				Annotations:  yyDollar[7].typeAnnotations,
				Line:         yyDollar[1].line,
			}
		}
	case 23:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line thrift.y:232
		{
			yyVAL.field = &ast.Field{
				ID:           int(yyDollar[2].i64),
				Name:         yyDollar[6].str,
				Type:         yyDollar[5].fieldType,
				Requiredness: yyDollar[4].fieldRequired,
				Default:      yyDollar[8].constantValue,
				Annotations:  yyDollar[9].typeAnnotations,
				Line:         yyDollar[1].line,
			}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:247
		{
			yyVAL.fieldRequired = ast.Required
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:248
		{
			yyVAL.fieldRequired = ast.Optional
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:249
		{
			yyVAL.fieldRequired = ast.Unspecified
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line thrift.y:258
		{
			yyVAL.fieldType = ast.BaseType{ID: yyDollar[1].baseTypeID, Annotations: yyDollar[2].typeAnnotations}
		}
	case 28:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line thrift.y:262
		{
			yyVAL.fieldType = ast.MapType{KeyType: yyDollar[3].fieldType, ValueType: yyDollar[5].fieldType, Annotations: yyDollar[7].typeAnnotations}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line thrift.y:264
		{
			yyVAL.fieldType = ast.ListType{ValueType: yyDollar[3].fieldType, Annotations: yyDollar[5].typeAnnotations}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line thrift.y:266
		{
			yyVAL.fieldType = ast.SetType{ValueType: yyDollar[3].fieldType, Annotations: yyDollar[5].typeAnnotations}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line thrift.y:268
		{
			yyVAL.fieldType = ast.TypeReference{Name: yyDollar[2].str, Line: yyDollar[1].line}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:272
		{
			yyVAL.baseTypeID = ast.BoolBaseTypeID
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:273
		{
			yyVAL.baseTypeID = ast.ByteBaseTypeID
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:274
		{
			yyVAL.baseTypeID = ast.I16BaseTypeID
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:275
		{
			yyVAL.baseTypeID = ast.I32BaseTypeID
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:276
		{
			yyVAL.baseTypeID = ast.I64BaseTypeID
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:277
		{
			yyVAL.baseTypeID = ast.DoubleBaseTypeID
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:278
		{
			yyVAL.baseTypeID = ast.StringBaseTypeID
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:279
		{
			yyVAL.baseTypeID = ast.BinaryBaseTypeID
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:287
		{
			yyVAL.constantValue = ast.ConstantInteger(yyDollar[1].i64)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:288
		{
			yyVAL.constantValue = ast.ConstantDouble(yyDollar[1].dub)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:289
		{
			yyVAL.constantValue = ast.ConstantBoolean(true)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:290
		{
			yyVAL.constantValue = ast.ConstantBoolean(false)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:291
		{
			yyVAL.constantValue = ast.ConstantString(yyDollar[1].str)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line thrift.y:293
		{
			yyVAL.constantValue = ast.ConstantReference{Name: yyDollar[2].str, Line: yyDollar[1].line}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:295
		{
			yyVAL.constantValue = ast.ConstantList{Items: yyDollar[2].constantValues}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:296
		{
			yyVAL.constantValue = ast.ConstantMap{Items: yyDollar[2].constantMapItems}
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:300
		{
			yyVAL.constantValues = nil
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:302
		{
			yyVAL.constantValues = append(yyDollar[1].constantValues, yyDollar[2].constantValue)
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:306
		{
			yyVAL.constantMapItems = nil
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line thrift.y:308
		{
			yyVAL.constantMapItems = append(yyDollar[1].constantMapItems, ast.ConstantMapItem{Key: yyDollar[2].constantValue, Value: yyDollar[4].constantValue})
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:316
		{
			yyVAL.typeAnnotations = nil
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:317
		{
			yyVAL.typeAnnotations = yyDollar[2].typeAnnotations
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:321
		{
			yyVAL.typeAnnotations = nil
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line thrift.y:323
		{
			yyVAL.typeAnnotations = append(yyDollar[1].typeAnnotations, &ast.Annotation{Name: yyDollar[3].str, Value: yyDollar[5].str, Line: yyDollar[2].line})
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:340
		{
			yyVAL.line = yylex.(*lexer).line
		}
	}
	goto yystack /* stack new state and value */
}
