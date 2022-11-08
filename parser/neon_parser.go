// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Neon

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

type NeonParser struct {
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
		"", "';'", "'('", "')'", "'{'", "'}'", "','", "'number'", "'string'",
		"'bool'", "'='", "'!'", "'void'", "", "", "", "", "'def'", "'return'",
		"'if'", "'elif'", "'else'", "'while'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "ADD_SUB", "MDM",
		"COMP", "BOOL", "DEF", "RETURN", "IF", "ELIF", "ELSE", "WHILE", "STRING",
		"ID", "INT", "WS",
	}
	staticData.ruleNames = []string{
		"program", "stat", "if", "elif", "else", "while", "func", "funccall",
		"funcarg", "decl", "assign", "return", "expr", "type",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 210, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 3, 0, 31, 8,
		0, 5, 0, 33, 8, 0, 10, 0, 12, 0, 36, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 44, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 50, 8, 1, 3, 1, 52,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 60, 8, 2, 10, 2, 12, 2,
		63, 9, 2, 1, 2, 1, 2, 5, 2, 67, 8, 2, 10, 2, 12, 2, 70, 9, 2, 1, 2, 3,
		2, 73, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 81, 8, 3, 10, 3,
		12, 3, 84, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 91, 8, 4, 10, 4, 12,
		4, 94, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 104,
		8, 5, 10, 5, 12, 5, 107, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 5, 6, 118, 8, 6, 10, 6, 12, 6, 121, 9, 6, 3, 6, 123, 8, 6,
		1, 6, 3, 6, 126, 8, 6, 1, 6, 1, 6, 5, 6, 130, 8, 6, 10, 6, 12, 6, 133,
		9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 142, 8, 7, 10, 7,
		12, 7, 145, 9, 7, 3, 7, 147, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 157, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 171, 8, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 192, 8, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 203, 8, 12, 10, 12,
		12, 12, 206, 9, 12, 1, 13, 1, 13, 1, 13, 0, 1, 24, 14, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 0, 1, 2, 0, 7, 9, 12, 12, 227, 0, 34, 1,
		0, 0, 0, 2, 51, 1, 0, 0, 0, 4, 53, 1, 0, 0, 0, 6, 74, 1, 0, 0, 0, 8, 87,
		1, 0, 0, 0, 10, 97, 1, 0, 0, 0, 12, 110, 1, 0, 0, 0, 14, 136, 1, 0, 0,
		0, 16, 156, 1, 0, 0, 0, 18, 170, 1, 0, 0, 0, 20, 172, 1, 0, 0, 0, 22, 176,
		1, 0, 0, 0, 24, 191, 1, 0, 0, 0, 26, 207, 1, 0, 0, 0, 28, 31, 3, 2, 1,
		0, 29, 31, 3, 12, 6, 0, 30, 28, 1, 0, 0, 0, 30, 29, 1, 0, 0, 0, 31, 33,
		1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0,
		34, 35, 1, 0, 0, 0, 35, 37, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 38, 5,
		0, 0, 1, 38, 1, 1, 0, 0, 0, 39, 44, 3, 18, 9, 0, 40, 44, 3, 20, 10, 0,
		41, 44, 3, 14, 7, 0, 42, 44, 3, 22, 11, 0, 43, 39, 1, 0, 0, 0, 43, 40,
		1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0,
		45, 46, 5, 1, 0, 0, 46, 52, 1, 0, 0, 0, 47, 50, 3, 4, 2, 0, 48, 50, 3,
		10, 5, 0, 49, 47, 1, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51,
		43, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 3, 1, 0, 0, 0, 53, 54, 5, 19, 0,
		0, 54, 55, 5, 2, 0, 0, 55, 56, 3, 24, 12, 0, 56, 57, 5, 3, 0, 0, 57, 61,
		5, 4, 0, 0, 58, 60, 3, 2, 1, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0,
		61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64, 1, 0, 0, 0, 63, 61, 1,
		0, 0, 0, 64, 68, 5, 5, 0, 0, 65, 67, 3, 6, 3, 0, 66, 65, 1, 0, 0, 0, 67,
		70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 72, 1, 0, 0,
		0, 70, 68, 1, 0, 0, 0, 71, 73, 3, 8, 4, 0, 72, 71, 1, 0, 0, 0, 72, 73,
		1, 0, 0, 0, 73, 5, 1, 0, 0, 0, 74, 75, 5, 20, 0, 0, 75, 76, 5, 2, 0, 0,
		76, 77, 3, 24, 12, 0, 77, 78, 5, 3, 0, 0, 78, 82, 5, 4, 0, 0, 79, 81, 3,
		2, 1, 0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82,
		83, 1, 0, 0, 0, 83, 85, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 5, 0,
		0, 86, 7, 1, 0, 0, 0, 87, 88, 5, 21, 0, 0, 88, 92, 5, 4, 0, 0, 89, 91,
		3, 2, 1, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0,
		92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 5,
		5, 0, 0, 96, 9, 1, 0, 0, 0, 97, 98, 5, 22, 0, 0, 98, 99, 5, 2, 0, 0, 99,
		100, 3, 24, 12, 0, 100, 101, 5, 3, 0, 0, 101, 105, 5, 4, 0, 0, 102, 104,
		3, 2, 1, 0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0,
		0, 0, 105, 106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0,
		108, 109, 5, 5, 0, 0, 109, 11, 1, 0, 0, 0, 110, 111, 5, 17, 0, 0, 111,
		112, 3, 26, 13, 0, 112, 125, 5, 24, 0, 0, 113, 122, 5, 2, 0, 0, 114, 119,
		3, 16, 8, 0, 115, 116, 5, 6, 0, 0, 116, 118, 3, 16, 8, 0, 117, 115, 1,
		0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0,
		0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 114, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 5, 3, 0, 0, 125, 113,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 131, 5, 4,
		0, 0, 128, 130, 3, 2, 1, 0, 129, 128, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0,
		131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 1, 0, 0, 0, 133,
		131, 1, 0, 0, 0, 134, 135, 5, 5, 0, 0, 135, 13, 1, 0, 0, 0, 136, 137, 5,
		24, 0, 0, 137, 146, 5, 2, 0, 0, 138, 143, 3, 24, 12, 0, 139, 140, 5, 6,
		0, 0, 140, 142, 3, 24, 12, 0, 141, 139, 1, 0, 0, 0, 142, 145, 1, 0, 0,
		0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145,
		143, 1, 0, 0, 0, 146, 138, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148,
		1, 0, 0, 0, 148, 149, 5, 3, 0, 0, 149, 15, 1, 0, 0, 0, 150, 151, 5, 7,
		0, 0, 151, 157, 5, 24, 0, 0, 152, 153, 5, 8, 0, 0, 153, 157, 5, 24, 0,
		0, 154, 155, 5, 9, 0, 0, 155, 157, 5, 24, 0, 0, 156, 150, 1, 0, 0, 0, 156,
		152, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 17, 1, 0, 0, 0, 158, 159, 5,
		7, 0, 0, 159, 160, 5, 24, 0, 0, 160, 161, 5, 10, 0, 0, 161, 171, 3, 24,
		12, 0, 162, 163, 5, 8, 0, 0, 163, 164, 5, 24, 0, 0, 164, 165, 5, 10, 0,
		0, 165, 171, 3, 24, 12, 0, 166, 167, 5, 9, 0, 0, 167, 168, 5, 24, 0, 0,
		168, 169, 5, 10, 0, 0, 169, 171, 3, 24, 12, 0, 170, 158, 1, 0, 0, 0, 170,
		162, 1, 0, 0, 0, 170, 166, 1, 0, 0, 0, 171, 19, 1, 0, 0, 0, 172, 173, 5,
		24, 0, 0, 173, 174, 5, 10, 0, 0, 174, 175, 3, 24, 12, 0, 175, 21, 1, 0,
		0, 0, 176, 177, 5, 18, 0, 0, 177, 178, 3, 24, 12, 0, 178, 23, 1, 0, 0,
		0, 179, 180, 6, 12, -1, 0, 180, 181, 5, 11, 0, 0, 181, 192, 3, 24, 12,
		7, 182, 192, 3, 14, 7, 0, 183, 192, 5, 25, 0, 0, 184, 192, 5, 16, 0, 0,
		185, 192, 5, 23, 0, 0, 186, 192, 5, 24, 0, 0, 187, 188, 5, 2, 0, 0, 188,
		189, 3, 24, 12, 0, 189, 190, 5, 3, 0, 0, 190, 192, 1, 0, 0, 0, 191, 179,
		1, 0, 0, 0, 191, 182, 1, 0, 0, 0, 191, 183, 1, 0, 0, 0, 191, 184, 1, 0,
		0, 0, 191, 185, 1, 0, 0, 0, 191, 186, 1, 0, 0, 0, 191, 187, 1, 0, 0, 0,
		192, 204, 1, 0, 0, 0, 193, 194, 10, 10, 0, 0, 194, 195, 5, 14, 0, 0, 195,
		203, 3, 24, 12, 11, 196, 197, 10, 9, 0, 0, 197, 198, 5, 13, 0, 0, 198,
		203, 3, 24, 12, 10, 199, 200, 10, 8, 0, 0, 200, 201, 5, 15, 0, 0, 201,
		203, 3, 24, 12, 9, 202, 193, 1, 0, 0, 0, 202, 196, 1, 0, 0, 0, 202, 199,
		1, 0, 0, 0, 203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205, 1, 0,
		0, 0, 205, 25, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 208, 7, 0, 0, 0,
		208, 27, 1, 0, 0, 0, 22, 30, 34, 43, 49, 51, 61, 68, 72, 82, 92, 105, 119,
		122, 125, 131, 143, 146, 156, 170, 191, 202, 204,
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

// NeonParserInit initializes any static state used to implement NeonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNeonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NeonParserInit() {
	staticData := &neonParserStaticData
	staticData.once.Do(neonParserInit)
}

// NewNeonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNeonParser(input antlr.TokenStream) *NeonParser {
	NeonParserInit()
	this := new(NeonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &neonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// NeonParser tokens.
const (
	NeonParserEOF     = antlr.TokenEOF
	NeonParserT__0    = 1
	NeonParserT__1    = 2
	NeonParserT__2    = 3
	NeonParserT__3    = 4
	NeonParserT__4    = 5
	NeonParserT__5    = 6
	NeonParserT__6    = 7
	NeonParserT__7    = 8
	NeonParserT__8    = 9
	NeonParserT__9    = 10
	NeonParserT__10   = 11
	NeonParserT__11   = 12
	NeonParserADD_SUB = 13
	NeonParserMDM     = 14
	NeonParserCOMP    = 15
	NeonParserBOOL    = 16
	NeonParserDEF     = 17
	NeonParserRETURN  = 18
	NeonParserIF      = 19
	NeonParserELIF    = 20
	NeonParserELSE    = 21
	NeonParserWHILE   = 22
	NeonParserSTRING  = 23
	NeonParserID      = 24
	NeonParserINT     = 25
	NeonParserWS      = 26
)

// NeonParser rules.
const (
	NeonParserRULE_program  = 0
	NeonParserRULE_stat     = 1
	NeonParserRULE_if       = 2
	NeonParserRULE_elif     = 3
	NeonParserRULE_else     = 4
	NeonParserRULE_while    = 5
	NeonParserRULE_func     = 6
	NeonParserRULE_funccall = 7
	NeonParserRULE_funcarg  = 8
	NeonParserRULE_decl     = 9
	NeonParserRULE_assign   = 10
	NeonParserRULE_return   = 11
	NeonParserRULE_expr     = 12
	NeonParserRULE_type     = 13
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
	p.RuleIndex = NeonParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(NeonParserEOF, 0)
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
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *NeonParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NeonParserRULE_program)
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21889920) != 0 {
		p.SetState(30)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NeonParserT__6, NeonParserT__7, NeonParserT__8, NeonParserRETURN, NeonParserIF, NeonParserWHILE, NeonParserID:
			{
				p.SetState(28)
				p.Stat()
			}

		case NeonParserDEF:
			{
				p.SetState(29)
				p.Func_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
		p.Match(NeonParserEOF)
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
	p.RuleIndex = NeonParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

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
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *NeonParser) Stat() (localctx IStatContext) {
	this := p
	_ = this

	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NeonParserRULE_stat)

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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NeonParserT__6, NeonParserT__7, NeonParserT__8, NeonParserRETURN, NeonParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(39)
				p.Decl()
			}

		case 2:
			{
				p.SetState(40)
				p.Assign()
			}

		case 3:
			{
				p.SetState(41)
				p.Funccall()
			}

		case 4:
			{
				p.SetState(42)
				p.Return_()
			}

		}
		{
			p.SetState(45)
			p.Match(NeonParserT__0)
		}

	case NeonParserIF, NeonParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(49)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NeonParserIF:
			{
				p.SetState(47)
				p.If_()
			}

		case NeonParserWHILE:
			{
				p.SetState(48)
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
	p.RuleIndex = NeonParserRULE_if
	return p
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(NeonParserIF, 0)
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
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitIf(s)
	}
}

