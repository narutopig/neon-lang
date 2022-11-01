// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // neon

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type neonParser struct {
	*antlr.BaseParser
}

var neonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func neonParserInit() {
	staticData := &neonParserStaticData
	staticData.literalNames = []string{
		"", "'number'", "'string'", "'bool'", "'void'", "'def'", "'return'",
		"'if'", "'elif'", "'else'", "'while'", "'=='", "'!='", "';'", "'<='",
		"'>='", "", "", "", "", "", "", "','", "'('", "')'", "'{'", "'}'", "'='",
		"'<'", "'>'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "NUMTYPE", "STRTYPE", "BOOLTYPE", "VOID", "DEF", "RETURN", "IF",
		"ELIF", "ELSE", "WHILE", "EQUALITY", "NOTEQUAL", "SEMI", "LESSEQ", "GREATEREQ",
		"STRING", "ID", "INT", "WS", "ADD_SUB", "MDM", "COMMA", "Lb", "Rb",
		"Lc", "Rc", "EQUAL", "LESS", "GREATER", "BANG", "BOOL",
	}
	staticData.ruleNames = []string{
		"program", "stat", "if", "elif", "else", "while", "func", "funccall",
		"type", "arithop", "compop", "op", "funcarg", "decl", "assign", "return",
		"expr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 204, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 3, 0, 37, 8, 0, 5, 0, 39, 8, 0, 10, 0, 12, 0,
		42, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 50, 8, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 56, 8, 1, 3, 1, 58, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 66, 8, 2, 10, 2, 12, 2, 69, 9, 2, 1, 2, 1, 2, 5, 2, 73,
		8, 2, 10, 2, 12, 2, 76, 9, 2, 1, 2, 3, 2, 79, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 87, 8, 3, 10, 3, 12, 3, 90, 9, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 5, 4, 97, 8, 4, 10, 4, 12, 4, 100, 9, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 110, 8, 5, 10, 5, 12, 5, 113, 9, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 120, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 5, 6, 127, 8, 6, 10, 6, 12, 6, 130, 9, 6, 3, 6, 132, 8, 6, 1, 6,
		3, 6, 135, 8, 6, 1, 6, 1, 6, 5, 6, 139, 8, 6, 10, 6, 12, 6, 142, 9, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 151, 8, 7, 10, 7, 12, 7,
		154, 9, 7, 3, 7, 156, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 3, 11, 168, 8, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 193, 8,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 199, 8, 16, 10, 16, 12, 16, 202,
		9, 16, 1, 16, 0, 1, 32, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 0, 3, 1, 0, 1, 3, 1, 0, 20, 21, 3, 0, 11, 11, 14, 15,
		28, 29, 213, 0, 40, 1, 0, 0, 0, 2, 57, 1, 0, 0, 0, 4, 59, 1, 0, 0, 0, 6,
		80, 1, 0, 0, 0, 8, 93, 1, 0, 0, 0, 10, 103, 1, 0, 0, 0, 12, 116, 1, 0,
		0, 0, 14, 145, 1, 0, 0, 0, 16, 159, 1, 0, 0, 0, 18, 161, 1, 0, 0, 0, 20,
		163, 1, 0, 0, 0, 22, 167, 1, 0, 0, 0, 24, 169, 1, 0, 0, 0, 26, 172, 1,
		0, 0, 0, 28, 177, 1, 0, 0, 0, 30, 181, 1, 0, 0, 0, 32, 192, 1, 0, 0, 0,
		34, 37, 3, 2, 1, 0, 35, 37, 3, 12, 6, 0, 36, 34, 1, 0, 0, 0, 36, 35, 1,
		0, 0, 0, 37, 39, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40,
		38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 43, 1, 0, 0, 0, 42, 40, 1, 0, 0,
		0, 43, 44, 5, 0, 0, 1, 44, 1, 1, 0, 0, 0, 45, 50, 3, 26, 13, 0, 46, 50,
		3, 28, 14, 0, 47, 50, 3, 14, 7, 0, 48, 50, 3, 30, 15, 0, 49, 45, 1, 0,
		0, 0, 49, 46, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 51,
		1, 0, 0, 0, 51, 52, 5, 13, 0, 0, 52, 58, 1, 0, 0, 0, 53, 56, 3, 4, 2, 0,
		54, 56, 3, 10, 5, 0, 55, 53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 58, 1,
		0, 0, 0, 57, 49, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 3, 1, 0, 0, 0, 59,
		60, 5, 7, 0, 0, 60, 61, 5, 23, 0, 0, 61, 62, 3, 32, 16, 0, 62, 63, 5, 24,
		0, 0, 63, 67, 5, 25, 0, 0, 64, 66, 3, 2, 1, 0, 65, 64, 1, 0, 0, 0, 66,
		69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 1, 0, 0,
		0, 69, 67, 1, 0, 0, 0, 70, 74, 5, 26, 0, 0, 71, 73, 3, 6, 3, 0, 72, 71,
		1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0,
		75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 79, 3, 8, 4, 0, 78, 77, 1,
		0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 5, 1, 0, 0, 0, 80, 81, 5, 8, 0, 0, 81,
		82, 5, 23, 0, 0, 82, 83, 3, 32, 16, 0, 83, 84, 5, 24, 0, 0, 84, 88, 5,
		25, 0, 0, 85, 87, 3, 2, 1, 0, 86, 85, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88,
		86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90, 88, 1, 0, 0,
		0, 91, 92, 5, 26, 0, 0, 92, 7, 1, 0, 0, 0, 93, 94, 5, 9, 0, 0, 94, 98,
		5, 25, 0, 0, 95, 97, 3, 2, 1, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0,
		0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 98,
		1, 0, 0, 0, 101, 102, 5, 26, 0, 0, 102, 9, 1, 0, 0, 0, 103, 104, 5, 10,
		0, 0, 104, 105, 5, 23, 0, 0, 105, 106, 3, 32, 16, 0, 106, 107, 5, 24, 0,
		0, 107, 111, 5, 25, 0, 0, 108, 110, 3, 2, 1, 0, 109, 108, 1, 0, 0, 0, 110,
		113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114,
		1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 115, 5, 26, 0, 0, 115, 11, 1, 0,
		0, 0, 116, 119, 5, 5, 0, 0, 117, 120, 3, 16, 8, 0, 118, 120, 5, 4, 0, 0,
		119, 117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121,
		134, 5, 17, 0, 0, 122, 131, 5, 23, 0, 0, 123, 128, 3, 24, 12, 0, 124, 125,
		5, 22, 0, 0, 125, 127, 3, 24, 12, 0, 126, 124, 1, 0, 0, 0, 127, 130, 1,
		0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 132, 1, 0, 0,
		0, 130, 128, 1, 0, 0, 0, 131, 123, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132,
		133, 1, 0, 0, 0, 133, 135, 5, 24, 0, 0, 134, 122, 1, 0, 0, 0, 134, 135,
		1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 140, 5, 25, 0, 0, 137, 139, 3, 2,
		1, 0, 138, 137, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0,
		140, 141, 1, 0, 0, 0, 141, 143, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143,
		144, 5, 26, 0, 0, 144, 13, 1, 0, 0, 0, 145, 146, 5, 17, 0, 0, 146, 155,
		5, 23, 0, 0, 147, 152, 3, 32, 16, 0, 148, 149, 5, 22, 0, 0, 149, 151, 3,
		32, 16, 0, 150, 148, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0,
		0, 0, 152, 153, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0,
		155, 147, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157,
		158, 5, 24, 0, 0, 158, 15, 1, 0, 0, 0, 159, 160, 7, 0, 0, 0, 160, 17, 1,
		0, 0, 0, 161, 162, 7, 1, 0, 0, 162, 19, 1, 0, 0, 0, 163, 164, 7, 2, 0,
		0, 164, 21, 1, 0, 0, 0, 165, 168, 3, 18, 9, 0, 166, 168, 3, 20, 10, 0,
		167, 165, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 23, 1, 0, 0, 0, 169, 170,
		3, 16, 8, 0, 170, 171, 5, 17, 0, 0, 171, 25, 1, 0, 0, 0, 172, 173, 3, 16,
		8, 0, 173, 174, 5, 17, 0, 0, 174, 175, 5, 27, 0, 0, 175, 176, 3, 32, 16,
		0, 176, 27, 1, 0, 0, 0, 177, 178, 5, 17, 0, 0, 178, 179, 5, 27, 0, 0, 179,
		180, 3, 32, 16, 0, 180, 29, 1, 0, 0, 0, 181, 182, 5, 6, 0, 0, 182, 183,
		3, 32, 16, 0, 183, 31, 1, 0, 0, 0, 184, 185, 6, 16, -1, 0, 185, 186, 5,
		30, 0, 0, 186, 193, 3, 32, 16, 6, 187, 193, 3, 14, 7, 0, 188, 193, 5, 18,
		0, 0, 189, 193, 5, 31, 0, 0, 190, 193, 5, 16, 0, 0, 191, 193, 5, 17, 0,
		0, 192, 184, 1, 0, 0, 0, 192, 187, 1, 0, 0, 0, 192, 188, 1, 0, 0, 0, 192,
		189, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 200,
		1, 0, 0, 0, 194, 195, 10, 7, 0, 0, 195, 196, 3, 22, 11, 0, 196, 197, 3,
		32, 16, 8, 197, 199, 1, 0, 0, 0, 198, 194, 1, 0, 0, 0, 199, 202, 1, 0,
		0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 33, 1, 0, 0, 0,
		202, 200, 1, 0, 0, 0, 21, 36, 40, 49, 55, 57, 67, 74, 78, 88, 98, 111,
		119, 128, 131, 134, 140, 152, 155, 167, 192, 200,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// neonParserInit initializes any static state used to implement neonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewneonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NeonParserInit() {
	staticData := &neonParserStaticData
	staticData.once.Do(neonParserInit)
}

// NewneonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewneonParser(input antlr.TokenStream) *neonParser {
	NeonParserInit()
	this := new(neonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &neonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// neonParser tokens.
const (
	neonParserEOF       = antlr.TokenEOF
	neonParserNUMTYPE   = 1
	neonParserSTRTYPE   = 2
	neonParserBOOLTYPE  = 3
	neonParserVOID      = 4
	neonParserDEF       = 5
	neonParserRETURN    = 6
	neonParserIF        = 7
	neonParserELIF      = 8
	neonParserELSE      = 9
	neonParserWHILE     = 10
	neonParserEQUALITY  = 11
	neonParserNOTEQUAL  = 12
	neonParserSEMI      = 13
	neonParserLESSEQ    = 14
	neonParserGREATEREQ = 15
	neonParserSTRING    = 16
	neonParserID        = 17
	neonParserINT       = 18
	neonParserWS        = 19
	neonParserADD_SUB   = 20
	neonParserMDM       = 21
	neonParserCOMMA     = 22
	neonParserLb        = 23
	neonParserRb        = 24
	neonParserLc        = 25
	neonParserRc        = 26
	neonParserEQUAL     = 27
	neonParserLESS      = 28
	neonParserGREATER   = 29
	neonParserBANG      = 30
	neonParserBOOL      = 31
)

// neonParser rules.
const (
	neonParserRULE_program  = 0
	neonParserRULE_stat     = 1
	neonParserRULE_if       = 2
	neonParserRULE_elif     = 3
	neonParserRULE_else     = 4
	neonParserRULE_while    = 5
	neonParserRULE_func     = 6
	neonParserRULE_funccall = 7
	neonParserRULE_type     = 8
	neonParserRULE_arithop  = 9
	neonParserRULE_compop   = 10
	neonParserRULE_op       = 11
	neonParserRULE_funcarg  = 12
	neonParserRULE_decl     = 13
	neonParserRULE_assign   = 14
	neonParserRULE_return   = 15
	neonParserRULE_expr     = 16
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(neonParserEOF, 0)
}

func (s *ProgramContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ProgramContext) AllFunc_() []IFuncContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncContext); ok {
			len++
		}
	}

	tst := make([]IFuncContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncContext); ok {
			tst[i] = t.(IFuncContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Func_(i int) IFuncContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *neonParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, neonParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132334) != 0 {
		p.SetState(36)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case neonParserNUMTYPE, neonParserSTRTYPE, neonParserBOOLTYPE, neonParserRETURN, neonParserIF, neonParserWHILE, neonParserID:
			{
				p.SetState(34)
				p.Stat()
			}

		case neonParserDEF:
			{
				p.SetState(35)
				p.Func_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(43)
		p.Match(neonParserEOF)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) SEMI() antlr.TerminalNode {
	return s.GetToken(neonParserSEMI, 0)
}

func (s *StatContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *StatContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StatContext) Funccall() IFunccallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunccallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunccallContext)
}

