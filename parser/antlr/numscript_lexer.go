// Code generated from Numscript.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type NumscriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var NumscriptLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func numscriptlexerLexerInit() {
	staticData := &NumscriptLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'max'", "'source'", "'destination'", "'send'",
		"'from'", "'to'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "MAX", "SOURCE",
		"DESTINATION", "SEND", "FROM", "TO", "LPARENS", "RPARENS", "LBRACKET",
		"RBRACKET", "LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL", "PERCENTAGE_PORTION_LITERAL",
		"NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.RuleNames = []string{
		"WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "MAX", "SOURCE",
		"DESTINATION", "SEND", "FROM", "TO", "LPARENS", "RPARENS", "LBRACKET",
		"RBRACKET", "LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL", "PERCENTAGE_PORTION_LITERAL",
		"NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 208, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 4, 0, 49, 8, 0, 11, 0, 12, 0, 50,
		1, 0, 1, 0, 1, 1, 4, 1, 56, 8, 1, 11, 1, 12, 1, 57, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 5, 2, 65, 8, 2, 10, 2, 12, 2, 68, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 79, 8, 3, 10, 3, 12, 3, 82, 9, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 4, 17, 139, 8, 17,
		11, 17, 12, 17, 140, 1, 17, 3, 17, 144, 8, 17, 1, 17, 1, 17, 3, 17, 148,
		8, 17, 1, 17, 4, 17, 151, 8, 17, 11, 17, 12, 17, 152, 1, 18, 4, 18, 156,
		8, 18, 11, 18, 12, 18, 157, 1, 18, 1, 18, 4, 18, 162, 8, 18, 11, 18, 12,
		18, 163, 3, 18, 166, 8, 18, 1, 18, 1, 18, 1, 19, 4, 19, 171, 8, 19, 11,
		19, 12, 19, 172, 1, 20, 1, 20, 4, 20, 177, 8, 20, 11, 20, 12, 20, 178,
		1, 20, 5, 20, 182, 8, 20, 10, 20, 12, 20, 185, 9, 20, 1, 21, 1, 21, 4,
		21, 189, 8, 21, 11, 21, 12, 21, 190, 1, 21, 1, 21, 4, 21, 195, 8, 21, 11,
		21, 12, 21, 196, 5, 21, 199, 8, 21, 10, 21, 12, 21, 202, 9, 21, 1, 22,
		4, 22, 205, 8, 22, 11, 22, 12, 22, 206, 2, 66, 80, 0, 23, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 1, 0, 8, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13,
		1, 0, 48, 57, 1, 0, 32, 32, 2, 0, 95, 95, 97, 122, 3, 0, 48, 57, 95, 95,
		97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 47, 57, 65,
		90, 226, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 1, 48, 1, 0, 0, 0, 3, 55, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 74, 1,
		0, 0, 0, 9, 87, 1, 0, 0, 0, 11, 91, 1, 0, 0, 0, 13, 98, 1, 0, 0, 0, 15,
		110, 1, 0, 0, 0, 17, 115, 1, 0, 0, 0, 19, 120, 1, 0, 0, 0, 21, 123, 1,
		0, 0, 0, 23, 125, 1, 0, 0, 0, 25, 127, 1, 0, 0, 0, 27, 129, 1, 0, 0, 0,
		29, 131, 1, 0, 0, 0, 31, 133, 1, 0, 0, 0, 33, 135, 1, 0, 0, 0, 35, 138,
		1, 0, 0, 0, 37, 155, 1, 0, 0, 0, 39, 170, 1, 0, 0, 0, 41, 174, 1, 0, 0,
		0, 43, 186, 1, 0, 0, 0, 45, 204, 1, 0, 0, 0, 47, 49, 7, 0, 0, 0, 48, 47,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0,
		51, 52, 1, 0, 0, 0, 52, 53, 6, 0, 0, 0, 53, 2, 1, 0, 0, 0, 54, 56, 7, 1,
		0, 0, 55, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58,
		1, 0, 0, 0, 58, 4, 1, 0, 0, 0, 59, 60, 5, 47, 0, 0, 60, 61, 5, 42, 0, 0,
		61, 66, 1, 0, 0, 0, 62, 65, 3, 5, 2, 0, 63, 65, 9, 0, 0, 0, 64, 62, 1,
		0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 66,
		64, 1, 0, 0, 0, 67, 69, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 70, 5, 42,
		0, 0, 70, 71, 5, 47, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 6, 2, 0, 0, 73,
		6, 1, 0, 0, 0, 74, 75, 5, 47, 0, 0, 75, 76, 5, 47, 0, 0, 76, 80, 1, 0,
		0, 0, 77, 79, 9, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 81,
		1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0,
		83, 84, 3, 3, 1, 0, 84, 85, 1, 0, 0, 0, 85, 86, 6, 3, 0, 0, 86, 8, 1, 0,
		0, 0, 87, 88, 5, 109, 0, 0, 88, 89, 5, 97, 0, 0, 89, 90, 5, 120, 0, 0,
		90, 10, 1, 0, 0, 0, 91, 92, 5, 115, 0, 0, 92, 93, 5, 111, 0, 0, 93, 94,
		5, 117, 0, 0, 94, 95, 5, 114, 0, 0, 95, 96, 5, 99, 0, 0, 96, 97, 5, 101,
		0, 0, 97, 12, 1, 0, 0, 0, 98, 99, 5, 100, 0, 0, 99, 100, 5, 101, 0, 0,
		100, 101, 5, 115, 0, 0, 101, 102, 5, 116, 0, 0, 102, 103, 5, 105, 0, 0,
		103, 104, 5, 110, 0, 0, 104, 105, 5, 97, 0, 0, 105, 106, 5, 116, 0, 0,
		106, 107, 5, 105, 0, 0, 107, 108, 5, 111, 0, 0, 108, 109, 5, 110, 0, 0,
		109, 14, 1, 0, 0, 0, 110, 111, 5, 115, 0, 0, 111, 112, 5, 101, 0, 0, 112,
		113, 5, 110, 0, 0, 113, 114, 5, 100, 0, 0, 114, 16, 1, 0, 0, 0, 115, 116,
		5, 102, 0, 0, 116, 117, 5, 114, 0, 0, 117, 118, 5, 111, 0, 0, 118, 119,
		5, 109, 0, 0, 119, 18, 1, 0, 0, 0, 120, 121, 5, 116, 0, 0, 121, 122, 5,
		111, 0, 0, 122, 20, 1, 0, 0, 0, 123, 124, 5, 40, 0, 0, 124, 22, 1, 0, 0,
		0, 125, 126, 5, 41, 0, 0, 126, 24, 1, 0, 0, 0, 127, 128, 5, 91, 0, 0, 128,
		26, 1, 0, 0, 0, 129, 130, 5, 93, 0, 0, 130, 28, 1, 0, 0, 0, 131, 132, 5,
		123, 0, 0, 132, 30, 1, 0, 0, 0, 133, 134, 5, 125, 0, 0, 134, 32, 1, 0,
		0, 0, 135, 136, 5, 61, 0, 0, 136, 34, 1, 0, 0, 0, 137, 139, 7, 2, 0, 0,
		138, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 143, 1, 0, 0, 0, 142, 144, 7, 3, 0, 0, 143, 142,
		1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 147, 5, 47,
		0, 0, 146, 148, 7, 3, 0, 0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0,
		148, 150, 1, 0, 0, 0, 149, 151, 7, 2, 0, 0, 150, 149, 1, 0, 0, 0, 151,
		152, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 36, 1,
		0, 0, 0, 154, 156, 7, 2, 0, 0, 155, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0,
		0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 165, 1, 0, 0, 0, 159,
		161, 5, 46, 0, 0, 160, 162, 7, 2, 0, 0, 161, 160, 1, 0, 0, 0, 162, 163,
		1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 1, 0,
		0, 0, 165, 159, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0,
		167, 168, 5, 37, 0, 0, 168, 38, 1, 0, 0, 0, 169, 171, 7, 2, 0, 0, 170,
		169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173,
		1, 0, 0, 0, 173, 40, 1, 0, 0, 0, 174, 176, 5, 36, 0, 0, 175, 177, 7, 4,
		0, 0, 176, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0,
		178, 179, 1, 0, 0, 0, 179, 183, 1, 0, 0, 0, 180, 182, 7, 5, 0, 0, 181,
		180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 42, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 188, 5, 64,
		0, 0, 187, 189, 7, 6, 0, 0, 188, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0,
		190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 200, 1, 0, 0, 0, 192,
		194, 5, 58, 0, 0, 193, 195, 7, 6, 0, 0, 194, 193, 1, 0, 0, 0, 195, 196,
		1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 1, 0,
		0, 0, 198, 192, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0,
		200, 201, 1, 0, 0, 0, 201, 44, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 205,
		7, 7, 0, 0, 204, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 204, 1, 0,
		0, 0, 206, 207, 1, 0, 0, 0, 207, 46, 1, 0, 0, 0, 20, 0, 50, 57, 64, 66,
		80, 140, 143, 147, 152, 157, 163, 165, 172, 178, 183, 190, 196, 200, 206,
		1, 6, 0, 0,
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