func (p *NeonParser) If_() (localctx IIfContext) {
	this := p
	_ = this

	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NeonParserRULE_if)
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
		p.SetState(53)
		p.Match(NeonParserIF)
	}
	{
		p.SetState(54)
		p.Match(NeonParserT__1)
	}
	{
		p.SetState(55)
		p.expr(0)
	}
	{
		p.SetState(56)
		p.Match(NeonParserT__2)
	}
	{
		p.SetState(57)
		p.Match(NeonParserT__3)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21758848) != 0 {
		{
			p.SetState(58)
			p.Stat()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(NeonParserT__4)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NeonParserELIF {
		{
			p.SetState(65)
			p.Elif()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NeonParserELSE {
		{
			p.SetState(71)
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
	p.RuleIndex = NeonParserRULE_elif
	return p
}

func (*ElifContext) IsElifContext() {}

func NewElifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifContext {
	var p = new(ElifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_elif

	return p
}

func (s *ElifContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifContext) ELIF() antlr.TerminalNode {
	return s.GetToken(NeonParserELIF, 0)
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
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterElif(s)
	}
}

func (s *ElifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitElif(s)
	}
}

func (p *NeonParser) Elif() (localctx IElifContext) {
	this := p
	_ = this

	localctx = NewElifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NeonParserRULE_elif)
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
		p.SetState(74)
		p.Match(NeonParserELIF)
	}
	{
		p.SetState(75)
		p.Match(NeonParserT__1)
	}
	{
		p.SetState(76)
		p.expr(0)
	}
	{
		p.SetState(77)
		p.Match(NeonParserT__2)
	}
	{
		p.SetState(78)
		p.Match(NeonParserT__3)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21758848) != 0 {
		{
			p.SetState(79)
			p.Stat()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
		p.Match(NeonParserT__4)
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
	p.RuleIndex = NeonParserRULE_else
	return p
}

func (*ElseContext) IsElseContext() {}

func NewElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseContext {
	var p = new(ElseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_else

	return p
}

func (s *ElseContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NeonParserELSE, 0)
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
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterElse(s)
	}
}

func (s *ElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitElse(s)
	}
}

func (p *NeonParser) Else_() (localctx IElseContext) {
	this := p
	_ = this

	localctx = NewElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NeonParserRULE_else)
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
		p.SetState(87)
		p.Match(NeonParserELSE)
	}
	{
		p.SetState(88)
		p.Match(NeonParserT__3)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21758848) != 0 {
		{
			p.SetState(89)
			p.Stat()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(NeonParserT__4)
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
	p.RuleIndex = NeonParserRULE_while
	return p
}

func (*WhileContext) IsWhileContext() {}

func NewWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileContext {
	var p = new(WhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_while

	return p
}

func (s *WhileContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(NeonParserWHILE, 0)
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
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (p *NeonParser) While() (localctx IWhileContext) {
	this := p
	_ = this

	localctx = NewWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NeonParserRULE_while)
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
		p.SetState(97)
		p.Match(NeonParserWHILE)
	}
	{
		p.SetState(98)
		p.Match(NeonParserT__1)
	}
	{
		p.SetState(99)
		p.expr(0)
	}
	{
		p.SetState(100)
		p.Match(NeonParserT__2)
	}
	{
		p.SetState(101)
		p.Match(NeonParserT__3)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21758848) != 0 {
		{
			p.SetState(102)
			p.Stat()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(108)
		p.Match(NeonParserT__4)
	}

	return localctx
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFunction returns the function token.
	GetFunction() antlr.Token

	// SetFunction sets the function token.
	SetFunction(antlr.Token)

	// GetFuncType returns the funcType rule contexts.
	GetFuncType() ITypeContext

	// SetFuncType sets the funcType rule contexts.
	SetFuncType(ITypeContext)

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	funcType ITypeContext
	function antlr.Token
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NeonParserRULE_func
	return p
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) GetFunction() antlr.Token { return s.function }

func (s *FuncContext) SetFunction(v antlr.Token) { s.function = v }

func (s *FuncContext) GetFuncType() ITypeContext { return s.funcType }

func (s *FuncContext) SetFuncType(v ITypeContext) { s.funcType = v }

func (s *FuncContext) DEF() antlr.TerminalNode {
	return s.GetToken(NeonParserDEF, 0)
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

func (s *FuncContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
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

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (p *NeonParser) Func_() (localctx IFuncContext) {
	this := p
	_ = this

	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NeonParserRULE_func)
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
		p.SetState(110)
		p.Match(NeonParserDEF)
	}
	{
		p.SetState(111)

		var _x = p.Type_()

		localctx.(*FuncContext).funcType = _x
	}
	{
		p.SetState(112)

		var _m = p.Match(NeonParserID)

		localctx.(*FuncContext).function = _m
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NeonParserT__1 {
		{
			p.SetState(113)
			p.Match(NeonParserT__1)
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0 {
			{
				p.SetState(114)
				p.Funcarg()
			}
			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == NeonParserT__5 {
				{
					p.SetState(115)
					p.Match(NeonParserT__5)
				}
				{
					p.SetState(116)
					p.Funcarg()
				}

				p.SetState(121)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(124)
			p.Match(NeonParserT__2)
		}

	}
	{
		p.SetState(127)
		p.Match(NeonParserT__3)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21758848) != 0 {
		{
			p.SetState(128)
			p.Stat()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(NeonParserT__4)
	}

	return localctx
}

// IFunccallContext is an interface to support dynamic dispatch.
type IFunccallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFunction returns the function token.
	GetFunction() antlr.Token

	// SetFunction sets the function token.
	SetFunction(antlr.Token)

	// IsFunccallContext differentiates from other interfaces.
	IsFunccallContext()
}

type FunccallContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	function antlr.Token
}

func NewEmptyFunccallContext() *FunccallContext {
	var p = new(FunccallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NeonParserRULE_funccall
	return p
}

func (*FunccallContext) IsFunccallContext() {}

func NewFunccallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunccallContext {
	var p = new(FunccallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_funccall

	return p
}

func (s *FunccallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunccallContext) GetFunction() antlr.Token { return s.function }

func (s *FunccallContext) SetFunction(v antlr.Token) { s.function = v }

func (s *FunccallContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
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

func (s *FunccallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunccallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunccallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterFunccall(s)
	}
}

func (s *FunccallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitFunccall(s)
	}
}

func (p *NeonParser) Funccall() (localctx IFunccallContext) {
	this := p
	_ = this

	localctx = NewFunccallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NeonParserRULE_funccall)
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
		p.SetState(136)

		var _m = p.Match(NeonParserID)

		localctx.(*FunccallContext).function = _m
	}
	{
		p.SetState(137)
		p.Match(NeonParserT__1)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&58787844) != 0 {
		{
			p.SetState(138)
			p.expr(0)
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NeonParserT__5 {
			{
				p.SetState(139)
				p.Match(NeonParserT__5)
			}
			{
				p.SetState(140)
				p.expr(0)
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(148)
		p.Match(NeonParserT__2)
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
	p.RuleIndex = NeonParserRULE_funcarg
	return p
}

func (*FuncargContext) IsFuncargContext() {}

func NewFuncargContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncargContext {
	var p = new(FuncargContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_funcarg

	return p
}

func (s *FuncargContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncargContext) CopyFrom(ctx *FuncargContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FuncargContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncargContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StrArgContext struct {
	*FuncargContext
}

func NewStrArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrArgContext {
	var p = new(StrArgContext)

	p.FuncargContext = NewEmptyFuncargContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FuncargContext))

	return p
}

func (s *StrArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrArgContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
}

func (s *StrArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterStrArg(s)
	}
}

func (s *StrArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitStrArg(s)
	}
}

type NumArgContext struct {
	*FuncargContext
}

func NewNumArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumArgContext {
	var p = new(NumArgContext)

	p.FuncargContext = NewEmptyFuncargContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FuncargContext))

	return p
}

func (s *NumArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumArgContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
}

func (s *NumArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterNumArg(s)
	}
}

func (s *NumArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitNumArg(s)
	}
}

type BoolArgContext struct {
	*FuncargContext
}

func NewBoolArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolArgContext {
	var p = new(BoolArgContext)

	p.FuncargContext = NewEmptyFuncargContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FuncargContext))

	return p
}