func (s *StatContext) Return_() IReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *StatContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *StatContext) While() IWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *neonParser) Stat() (localctx IStatContext) {
	this := p
	_ = this

	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, neonParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case neonParserNUMTYPE, neonParserSTRTYPE, neonParserBOOLTYPE, neonParserRETURN, neonParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(45)
				p.Decl()
			}

		case 2:
			{
				p.SetState(46)
				p.Assign()
			}

		case 3:
			{
				p.SetState(47)
				p.Funccall()
			}

		case 4:
			{
				p.SetState(48)
				p.Return_()
			}

		}
		{
			p.SetState(51)
			p.Match(neonParserSEMI)
		}

	case neonParserIF, neonParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(55)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case neonParserIF:
			{
				p.SetState(53)
				p.If_()
			}

		case neonParserWHILE:
			{
				p.SetState(54)
				p.While()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfContext() *IfContext {
	var p = new(IfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_if
	return p
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(neonParserIF, 0)
}

func (s *IfContext) Lb() antlr.TerminalNode {
	return s.GetToken(neonParserLb, 0)
}

func (s *IfContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfContext) Rb() antlr.TerminalNode {
	return s.GetToken(neonParserRb, 0)
}

func (s *IfContext) Lc() antlr.TerminalNode {
	return s.GetToken(neonParserLc, 0)
}

func (s *IfContext) Rc() antlr.TerminalNode {
	return s.GetToken(neonParserRc, 0)
}

func (s *IfContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *IfContext) AllElif() []IElifContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElifContext); ok {
			len++
		}
	}

	tst := make([]IElifContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElifContext); ok {
			tst[i] = t.(IElifContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) Elif(i int) IElifContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElifContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElifContext)
}

