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
		"", "'number'", "'string'", "'bool'", "'def'", "'!='", "';'", "'<='",
		"'>='", "", "", "", "", "", "", "','", "'('", "')'", "'{'", "'}'", "'='",
		"'<'", "'>'",
	}
	staticData.symbolicNames = []string{
		"", "NUMTYPE", "STRTYPE", "BOOLTYPE", "DEF", "NOTEQUAL", "SEMI", "LESSEQ",
		"GREATEREQ", "STRING", "ID", "INT", "WS", "ADD_SUB", "MUL_DIV", "COMMA",
		"Lb", "Rb", "Lc", "Rc", "EQUAL", "LESS", "GREATER", "BOOL",
	}
	staticData.ruleNames = []string{
		"program", "stat", "func", "type", "arithop", "compop", "op", "funcarg",
		"decl", "assign", "expr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 97, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 3, 0, 25, 8, 0, 5, 0, 27, 8, 0, 10, 0, 12, 0, 30, 9, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 36, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 5, 2, 44, 8, 2, 10, 2, 12, 2, 47, 9, 2, 1, 2, 1, 2, 1, 2, 5, 2, 52,
		8, 2, 10, 2, 12, 2, 55, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 3, 6, 67, 8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 86, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 92, 8, 10, 10, 10, 12,
		10, 95, 9, 10, 1, 10, 0, 1, 20, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 0, 3, 1, 0, 1, 3, 1, 0, 13, 14, 2, 0, 7, 8, 21, 22, 95, 0, 28, 1, 0,
		0, 0, 2, 35, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 58, 1, 0, 0, 0, 8, 60, 1,
		0, 0, 0, 10, 62, 1, 0, 0, 0, 12, 66, 1, 0, 0, 0, 14, 68, 1, 0, 0, 0, 16,
		71, 1, 0, 0, 0, 18, 76, 1, 0, 0, 0, 20, 85, 1, 0, 0, 0, 22, 25, 3, 2, 1,
		0, 23, 25, 3, 4, 2, 0, 24, 22, 1, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 27,
		1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0,
		28, 29, 1, 0, 0, 0, 29, 31, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 32, 5,
		0, 0, 1, 32, 1, 1, 0, 0, 0, 33, 36, 3, 16, 8, 0, 34, 36, 3, 18, 9, 0, 35,
		33, 1, 0, 0, 0, 35, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 5, 6, 0,
		0, 38, 3, 1, 0, 0, 0, 39, 40, 5, 4, 0, 0, 40, 41, 5, 10, 0, 0, 41, 45,
		5, 16, 0, 0, 42, 44, 3, 14, 7, 0, 43, 42, 1, 0, 0, 0, 44, 47, 1, 0, 0,
		0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 48, 1, 0, 0, 0, 47, 45,
		1, 0, 0, 0, 48, 49, 5, 17, 0, 0, 49, 53, 5, 18, 0, 0, 50, 52, 3, 2, 1,
		0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54,
		1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 5, 19, 0, 0,
		57, 5, 1, 0, 0, 0, 58, 59, 7, 0, 0, 0, 59, 7, 1, 0, 0, 0, 60, 61, 7, 1,
		0, 0, 61, 9, 1, 0, 0, 0, 62, 63, 7, 2, 0, 0, 63, 11, 1, 0, 0, 0, 64, 67,
		3, 8, 4, 0, 65, 67, 3, 10, 5, 0, 66, 64, 1, 0, 0, 0, 66, 65, 1, 0, 0, 0,
		67, 13, 1, 0, 0, 0, 68, 69, 3, 6, 3, 0, 69, 70, 5, 10, 0, 0, 70, 15, 1,
		0, 0, 0, 71, 72, 3, 6, 3, 0, 72, 73, 5, 10, 0, 0, 73, 74, 5, 20, 0, 0,
		74, 75, 3, 20, 10, 0, 75, 17, 1, 0, 0, 0, 76, 77, 5, 10, 0, 0, 77, 78,
		5, 20, 0, 0, 78, 79, 3, 20, 10, 0, 79, 19, 1, 0, 0, 0, 80, 81, 6, 10, -1,
		0, 81, 86, 5, 11, 0, 0, 82, 86, 5, 23, 0, 0, 83, 86, 5, 9, 0, 0, 84, 86,
		5, 10, 0, 0, 85, 80, 1, 0, 0, 0, 85, 82, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0,
		85, 84, 1, 0, 0, 0, 86, 93, 1, 0, 0, 0, 87, 88, 10, 5, 0, 0, 88, 89, 3,
		12, 6, 0, 89, 90, 3, 20, 10, 6, 90, 92, 1, 0, 0, 0, 91, 87, 1, 0, 0, 0,
		92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 21, 1,
		0, 0, 0, 95, 93, 1, 0, 0, 0, 8, 24, 28, 35, 45, 53, 66, 85, 93,
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
	neonParserDEF       = 4
	neonParserNOTEQUAL  = 5
	neonParserSEMI      = 6
	neonParserLESSEQ    = 7
	neonParserGREATEREQ = 8
	neonParserSTRING    = 9
	neonParserID        = 10
	neonParserINT       = 11
	neonParserWS        = 12
	neonParserADD_SUB   = 13
	neonParserMUL_DIV   = 14
	neonParserCOMMA     = 15
	neonParserLb        = 16
	neonParserRb        = 17
	neonParserLc        = 18
	neonParserRc        = 19
	neonParserEQUAL     = 20
	neonParserLESS      = 21
	neonParserGREATER   = 22
	neonParserBOOL      = 23
)

