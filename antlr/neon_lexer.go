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
		"", "'number'", "'string'", "'bool'", "'def'", "'return'", "'if'", "'while'",
		"'=='", "'!='", "';'", "'<='", "'>='", "", "", "", "", "", "", "','",
		"'('", "')'", "'{'", "'}'", "'='", "'<'", "'>'",
	}
	staticData.symbolicNames = []string{
		"", "NUMTYPE", "STRTYPE", "BOOLTYPE", "DEF", "RETURN", "IF", "WHILE",
		"EQUALITY", "NOTEQUAL", "SEMI", "LESSEQ", "GREATEREQ", "STRING", "ID",
		"INT", "WS", "ADD_SUB", "MDM", "COMMA", "Lb", "Rb", "Lc", "Rc", "EQUAL",
		"LESS", "GREATER",
	}
	staticData.ruleNames = []string{
		"NUMTYPE", "STRTYPE", "BOOLTYPE", "DEF", "RETURN", "IF", "WHILE", "EQUALITY",
		"NOTEQUAL", "SEMI", "LESSEQ", "GREATEREQ", "STRING", "ID", "INT", "WS",
		"ADD_SUB", "MDM", "COMMA", "Lb", "Rb", "Lc", "Rc", "EQUAL", "LESS",
		"GREATER",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 152, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 5, 12, 109, 8, 12, 10,
		12, 12, 12, 112, 9, 12, 1, 12, 1, 12, 1, 13, 4, 13, 117, 8, 13, 11, 13,
		12, 13, 118, 1, 14, 4, 14, 122, 8, 14, 11, 14, 12, 14, 123, 1, 15, 4, 15,
		127, 8, 15, 11, 15, 12, 15, 128, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 0, 0, 26, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 1, 0, 4, 2, 0, 65, 90, 97, 122, 3, 0, 9, 10,
		13, 13, 32, 32, 2, 0, 43, 43, 45, 45, 3, 0, 37, 37, 42, 42, 47, 47, 155,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 1, 53, 1, 0, 0, 0,
		3, 60, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 72, 1, 0, 0, 0, 9, 76, 1, 0, 0,
		0, 11, 83, 1, 0, 0, 0, 13, 86, 1, 0, 0, 0, 15, 92, 1, 0, 0, 0, 17, 95,
		1, 0, 0, 0, 19, 98, 1, 0, 0, 0, 21, 100, 1, 0, 0, 0, 23, 103, 1, 0, 0,
		0, 25, 106, 1, 0, 0, 0, 27, 116, 1, 0, 0, 0, 29, 121, 1, 0, 0, 0, 31, 126,
		1, 0, 0, 0, 33, 132, 1, 0, 0, 0, 35, 134, 1, 0, 0, 0, 37, 136, 1, 0, 0,
		0, 39, 138, 1, 0, 0, 0, 41, 140, 1, 0, 0, 0, 43, 142, 1, 0, 0, 0, 45, 144,
		1, 0, 0, 0, 47, 146, 1, 0, 0, 0, 49, 148, 1, 0, 0, 0, 51, 150, 1, 0, 0,
		0, 53, 54, 5, 110, 0, 0, 54, 55, 5, 117, 0, 0, 55, 56, 5, 109, 0, 0, 56,
		57, 5, 98, 0, 0, 57, 58, 5, 101, 0, 0, 58, 59, 5, 114, 0, 0, 59, 2, 1,
		0, 0, 0, 60, 61, 5, 115, 0, 0, 61, 62, 5, 116, 0, 0, 62, 63, 5, 114, 0,
		0, 63, 64, 5, 105, 0, 0, 64, 65, 5, 110, 0, 0, 65, 66, 5, 103, 0, 0, 66,
		4, 1, 0, 0, 0, 67, 68, 5, 98, 0, 0, 68, 69, 5, 111, 0, 0, 69, 70, 5, 111,
		0, 0, 70, 71, 5, 108, 0, 0, 71, 6, 1, 0, 0, 0, 72, 73, 5, 100, 0, 0, 73,
		74, 5, 101, 0, 0, 74, 75, 5, 102, 0, 0, 75, 8, 1, 0, 0, 0, 76, 77, 5, 114,
		0, 0, 77, 78, 5, 101, 0, 0, 78, 79, 5, 116, 0, 0, 79, 80, 5, 117, 0, 0,
		80, 81, 5, 114, 0, 0, 81, 82, 5, 110, 0, 0, 82, 10, 1, 0, 0, 0, 83, 84,
		5, 105, 0, 0, 84, 85, 5, 102, 0, 0, 85, 12, 1, 0, 0, 0, 86, 87, 5, 119,
		0, 0, 87, 88, 5, 104, 0, 0, 88, 89, 5, 105, 0, 0, 89, 90, 5, 108, 0, 0,
		90, 91, 5, 101, 0, 0, 91, 14, 1, 0, 0, 0, 92, 93, 5, 61, 0, 0, 93, 94,
		5, 61, 0, 0, 94, 16, 1, 0, 0, 0, 95, 96, 5, 33, 0, 0, 96, 97, 5, 61, 0,
		0, 97, 18, 1, 0, 0, 0, 98, 99, 5, 59, 0, 0, 99, 20, 1, 0, 0, 0, 100, 101,
		5, 60, 0, 0, 101, 102, 5, 61, 0, 0, 102, 22, 1, 0, 0, 0, 103, 104, 5, 62,
		0, 0, 104, 105, 5, 61, 0, 0, 105, 24, 1, 0, 0, 0, 106, 110, 5, 34, 0, 0,
		107, 109, 2, 32, 126, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110,
		108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113, 1, 0, 0, 0, 112, 110,
		1, 0, 0, 0, 113, 114, 5, 34, 0, 0, 114, 26, 1, 0, 0, 0, 115, 117, 7, 0,
		0, 0, 116, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0,
		118, 119, 1, 0, 0, 0, 119, 28, 1, 0, 0, 0, 120, 122, 2, 48, 57, 0, 121,
		120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124,
		1, 0, 0, 0, 124, 30, 1, 0, 0, 0, 125, 127, 7, 1, 0, 0, 126, 125, 1, 0,
		0, 0, 127, 128, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0,
		129, 130, 1, 0, 0, 0, 130, 131, 6, 15, 0, 0, 131, 32, 1, 0, 0, 0, 132,
		133, 7, 2, 0, 0, 133, 34, 1, 0, 0, 0, 134, 135, 7, 3, 0, 0, 135, 36, 1,
		0, 0, 0, 136, 137, 5, 44, 0, 0, 137, 38, 1, 0, 0, 0, 138, 139, 5, 40, 0,
		0, 139, 40, 1, 0, 0, 0, 140, 141, 5, 41, 0, 0, 141, 42, 1, 0, 0, 0, 142,
		143, 5, 123, 0, 0, 143, 44, 1, 0, 0, 0, 144, 145, 5, 125, 0, 0, 145, 46,
		1, 0, 0, 0, 146, 147, 5, 61, 0, 0, 147, 48, 1, 0, 0, 0, 148, 149, 5, 60,
		0, 0, 149, 50, 1, 0, 0, 0, 150, 151, 5, 62, 0, 0, 151, 52, 1, 0, 0, 0,
		5, 0, 110, 118, 123, 128, 1, 6, 0, 0,
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
	neonLexerRETURN    = 5
	neonLexerIF        = 6
	neonLexerWHILE     = 7
	neonLexerEQUALITY  = 8
	neonLexerNOTEQUAL  = 9
	neonLexerSEMI      = 10
	neonLexerLESSEQ    = 11
	neonLexerGREATEREQ = 12
	neonLexerSTRING    = 13
	neonLexerID        = 14
	neonLexerINT       = 15
	neonLexerWS        = 16
	neonLexerADD_SUB   = 17
	neonLexerMDM       = 18
	neonLexerCOMMA     = 19
	neonLexerLb        = 20
	neonLexerRb        = 21
	neonLexerLc        = 22
	neonLexerRc        = 23
	neonLexerEQUAL     = 24
	neonLexerLESS      = 25
	neonLexerGREATER   = 26
)