func (s *IfContext) Else_() IElseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseContext)
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitIf(s)
	}
}

func (p *neonParser) If_() (localctx IIfContext) {
	this := p
	_ = this

	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, neonParserRULE_if)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(neonParserIF)
	}
	{
		p.SetState(60)
		p.Match(neonParserLb)
	}
	{
		p.SetState(61)
		p.expr(0)
	}
	{
		p.SetState(62)
		p.Match(neonParserRb)
	}
	{
		p.SetState(63)
		p.Match(neonParserLc)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132302) != 0 {
		{
			p.SetState(64)
			p.Stat()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(neonParserRc)
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == neonParserELIF {
		{
			p.SetState(71)
			p.Elif()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == neonParserELSE {
		{
			p.SetState(77)
			p.Else_()
		}

	}

	return localctx
}

// IElifContext is an interface to support dynamic dispatch.
type IElifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElifContext differentiates from other interfaces.
	IsElifContext()
}

type ElifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElifContext() *ElifContext {
	var p = new(ElifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_elif
	return p
}

func (*ElifContext) IsElifContext() {}

func NewElifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifContext {
	var p = new(ElifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_elif

	return p
}

func (s *ElifContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifContext) ELIF() antlr.TerminalNode {
	return s.GetToken(neonParserELIF, 0)
}

func (s *ElifContext) Lb() antlr.TerminalNode {
	return s.GetToken(neonParserLb, 0)
}

func (s *ElifContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElifContext) Rb() antlr.TerminalNode {
	return s.GetToken(neonParserRb, 0)
}

func (s *ElifContext) Lc() antlr.TerminalNode {
	return s.GetToken(neonParserLc, 0)
}

func (s *ElifContext) Rc() antlr.TerminalNode {
	return s.GetToken(neonParserRc, 0)
}

func (s *ElifContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *ElifContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ElifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterElif(s)
	}
}

func (s *ElifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitElif(s)
	}
}

func (p *neonParser) Elif() (localctx IElifContext) {
	this := p
	_ = this

	localctx = NewElifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, neonParserRULE_elif)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(neonParserELIF)
	}
	{
		p.SetState(81)
		p.Match(neonParserLb)
	}
	{
		p.SetState(82)
		p.expr(0)
	}
	{
		p.SetState(83)
		p.Match(neonParserRb)
	}
	{
		p.SetState(84)
		p.Match(neonParserLc)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132302) != 0 {
		{
			p.SetState(85)
			p.Stat()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.Match(neonParserRc)
	}

	return localctx
}

// IElseContext is an interface to support dynamic dispatch.
type IElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseContext differentiates from other interfaces.
	IsElseContext()
}

type ElseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseContext() *ElseContext {
	var p = new(ElseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_else
	return p
}

func (*ElseContext) IsElseContext() {}

func NewElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseContext {
	var p = new(ElseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_else

	return p
}

func (s *ElseContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(neonParserELSE, 0)
}

func (s *ElseContext) Lc() antlr.TerminalNode {
	return s.GetToken(neonParserLc, 0)
}

func (s *ElseContext) Rc() antlr.TerminalNode {
	return s.GetToken(neonParserRc, 0)
}

func (s *ElseContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *ElseContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterElse(s)
	}
}

func (s *ElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitElse(s)
	}
}

func (p *neonParser) Else_() (localctx IElseContext) {
	this := p
	_ = this

	localctx = NewElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, neonParserRULE_else)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(neonParserELSE)
	}
	{
		p.SetState(94)
		p.Match(neonParserLc)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132302) != 0 {
		{
			p.SetState(95)
			p.Stat()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
		p.Match(neonParserRc)
	}

	return localctx
}

// IWhileContext is an interface to support dynamic dispatch.
type IWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileContext differentiates from other interfaces.
	IsWhileContext()
}

type WhileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileContext() *WhileContext {
	var p = new(WhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_while
	return p
}

func (*WhileContext) IsWhileContext() {}

func NewWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileContext {
	var p = new(WhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_while

	return p
}

func (s *WhileContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(neonParserWHILE, 0)
}

func (s *WhileContext) Lb() antlr.TerminalNode {
	return s.GetToken(neonParserLb, 0)
}

func (s *WhileContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileContext) Rb() antlr.TerminalNode {
	return s.GetToken(neonParserRb, 0)
}

func (s *WhileContext) Lc() antlr.TerminalNode {
	return s.GetToken(neonParserLc, 0)
}

func (s *WhileContext) Rc() antlr.TerminalNode {
	return s.GetToken(neonParserRc, 0)
}

func (s *WhileContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *WhileContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (p *neonParser) While() (localctx IWhileContext) {
	this := p
	_ = this

	localctx = NewWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, neonParserRULE_while)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(neonParserWHILE)
	}
	{
		p.SetState(104)
		p.Match(neonParserLb)
	}
	{
		p.SetState(105)
		p.expr(0)
	}
	{
		p.SetState(106)
		p.Match(neonParserRb)
	}
	{
		p.SetState(107)
		p.Match(neonParserLc)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132302) != 0 {
		{
			p.SetState(108)
			p.Stat()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(114)
		p.Match(neonParserRc)
	}

	return localctx
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_func
	return p
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) DEF() antlr.TerminalNode {
	return s.GetToken(neonParserDEF, 0)
}

func (s *FuncContext) ID() antlr.TerminalNode {
	return s.GetToken(neonParserID, 0)
}

func (s *FuncContext) Lc() antlr.TerminalNode {
	return s.GetToken(neonParserLc, 0)
}

func (s *FuncContext) Rc() antlr.TerminalNode {
	return s.GetToken(neonParserRc, 0)
}

func (s *FuncContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncContext) VOID() antlr.TerminalNode {
	return s.GetToken(neonParserVOID, 0)
}

func (s *FuncContext) Lb() antlr.TerminalNode {
	return s.GetToken(neonParserLb, 0)
}

func (s *FuncContext) Rb() antlr.TerminalNode {
	return s.GetToken(neonParserRb, 0)
}

func (s *FuncContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *FuncContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *FuncContext) AllFuncarg() []IFuncargContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncargContext); ok {
			len++
		}
	}

	tst := make([]IFuncargContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncargContext); ok {
			tst[i] = t.(IFuncargContext)
			i++
		}
	}

	return tst
}

