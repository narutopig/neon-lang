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
		"", "'number'", "'string'", "'bool'", "'true'", "'false'", "'void'",
		"'def'", "'return'", "'if'", "'elif'", "'else'", "'while'", "'=='",
		"'!='", "';'", "'<='", "'>='", "", "", "", "", "", "", "','", "'('",
		"')'", "'{'", "'}'", "'='", "'<'", "'>'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "NUMTYPE", "STRTYPE", "BOOLTYPE", "TRUE", "FALSE", "VOID", "DEF",
		"RETURN", "IF", "ELIF", "ELSE", "WHILE", "EQUALITY", "NOTEQUAL", "SEMI",
		"LESSEQ", "GREATEREQ", "STRING", "ID", "INT", "WS", "ADD_SUB", "MDM",
		"COMMA", "Lb", "Rb", "Lc", "Rc", "EQUAL", "LESS", "GREATER", "BANG",
	}
	staticData.ruleNames = []string{
		"NUMTYPE", "STRTYPE", "BOOLTYPE", "TRUE", "FALSE", "VOID", "DEF", "RETURN",
		"IF", "ELIF", "ELSE", "WHILE", "EQUALITY", "NOTEQUAL", "SEMI", "LESSEQ",
		"GREATEREQ", "STRING", "ID", "INT", "WS", "ADD_SUB", "MDM", "COMMA",
		"Lb", "Rb", "Lc", "Rc", "EQUAL", "LESS", "GREATER", "BANG",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 192, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 5, 17, 147, 8, 17, 10, 17, 12, 17, 150, 9, 17, 1, 17,
		1, 17, 1, 18, 4, 18, 155, 8, 18, 11, 18, 12, 18, 156, 1, 19, 4, 19, 160,
		8, 19, 11, 19, 12, 19, 161, 1, 20, 4, 20, 165, 8, 20, 11, 20, 12, 20, 166,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 31, 1, 31, 0, 0, 32, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 1, 0, 4, 2, 0,
		65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 43, 43, 45, 45, 3,
		0, 37, 37, 42, 42, 47, 47, 195, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 1, 65, 1, 0, 0, 0,
		3, 72, 1, 0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 84, 1, 0, 0, 0, 9, 89, 1, 0, 0,
		0, 11, 95, 1, 0, 0, 0, 13, 100, 1, 0, 0, 0, 15, 104, 1, 0, 0, 0, 17, 111,
		1, 0, 0, 0, 19, 114, 1, 0, 0, 0, 21, 119, 1, 0, 0, 0, 23, 124, 1, 0, 0,
		0, 25, 130, 1, 0, 0, 0, 27, 133, 1, 0, 0, 0, 29, 136, 1, 0, 0, 0, 31, 138,
		1, 0, 0, 0, 33, 141, 1, 0, 0, 0, 35, 144, 1, 0, 0, 0, 37, 154, 1, 0, 0,
		0, 39, 159, 1, 0, 0, 0, 41, 164, 1, 0, 0, 0, 43, 170, 1, 0, 0, 0, 45, 172,
		1, 0, 0, 0, 47, 174, 1, 0, 0, 0, 49, 176, 1, 0, 0, 0, 51, 178, 1, 0, 0,
		0, 53, 180, 1, 0, 0, 0, 55, 182, 1, 0, 0, 0, 57, 184, 1, 0, 0, 0, 59, 186,
		1, 0, 0, 0, 61, 188, 1, 0, 0, 0, 63, 190, 1, 0, 0, 0, 65, 66, 5, 110, 0,
		0, 66, 67, 5, 117, 0, 0, 67, 68, 5, 109, 0, 0, 68, 69, 5, 98, 0, 0, 69,
		70, 5, 101, 0, 0, 70, 71, 5, 114, 0, 0, 71, 2, 1, 0, 0, 0, 72, 73, 5, 115,
		0, 0, 73, 74, 5, 116, 0, 0, 74, 75, 5, 114, 0, 0, 75, 76, 5, 105, 0, 0,
		76, 77, 5, 110, 0, 0, 77, 78, 5, 103, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80,
		5, 98, 0, 0, 80, 81, 5, 111, 0, 0, 81, 82, 5, 111, 0, 0, 82, 83, 5, 108,
		0, 0, 83, 6, 1, 0, 0, 0, 84, 85, 5, 116, 0, 0, 85, 86, 5, 114, 0, 0, 86,
		87, 5, 117, 0, 0, 87, 88, 5, 101, 0, 0, 88, 8, 1, 0, 0, 0, 89, 90, 5, 102,
		0, 0, 90, 91, 5, 97, 0, 0, 91, 92, 5, 108, 0, 0, 92, 93, 5, 115, 0, 0,
		93, 94, 5, 101, 0, 0, 94, 10, 1, 0, 0, 0, 95, 96, 5, 118, 0, 0, 96, 97,
		5, 111, 0, 0, 97, 98, 5, 105, 0, 0, 98, 99, 5, 100, 0, 0, 99, 12, 1, 0,
		0, 0, 100, 101, 5, 100, 0, 0, 101, 102, 5, 101, 0, 0, 102, 103, 5, 102,
		0, 0, 103, 14, 1, 0, 0, 0, 104, 105, 5, 114, 0, 0, 105, 106, 5, 101, 0,
		0, 106, 107, 5, 116, 0, 0, 107, 108, 5, 117, 0, 0, 108, 109, 5, 114, 0,
		0, 109, 110, 5, 110, 0, 0, 110, 16, 1, 0, 0, 0, 111, 112, 5, 105, 0, 0,
		112, 113, 5, 102, 0, 0, 113, 18, 1, 0, 0, 0, 114, 115, 5, 101, 0, 0, 115,
		116, 5, 108, 0, 0, 116, 117, 5, 105, 0, 0, 117, 118, 5, 102, 0, 0, 118,
		20, 1, 0, 0, 0, 119, 120, 5, 101, 0, 0, 120, 121, 5, 108, 0, 0, 121, 122,
		5, 115, 0, 0, 122, 123, 5, 101, 0, 0, 123, 22, 1, 0, 0, 0, 124, 125, 5,
		119, 0, 0, 125, 126, 5, 104, 0, 0, 126, 127, 5, 105, 0, 0, 127, 128, 5,
		108, 0, 0, 128, 129, 5, 101, 0, 0, 129, 24, 1, 0, 0, 0, 130, 131, 5, 61,
		0, 0, 131, 132, 5, 61, 0, 0, 132, 26, 1, 0, 0, 0, 133, 134, 5, 33, 0, 0,
		134, 135, 5, 61, 0, 0, 135, 28, 1, 0, 0, 0, 136, 137, 5, 59, 0, 0, 137,
		30, 1, 0, 0, 0, 138, 139, 5, 60, 0, 0, 139, 140, 5, 61, 0, 0, 140, 32,
		1, 0, 0, 0, 141, 142, 5, 62, 0, 0, 142, 143, 5, 61, 0, 0, 143, 34, 1, 0,
		0, 0, 144, 148, 5, 34, 0, 0, 145, 147, 2, 32, 126, 0, 146, 145, 1, 0, 0,
		0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		151, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 152, 5, 34, 0, 0, 152, 36,
		1, 0, 0, 0, 153, 155, 7, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0,
		0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 38, 1, 0, 0, 0,
		158, 160, 2, 48, 57, 0, 159, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161,
		159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 40, 1, 0, 0, 0, 163, 165, 7,
		1, 0, 0, 164, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 164, 1, 0, 0,
		0, 166, 167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 6, 20, 0, 0, 169,
		42, 1, 0, 0, 0, 170, 171, 7, 2, 0, 0, 171, 44, 1, 0, 0, 0, 172, 173, 7,
		3, 0, 0, 173, 46, 1, 0, 0, 0, 174, 175, 5, 44, 0, 0, 175, 48, 1, 0, 0,
		0, 176, 177, 5, 40, 0, 0, 177, 50, 1, 0, 0, 0, 178, 179, 5, 41, 0, 0, 179,
		52, 1, 0, 0, 0, 180, 181, 5, 123, 0, 0, 181, 54, 1, 0, 0, 0, 182, 183,
		5, 125, 0, 0, 183, 56, 1, 0, 0, 0, 184, 185, 5, 61, 0, 0, 185, 58, 1, 0,
		0, 0, 186, 187, 5, 60, 0, 0, 187, 60, 1, 0, 0, 0, 188, 189, 5, 62, 0, 0,
		189, 62, 1, 0, 0, 0, 190, 191, 5, 33, 0, 0, 191, 64, 1, 0, 0, 0, 5, 0,
		148, 156, 161, 166, 1, 6, 0, 0,
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
	neonLexerTRUE      = 4
	neonLexerFALSE     = 5
	neonLexerVOID      = 6
	neonLexerDEF       = 7
	neonLexerRETURN    = 8
	neonLexerIF        = 9
	neonLexerELIF      = 10
	neonLexerELSE      = 11
	neonLexerWHILE     = 12
	neonLexerEQUALITY  = 13
	neonLexerNOTEQUAL  = 14
	neonLexerSEMI      = 15
	neonLexerLESSEQ    = 16
	neonLexerGREATEREQ = 17
	neonLexerSTRING    = 18
	neonLexerID        = 19
	neonLexerINT       = 20
	neonLexerWS        = 21
	neonLexerADD_SUB   = 22
	neonLexerMDM       = 23
	neonLexerCOMMA     = 24
	neonLexerLb        = 25
	neonLexerRb        = 26
	neonLexerLc        = 27
	neonLexerRc        = 28
	neonLexerEQUAL     = 29
	neonLexerLESS      = 30
	neonLexerGREATER   = 31
	neonLexerBANG      = 32
)