func (s *BoolArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolArgContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
}

func (s *BoolArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterBoolArg(s)
	}
}

func (s *BoolArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitBoolArg(s)
	}
}

func (p *NeonParser) Funcarg() (localctx IFuncargContext) {
	this := p
	_ = this

	localctx = NewFuncargContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NeonParserRULE_funcarg)

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

	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NeonParserT__6:
		localctx = NewNumArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Match(NeonParserT__6)
		}
		{
			p.SetState(151)
			p.Match(NeonParserID)
		}

	case NeonParserT__7:
		localctx = NewStrArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Match(NeonParserT__7)
		}
		{
			p.SetState(153)
			p.Match(NeonParserID)
		}

	case NeonParserT__8:
		localctx = NewBoolArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.Match(NeonParserT__8)
		}
		{
			p.SetState(155)
			p.Match(NeonParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = NeonParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) CopyFrom(ctx *DeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumVarContext struct {
	*DeclContext
	var_ antlr.Token
}

func NewNumVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumVarContext {
	var p = new(NumVarContext)

	p.DeclContext = NewEmptyDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclContext))

	return p
}

func (s *NumVarContext) GetVar_() antlr.Token { return s.var_ }

func (s *NumVarContext) SetVar_(v antlr.Token) { s.var_ = v }