func (s *FuncContext) Funcarg(i int) IFuncargContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncargContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncargContext)
}

func (s *FuncContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(neonParserCOMMA)
}

func (s *FuncContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(neonParserCOMMA, i)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (p *neonParser) Func_() (localctx IFuncContext) {
	this := p
	_ = this

	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, neonParserRULE_func)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(neonParserDEF)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case neonParserNUMTYPE, neonParserSTRTYPE, neonParserBOOLTYPE:
		{
			p.SetState(117)
			p.Type_()
		}

	case neonParserVOID:
		{
			p.SetState(118)
			p.Match(neonParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(121)
		p.Match(neonParserID)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == neonParserLb {
		{
			p.SetState(122)
			p.Match(neonParserLb)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0 {
			{
				p.SetState(123)
				p.Funcarg()
			}
			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == neonParserCOMMA {
				{
					p.SetState(124)
					p.Match(neonParserCOMMA)
				}
				{
					p.SetState(125)
					p.Funcarg()
				}

				p.SetState(130)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(133)
			p.Match(neonParserRb)
		}

	}
	{
		p.SetState(136)
		p.Match(neonParserLc)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132302) != 0 {
		{
			p.SetState(137)
			p.Stat()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(neonParserRc)
	}

	return localctx
}

// IFunccallContext is an interface to support dynamic dispatch.
type IFunccallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunccallContext differentiates from other interfaces.
	IsFunccallContext()
}

type FunccallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunccallContext() *FunccallContext {
	var p = new(FunccallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_funccall
	return p
}

func (*FunccallContext) IsFunccallContext() {}

func NewFunccallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunccallContext {
	var p = new(FunccallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_funccall

	return p
}

func (s *FunccallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunccallContext) ID() antlr.TerminalNode {
	return s.GetToken(neonParserID, 0)
}

func (s *FunccallContext) Lb() antlr.TerminalNode {
	return s.GetToken(neonParserLb, 0)
}

func (s *FunccallContext) Rb() antlr.TerminalNode {
	return s.GetToken(neonParserRb, 0)
}

func (s *FunccallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FunccallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunccallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(neonParserCOMMA)
}

func (s *FunccallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(neonParserCOMMA, i)
}

func (s *FunccallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunccallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunccallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterFunccall(s)
	}
}

func (s *FunccallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitFunccall(s)
	}
}

func (p *neonParser) Funccall() (localctx IFunccallContext) {
	this := p
	_ = this

	localctx = NewFunccallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, neonParserRULE_funccall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(neonParserID)
	}
	{
		p.SetState(146)
		p.Match(neonParserLb)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3221684224) != 0 {
		{
			p.SetState(147)
			p.expr(0)
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == neonParserCOMMA {
			{
				p.SetState(148)
				p.Match(neonParserCOMMA)
			}
			{
				p.SetState(149)
				p.expr(0)
			}

			p.SetState(154)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(157)
		p.Match(neonParserRb)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) NUMTYPE() antlr.TerminalNode {
	return s.GetToken(neonParserNUMTYPE, 0)
}

func (s *TypeContext) STRTYPE() antlr.TerminalNode {
	return s.GetToken(neonParserSTRTYPE, 0)
}

func (s *TypeContext) BOOLTYPE() antlr.TerminalNode {
	return s.GetToken(neonParserBOOLTYPE, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *neonParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, neonParserRULE_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArithopContext is an interface to support dynamic dispatch.
type IArithopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithopContext differentiates from other interfaces.
	IsArithopContext()
}

type ArithopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithopContext() *ArithopContext {
	var p = new(ArithopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_arithop
	return p
}

func (*ArithopContext) IsArithopContext() {}

func NewArithopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithopContext {
	var p = new(ArithopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_arithop

	return p
}

func (s *ArithopContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithopContext) ADD_SUB() antlr.TerminalNode {
	return s.GetToken(neonParserADD_SUB, 0)
}

func (s *ArithopContext) MDM() antlr.TerminalNode {
	return s.GetToken(neonParserMDM, 0)
}

func (s *ArithopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterArithop(s)
	}
}

func (s *ArithopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitArithop(s)
	}
}

func (p *neonParser) Arithop() (localctx IArithopContext) {
	this := p
	_ = this

	localctx = NewArithopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, neonParserRULE_arithop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		_la = p.GetTokenStream().LA(1)

		if !(_la == neonParserADD_SUB || _la == neonParserMDM) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICompopContext is an interface to support dynamic dispatch.
type ICompopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompopContext differentiates from other interfaces.
	IsCompopContext()
}

type CompopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompopContext() *CompopContext {
	var p = new(CompopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_compop
	return p
}

func (*CompopContext) IsCompopContext() {}

func NewCompopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompopContext {
	var p = new(CompopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_compop

	return p
}

func (s *CompopContext) GetParser() antlr.Parser { return s.parser }

func (s *CompopContext) EQUALITY() antlr.TerminalNode {
	return s.GetToken(neonParserEQUALITY, 0)
}

func (s *CompopContext) LESS() antlr.TerminalNode {
	return s.GetToken(neonParserLESS, 0)
}

func (s *CompopContext) GREATER() antlr.TerminalNode {
	return s.GetToken(neonParserGREATER, 0)
}

func (s *CompopContext) LESSEQ() antlr.TerminalNode {
	return s.GetToken(neonParserLESSEQ, 0)
}

func (s *CompopContext) GREATEREQ() antlr.TerminalNode {
	return s.GetToken(neonParserGREATEREQ, 0)
}

func (s *CompopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterCompop(s)
	}
}

func (s *CompopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitCompop(s)
	}
}

func (p *neonParser) Compop() (localctx ICompopContext) {
	this := p
	_ = this

	localctx = NewCompopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, neonParserRULE_compop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&805357568) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) Arithop() IArithopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithopContext)
}

func (s *OpContext) Compop() ICompopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompopContext)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitOp(s)
	}
}

func (p *neonParser) Op() (localctx IOpContext) {
	this := p
	_ = this

	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, neonParserRULE_op)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case neonParserADD_SUB, neonParserMDM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Arithop()
		}

	case neonParserEQUALITY, neonParserLESSEQ, neonParserGREATEREQ, neonParserLESS, neonParserGREATER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Compop()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncargContext is an interface to support dynamic dispatch.
type IFuncargContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncargContext differentiates from other interfaces.
	IsFuncargContext()
}

type FuncargContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncargContext() *FuncargContext {
	var p = new(FuncargContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_funcarg
	return p
}

func (*FuncargContext) IsFuncargContext() {}

func NewFuncargContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncargContext {
	var p = new(FuncargContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_funcarg

	return p
}

func (s *FuncargContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncargContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncargContext) ID() antlr.TerminalNode {
	return s.GetToken(neonParserID, 0)
}

func (s *FuncargContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncargContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncargContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterFuncarg(s)
	}
}

func (s *FuncargContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitFuncarg(s)
	}
}

func (p *neonParser) Funcarg() (localctx IFuncargContext) {
	this := p
	_ = this

	localctx = NewFuncargContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, neonParserRULE_funcarg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Type_()
	}
	{
		p.SetState(170)
		p.Match(neonParserID)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclContext) ID() antlr.TerminalNode {
	return s.GetToken(neonParserID, 0)
}

func (s *DeclContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(neonParserEQUAL, 0)
}

func (s *DeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *neonParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, neonParserRULE_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Type_()
	}
	{
		p.SetState(173)
		p.Match(neonParserID)
	}
	{
		p.SetState(174)
		p.Match(neonParserEQUAL)
	}
	{
		p.SetState(175)
		p.expr(0)
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(neonParserID, 0)
}