// NumscriptLexerInit initializes any static state used to implement NumscriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNumscriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NumscriptLexerInit() {
	staticData := &NumscriptLexerLexerStaticData
	staticData.once.Do(numscriptlexerLexerInit)
}

// NewNumscriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNumscriptLexer(input antlr.CharStream) *NumscriptLexer {
	NumscriptLexerInit()
	l := new(NumscriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &NumscriptLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Numscript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumscriptLexer tokens.
const (
	NumscriptLexerWS                         = 1
	NumscriptLexerNEWLINE                    = 2
	NumscriptLexerMULTILINE_COMMENT          = 3
	NumscriptLexerLINE_COMMENT               = 4
	NumscriptLexerMAX                        = 5
	NumscriptLexerSOURCE                     = 6
	NumscriptLexerDESTINATION                = 7
	NumscriptLexerSEND                       = 8
	NumscriptLexerFROM                       = 9
	NumscriptLexerTO                         = 10
	NumscriptLexerLPARENS                    = 11
	NumscriptLexerRPARENS                    = 12
	NumscriptLexerLBRACKET                   = 13
	NumscriptLexerRBRACKET                   = 14
	NumscriptLexerLBRACE                     = 15
	NumscriptLexerRBRACE                     = 16
	NumscriptLexerEQ                         = 17
	NumscriptLexerRATIO_PORTION_LITERAL      = 18
	NumscriptLexerPERCENTAGE_PORTION_LITERAL = 19
	NumscriptLexerNUMBER                     = 20
	NumscriptLexerVARIABLE_NAME              = 21
	NumscriptLexerACCOUNT                    = 22
	NumscriptLexerASSET                      = 23
)