func (s *NumVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumVarContext) Expr() IExprContext {
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

func (s *NumVarContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
}

func (s *NumVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterNumVar(s)
	}
}

func (s *NumVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitNumVar(s)
	}
}

type StrVarContext struct {
	*DeclContext
	var_ antlr.Token
}

func NewStrVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrVarContext {
	var p = new(StrVarContext)

	p.DeclContext = NewEmptyDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclContext))

	return p
}

func (s *StrVarContext) GetVar_() antlr.Token { return s.var_ }

func (s *StrVarContext) SetVar_(v antlr.Token) { s.var_ = v }

func (s *StrVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrVarContext) Expr() IExprContext {
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

func (s *StrVarContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
}

func (s *StrVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterStrVar(s)
	}
}

func (s *StrVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitStrVar(s)
	}
}

type BoolVarContext struct {
	*DeclContext
	var_ antlr.Token
}

func NewBoolVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolVarContext {
	var p = new(BoolVarContext)

	p.DeclContext = NewEmptyDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclContext))

	return p
}

func (s *BoolVarContext) GetVar_() antlr.Token { return s.var_ }

func (s *BoolVarContext) SetVar_(v antlr.Token) { s.var_ = v }

func (s *BoolVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolVarContext) Expr() IExprContext {
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

func (s *BoolVarContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
}

func (s *BoolVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterBoolVar(s)
	}
}