func (s *AssignContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(neonParserEQUAL, 0)
}

func (s *AssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *neonParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, neonParserRULE_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(neonParserID)
	}
	{
		p.SetState(178)
		p.Match(neonParserEQUAL)
	}
	{
		p.SetState(179)
		p.expr(0)
	}

	return localctx
}

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_return
	return p
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(neonParserRETURN, 0)
}

func (s *ReturnContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (p *neonParser) Return_() (localctx IReturnContext) {
	this := p
	_ = this

	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, neonParserRULE_return)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(neonParserRETURN)
	}
	{
		p.SetState(182)
		p.expr(0)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = neonParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = neonParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) BANG() antlr.TerminalNode {
	return s.GetToken(neonParserBANG, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Funccall() IFunccallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunccallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunccallContext)
}

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(neonParserINT, 0)
}

func (s *ExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(neonParserBOOL, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(neonParserSTRING, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(neonParserID, 0)
}

func (s *ExprContext) Op() IOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(neonListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *neonParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *neonParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, neonParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(185)
			p.Match(neonParserBANG)
		}
		{
			p.SetState(186)
			p.expr(6)
		}

	case 2:
		{
			p.SetState(187)
			p.Funccall()
		}

	case 3:
		{
			p.SetState(188)
			p.Match(neonParserINT)
		}

	case 4:
		{
			p.SetState(189)
			p.Match(neonParserBOOL)
		}

	case 5:
		{
			p.SetState(190)
			p.Match(neonParserSTRING)
		}

	case 6:
		{
			p.SetState(191)
			p.Match(neonParserID)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, neonParserRULE_expr)
			p.SetState(194)

			if !(p.Precpred(p.GetParserRuleContext(), 7)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
			}
			{
				p.SetState(195)
				p.Op()
			}
			{
				p.SetState(196)
				p.expr(8)
			}

		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

func (p *neonParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *neonParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
