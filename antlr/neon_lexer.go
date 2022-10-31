// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type neonLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var neonlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func neonlexerLexerInit() {
	staticData := &neonlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'number'", "'string'", "'bool'", "'def'", "'!='", "';'", "'<='",
		"'>='", "", "", "", "", "", "", "','", "'('", "')'", "'{'", "'}'", "'='",
		"'<'", "'>'",
	}
	staticData.symbolicNames = []string{
		"", "NUMTYPE", "STRTYPE", "BOOLTYPE", "DEF", "NOTEQUAL", "SEMI", "LESSEQ",
		"GREATEREQ", "STRING", "ID", "INT", "WS", "ADD_SUB", "MUL_DIV", "COMMA",
		"Lb", "Rb", "Lc", "Rc", "EQUAL", "LESS", "GREATER",
	}
	staticData.ruleNames = []string{
		"NUMTYPE", "STRTYPE", "BOOLTYPE", "DEF", "NOTEQUAL", "SEMI", "LESSEQ",
		"GREATEREQ", "STRING", "ID", "INT", "WS", "ADD_SUB", "MUL_DIV", "COMMA",
		"Lb", "Rb", "Lc", "Rc", "EQUAL", "LESS", "GREATER",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 125, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 5, 8, 82, 8, 8, 10, 8, 12, 8, 85, 9, 8, 1, 8, 1, 8, 1,
		9, 4, 9, 90, 8, 9, 11, 9, 12, 9, 91, 1, 10, 4, 10, 95, 8, 10, 11, 10, 12,
		10, 96, 1, 11, 4, 11, 100, 8, 11, 11, 11, 12, 11, 101, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 0, 0, 22,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 1, 0, 4, 2, 0, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32,
		32, 2, 0, 43, 43, 45, 45, 2, 0, 42, 42, 47, 47, 128, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 1, 45, 1, 0, 0, 0, 3, 52, 1, 0, 0, 0, 5,
		59, 1, 0, 0, 0, 7, 64, 1, 0, 0, 0, 9, 68, 1, 0, 0, 0, 11, 71, 1, 0, 0,
		0, 13, 73, 1, 0, 0, 0, 15, 76, 1, 0, 0, 0, 17, 79, 1, 0, 0, 0, 19, 89,
		1, 0, 0, 0, 21, 94, 1, 0, 0, 0, 23, 99, 1, 0, 0, 0, 25, 105, 1, 0, 0, 0,
		27, 107, 1, 0, 0, 0, 29, 109, 1, 0, 0, 0, 31, 111, 1, 0, 0, 0, 33, 113,
		1, 0, 0, 0, 35, 115, 1, 0, 0, 0, 37, 117, 1, 0, 0, 0, 39, 119, 1, 0, 0,
		0, 41, 121, 1, 0, 0, 0, 43, 123, 1, 0, 0, 0, 45, 46, 5, 110, 0, 0, 46,
		47, 5, 117, 0, 0, 47, 48, 5, 109, 0, 0, 48, 49, 5, 98, 0, 0, 49, 50, 5,
		101, 0, 0, 50, 51, 5, 114, 0, 0, 51, 2, 1, 0, 0, 0, 52, 53, 5, 115, 0,
		0, 53, 54, 5, 116, 0, 0, 54, 55, 5, 114, 0, 0, 55, 56, 5, 105, 0, 0, 56,
		57, 5, 110, 0, 0, 57, 58, 5, 103, 0, 0, 58, 4, 1, 0, 0, 0, 59, 60, 5, 98,
		0, 0, 60, 61, 5, 111, 0, 0, 61, 62, 5, 111, 0, 0, 62, 63, 5, 108, 0, 0,
		63, 6, 1, 0, 0, 0, 64, 65, 5, 100, 0, 0, 65, 66, 5, 101, 0, 0, 66, 67,
		5, 102, 0, 0, 67, 8, 1, 0, 0, 0, 68, 69, 5, 33, 0, 0, 69, 70, 5, 61, 0,
		0, 70, 10, 1, 0, 0, 0, 71, 72, 5, 59, 0, 0, 72, 12, 1, 0, 0, 0, 73, 74,
		5, 60, 0, 0, 74, 75, 5, 61, 0, 0, 75, 14, 1, 0, 0, 0, 76, 77, 5, 62, 0,
		0, 77, 78, 5, 61, 0, 0, 78, 16, 1, 0, 0, 0, 79, 83, 5, 34, 0, 0, 80, 82,
		2, 32, 126, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0,
		0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87,
		5, 34, 0, 0, 87, 18, 1, 0, 0, 0, 88, 90, 7, 0, 0, 0, 89, 88, 1, 0, 0, 0,
		90, 91, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 20, 1,
		0, 0, 0, 93, 95, 2, 48, 57, 0, 94, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0,
		96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 22, 1, 0, 0, 0, 98, 100, 7,
		1, 0, 0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0,
		101, 102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 6, 11, 0, 0, 104,
		24, 1, 0, 0, 0, 105, 106, 7, 2, 0, 0, 106, 26, 1, 0, 0, 0, 107, 108, 7,
		3, 0, 0, 108, 28, 1, 0, 0, 0, 109, 110, 5, 44, 0, 0, 110, 30, 1, 0, 0,
		0, 111, 112, 5, 40, 0, 0, 112, 32, 1, 0, 0, 0, 113, 114, 5, 41, 0, 0, 114,
		34, 1, 0, 0, 0, 115, 116, 5, 123, 0, 0, 116, 36, 1, 0, 0, 0, 117, 118,
		5, 125, 0, 0, 118, 38, 1, 0, 0, 0, 119, 120, 5, 61, 0, 0, 120, 40, 1, 0,
		0, 0, 121, 122, 5, 60, 0, 0, 122, 42, 1, 0, 0, 0, 123, 124, 5, 62, 0, 0,
		124, 44, 1, 0, 0, 0, 5, 0, 83, 91, 96, 101, 1, 6, 0, 0,
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

// neonLexerInit initializes any static state used to implement neonLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewneonLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NeonLexerInit() {
	staticData := &neonlexerLexerStaticData
	staticData.once.Do(neonlexerLexerInit)
}

// NewneonLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewneonLexer(input antlr.CharStream) *neonLexer {
	NeonLexerInit()
	l := new(neonLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &neonlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "neon.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// neonLexer tokens.
const (
	neonLexerNUMTYPE   = 1
	neonLexerSTRTYPE   = 2
	neonLexerBOOLTYPE  = 3
	neonLexerDEF       = 4
	neonLexerNOTEQUAL  = 5
	neonLexerSEMI      = 6
	neonLexerLESSEQ    = 7
	neonLexerGREATEREQ = 8
	neonLexerSTRING    = 9
	neonLexerID        = 10
	neonLexerINT       = 11
	neonLexerWS        = 12
	neonLexerADD_SUB   = 13
	neonLexerMUL_DIV   = 14
	neonLexerCOMMA     = 15
	neonLexerLb        = 16
	neonLexerRb        = 17
	neonLexerLc        = 18
	neonLexerRc        = 19
	neonLexerEQUAL     = 20
	neonLexerLESS      = 21
	neonLexerGREATER   = 22
)