func (s *BoolVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitBoolVar(s)
	}
}

func (p *NeonParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NeonParserRULE_decl)

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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NeonParserT__6:
		localctx = NewNumVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Match(NeonParserT__6)
		}
		{
			p.SetState(159)

			var _m = p.Match(NeonParserID)

			localctx.(*NumVarContext).var_ = _m
		}
		{
			p.SetState(160)
			p.Match(NeonParserT__9)
		}
		{
			p.SetState(161)
			p.expr(0)
		}

	case NeonParserT__7:
		localctx = NewStrVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(NeonParserT__7)
		}
		{
			p.SetState(163)

			var _m = p.Match(NeonParserID)

			localctx.(*StrVarContext).var_ = _m
		}
		{
			p.SetState(164)
			p.Match(NeonParserT__9)
		}
		{
			p.SetState(165)
			p.expr(0)
		}

	case NeonParserT__8:
		localctx = NewBoolVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Match(NeonParserT__8)
		}
		{
			p.SetState(167)

			var _m = p.Match(NeonParserID)

			localctx.(*BoolVarContext).var_ = _m
		}
		{
			p.SetState(168)
			p.Match(NeonParserT__9)
		}
		{
			p.SetState(169)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVar_ returns the var_ token.
	GetVar_() antlr.Token

	// SetVar_ sets the var_ token.
	SetVar_(antlr.Token)

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	var_   antlr.Token
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NeonParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) GetVar_() antlr.Token { return s.var_ }