// neonParser rules.
const (
	neonParserRULE_program = 0
	neonParserRULE_stat    = 1
	neonParserRULE_func    = 2
	neonParserRULE_type    = 3
	neonParserRULE_arithop = 4
	neonParserRULE_compop  = 5
	neonParserRULE_op      = 6
	neonParserRULE_funcarg = 7
	neonParserRULE_decl    = 8
	neonParserRULE_assign  = 9
	neonParserRULE_expr    = 10
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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1054) != 0 {
		p.SetState(24)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case neonParserNUMTYPE, neonParserSTRTYPE, neonParserBOOLTYPE, neonParserID:
			{
				p.SetState(22)
				p.Stat()
			}

		case neonParserDEF:
			{
				p.SetState(23)
				p.Func_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case neonParserNUMTYPE, neonParserSTRTYPE, neonParserBOOLTYPE:
		{
			p.SetState(33)
			p.Decl()
		}

	case neonParserID:
		{
			p.SetState(34)
			p.Assign()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(37)
		p.Match(neonParserSEMI)
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

func (s *FuncContext) Lb() antlr.TerminalNode {
	return s.GetToken(neonParserLb, 0)
}

func (s *FuncContext) Rb() antlr.TerminalNode {
	return s.GetToken(neonParserRb, 0)
}

func (s *FuncContext) Lc() antlr.TerminalNode {
	return s.GetToken(neonParserLc, 0)
}

func (s *FuncContext) Rc() antlr.TerminalNode {
	return s.GetToken(neonParserRc, 0)
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
	p.EnterRule(localctx, 4, neonParserRULE_func)
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
		p.SetState(39)
		p.Match(neonParserDEF)
	}
	{
		p.SetState(40)
		p.Match(neonParserID)
	}
	{
		p.SetState(41)
		p.Match(neonParserLb)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0 {
		{
			p.SetState(42)
			p.Funcarg()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
		p.Match(neonParserRb)
	}
	{
		p.SetState(49)
		p.Match(neonParserLc)
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1038) != 0 {
		{
			p.SetState(50)
			p.Stat()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(neonParserRc)
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
	p.EnterRule(localctx, 6, neonParserRULE_type)
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
		p.SetState(58)
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

func (s *ArithopContext) MUL_DIV() antlr.TerminalNode {
	return s.GetToken(neonParserMUL_DIV, 0)
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
	p.EnterRule(localctx, 8, neonParserRULE_arithop)
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
		p.SetState(60)
		_la = p.GetTokenStream().LA(1)

		if !(_la == neonParserADD_SUB || _la == neonParserMUL_DIV) {
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
	p.EnterRule(localctx, 10, neonParserRULE_compop)
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
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6291840) != 0) {
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
	p.EnterRule(localctx, 12, neonParserRULE_op)

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

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case neonParserADD_SUB, neonParserMUL_DIV:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Arithop()
		}

	case neonParserLESSEQ, neonParserGREATEREQ, neonParserLESS, neonParserGREATER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
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
	p.EnterRule(localctx, 14, neonParserRULE_funcarg)

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
		p.SetState(68)
		p.Type_()
	}
	{
		p.SetState(69)
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
	p.EnterRule(localctx, 16, neonParserRULE_decl)

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
		p.SetState(71)
		p.Type_()
	}
	{
		p.SetState(72)
		p.Match(neonParserID)
	}
	{
		p.SetState(73)
		p.Match(neonParserEQUAL)
	}
	{
		p.SetState(74)
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
	p.EnterRule(localctx, 18, neonParserRULE_assign)

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
		p.SetState(76)
		p.Match(neonParserID)
	}
	{
		p.SetState(77)
		p.Match(neonParserEQUAL)
	}
	{
		p.SetState(78)
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, neonParserRULE_expr, _p)

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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case neonParserINT:
		{
			p.SetState(81)
			p.Match(neonParserINT)
		}

	case neonParserBOOL:
		{
			p.SetState(82)
			p.Match(neonParserBOOL)
		}

	case neonParserSTRING:
		{
			p.SetState(83)
			p.Match(neonParserSTRING)
		}

	case neonParserID:
		{
			p.SetState(84)
			p.Match(neonParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, neonParserRULE_expr)
			p.SetState(87)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(88)
				p.Op()
			}
			{
				p.SetState(89)
				p.expr(6)
			}

		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

func (p *neonParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
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
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