func (s *AssignContext) SetVar_(v antlr.Token) { s.var_ = v }

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

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *NeonParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NeonParserRULE_assign)

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

		var _m = p.Match(NeonParserID)

		localctx.(*AssignContext).var_ = _m
	}
	{
		p.SetState(173)
		p.Match(NeonParserT__9)
	}
	{
		p.SetState(174)
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
	p.RuleIndex = NeonParserRULE_return
	return p
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(NeonParserRETURN, 0)
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
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (p *NeonParser) Return_() (localctx IReturnContext) {
	this := p
	_ = this

	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NeonParserRULE_return)

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
		p.SetState(176)
		p.Match(NeonParserRETURN)
	}
	{
		p.SetState(177)
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
	p.RuleIndex = NeonParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncCallContext struct {
	*ExprContext
}

func NewFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallContext {
	var p = new(FuncCallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) Funccall() IFunccallContext {
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

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

type IdentifierContext struct {
	*ExprContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(NeonParserID, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

type BoolContext struct {
	*ExprContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(NeonParserBOOL, 0)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitBool(s)
	}
}

type AddSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
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

func (s *AddSubContext) Expr(i int) IExprContext {
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

func (s *AddSubContext) ADD_SUB() antlr.TerminalNode {
	return s.GetToken(NeonParserADD_SUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type ComparisonContext struct {
	*ExprContext
	op antlr.Token
}

func NewComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonContext {
	var p = new(ComparisonContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ComparisonContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) AllExpr() []IExprContext {
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

func (s *ComparisonContext) Expr(i int) IExprContext {
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

func (s *ComparisonContext) COMP() antlr.TerminalNode {
	return s.GetToken(NeonParserCOMP, 0)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitComparison(s)
	}
}

type MDMContext struct {
	*ExprContext
	op antlr.Token
}

func NewMDMContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MDMContext {
	var p = new(MDMContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MDMContext) GetOp() antlr.Token { return s.op }

func (s *MDMContext) SetOp(v antlr.Token) { s.op = v }

func (s *MDMContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MDMContext) AllExpr() []IExprContext {
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

func (s *MDMContext) Expr(i int) IExprContext {
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

func (s *MDMContext) MDM() antlr.TerminalNode {
	return s.GetToken(NeonParserMDM, 0)
}

func (s *MDMContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterMDM(s)
	}
}

func (s *MDMContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitMDM(s)
	}
}

type NotExprContext struct {
	*ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
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

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

type StringContext struct {
	*ExprContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(NeonParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitString(s)
	}
}

type NoClueContext struct {
	*ExprContext
}

func NewNoClueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NoClueContext {
	var p = new(NoClueContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NoClueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoClueContext) Expr() IExprContext {
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

func (s *NoClueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterNoClue(s)
	}
}

func (s *NoClueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitNoClue(s)
	}
}

type IntContext struct {
	*ExprContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(NeonParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitInt(s)
	}
}

func (p *NeonParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *NeonParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, NeonParserRULE_expr, _p)

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
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(180)
			p.Match(NeonParserT__10)
		}
		{
			p.SetState(181)
			p.expr(7)
		}

	case 2:
		localctx = NewFuncCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(182)
			p.Funccall()
		}

	case 3:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(183)
			p.Match(NeonParserINT)
		}

	case 4:
		localctx = NewBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(184)
			p.Match(NeonParserBOOL)
		}

	case 5:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(185)
			p.Match(NeonParserSTRING)
		}

	case 6:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(186)
			p.Match(NeonParserID)
		}

	case 7:
		localctx = NewNoClueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(187)
			p.Match(NeonParserT__1)
		}
		{
			p.SetState(188)
			p.expr(0)
		}
		{
			p.SetState(189)
			p.Match(NeonParserT__2)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMDMContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NeonParserRULE_expr)
				p.SetState(193)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(194)

					var _m = p.Match(NeonParserMDM)

					localctx.(*MDMContext).op = _m
				}
				{
					p.SetState(195)
					p.expr(11)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NeonParserRULE_expr)
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(197)

					var _m = p.Match(NeonParserADD_SUB)

					localctx.(*AddSubContext).op = _m
				}
				{
					p.SetState(198)
					p.expr(10)
				}

			case 3:
				localctx = NewComparisonContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NeonParserRULE_expr)
				p.SetState(199)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(200)

					var _m = p.Match(NeonParserCOMP)

					localctx.(*ComparisonContext).op = _m
				}
				{
					p.SetState(201)
					p.expr(9)
				}

			}

		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.RuleIndex = NeonParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeonParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeonListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *NeonParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NeonParserRULE_type)
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
		p.SetState(207)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4992) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *NeonParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NeonParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
